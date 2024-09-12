package listener

import (
    "context"
    "log/slog"
    "math/big"
    el "github.com/x1rh/event-listener"

    "github.com/ethereum/go-ethereum/common"
    "github.com/pkg/errors"
    "dexpert-event-listener/gorm/query"
    "dexpert-event-listener/gorm/model"
    "dexpert-event-listener/constant"
    "time"
    "github.com/shopspring/decimal"
    "dexpert-event-listener/abi/standardTokenFactory01"
)

func NewStandardTokenFactory01EventListener(ltCtx *Context) (*el.EventListener, error) {
    c := el.ChainConfig{
        ChainId:   ltCtx.Chain.ChainId,
        ChainName: ltCtx.Chain.ChainName,
        URL:       ltCtx.Chain.URL,
    }

    client := ltCtx.AbiProxy.WithChainID(c.ChainId).Client
    if client == nil {
        return nil, errors.Errorf("newStandardTokenFactory01EventListener client is nil")
    }

    tokenFactory, err := el.NewContract(ltCtx.StandardTokenFactory01Address, ltCtx.StandardTokenFactory01ABIStr, big.NewInt(ltCtx.StandardTokenFactory01BlockNumber), big.NewInt(ltCtx.StandardTokenFactory01Step))
    if err != nil {
        panic(err)
    }
    tokenFactory.SetLogHandler(standardTokenFactory01EventLogHandler(ltCtx, tokenFactory))

    _el, err := el.New(
        c,
        el.WithClient(client),
        el.WithContract(*tokenFactory),
    )
    if err != nil {
        return nil, errors.Wrap(err, "fail to new an EventListener object")
    }
    return _el, nil
}

// LogTokenCreated signature: TokenCreated(address,address,uint8,uint96,uint256)
type LogTokenCreated struct {
    Owner        common.Address // token owner
    Token        common.Address // token address
    TokenType    uint8
    TokenVersion *big.Int
    Level        *big.Int
}

func standardTokenFactory01EventLogHandler(ltCtx *Context, c *el.Contract) el.LogHandleFunc {
    return func(ctx context.Context, event *el.Event) error {
        switch event.Name {
        case "TokenCreated":
            var l LogTokenCreated
            if err := c.Abi.UnpackIntoInterface(&l, event.Name, event.Data); err != nil {
                return errors.Wrap(err, "fail to unpack log")
            }
            // handle indexed topic
            l.Owner = el.HashToAddress(event.IndexedParams[0])
            l.Token = el.HashToAddress(event.IndexedParams[1])
            slog.Info("TokenCreated event", slog.Any("event", l))

            _ = query.Q.Transaction(func(tx *query.Query) error {
                block, err := ltCtx.AbiProxy.WithChainID(ltCtx.Chain.ChainId).BlockByNumber(ctx, new(big.Int).SetUint64(event.BlockNumber))
                if err != nil {
                    slog.Error("TokenCreated event", "fail to get block,err is: ", err)
                    return err
                }

                _fees, err := standardTokenFactory01.GetFees(ltCtx.StandardTokenFactory01Address, l.Level, block.Number(), ltCtx.AbiProxy.WithChainID(ltCtx.Chain.ChainId).Client)
                if err != nil {
                    slog.Error("TokenCreated event", "fail to get fees,err is: ", err)
                    return err
                }
                fee := decimal.NewFromBigInt(_fees, -(ltCtx.StandardTokenFactory01FeeDecimal)).String()

                uw := tx.UserWallet
                userWallet, err := uw.WithContext(ctx).Where(uw.Address.Eq(l.Owner.String()), uw.ChainID.Eq(int32(ltCtx.Chain.ChainId))).Take()
                if err != nil {
                    userWallet = &model.UserWallet{UID: 0}
                    slog.Error("TokenCreated event", "fail to get user wallet,err is: ", err)
                }

                eventTime := time.Unix(int64(block.Time()), 0)
                userLaunchTx := model.UserLaunchTx{
                    UID:             userWallet.UID,
                    ContractAddress: l.Token.String(),
                    ChainID:         int32(ltCtx.Chain.ChainId),
                    PairAddress:     "",
                    Fee:             fee,
                    FeeTokenSymbol:  ltCtx.StandardTokenFactory01FeeSymbol,
                    Timestamp:       eventTime,
                    TypeName:        constant.WithSwapName(constant.SwapTypeLaunch),
                    ChainName:       ltCtx.Chain.ChainName,
                    Tx:              event.TxHash.String(),
                    Owner:           l.Owner.String(),
                    Level:           l.Level.String(),
                }
                if err = tx.WithContext(ctx).UserLaunchTx.Create(&userLaunchTx); err != nil {
                    slog.Error("TokenCreated event", "fail to create user launch tx,err is: ", err)
                    return errors.Wrap(err, "fail to create user launch tx")
                }

                if err = tx.WithContext(ctx).UserTransaction.Create(&model.UserTransaction{
                    UID:             userWallet.UID,
                    Tid:             userLaunchTx.ID,
                    SwapType:        constant.SwapTypeLaunch,
                    Volume:          "0.00",
                    Timestamp:       eventTime,
                    SwapName:        constant.WithSwapName(constant.SwapTypeLaunch),
                    ChainName:       ltCtx.Chain.ChainName,
                    ChainID:         int32(ltCtx.Chain.ChainId),
                    Fee:             fee,
                    FeeTokenSymbol:  ltCtx.StandardTokenFactory01FeeSymbol,
                    FeeTokenDecimal: ltCtx.StandardTokenFactory01FeeDecimal,
                }); err != nil {
                    slog.Error("TokenCreated event", "fail to create user transaction,err is: ", err)
                    return errors.Wrap(err, "fail to create user transaction")
                }

                return nil
            })
        default:
            // do nothing
        }
        return nil
    }
}
