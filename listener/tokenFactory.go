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
)

func NewTokenFactoryEventListener(ltCtx *Context) (*el.EventListener, error) {
    c := el.ChainConfig{
        ChainId:   ltCtx.Chain.ChainId,
        ChainName: ltCtx.Chain.ChainName,
        URL:       ltCtx.Chain.URL,
    }

    client := ltCtx.AbiProxy.WithChainID(c.ChainId).Client
    if client == nil {
        return nil, errors.Errorf("newTokenFactoryEventListener client is nil")
    }

    tokenFactory, err := el.NewContract(ltCtx.TokenFactoryAddress, ltCtx.TokenFactoryABIStr, big.NewInt(ltCtx.TokenFactoryBlockNumber), big.NewInt(ltCtx.TokenFactoryStep))
    if err != nil {
        panic(err)
    }
    tokenFactory.SetLogHandler(tokenFactoryLogHandler(ltCtx, tokenFactory))

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

func tokenFactoryLogHandler(ltCtx *Context, c *el.Contract) el.LogHandleFunc {
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
                    Fee:             ltCtx.TokenFactoryLaunchFee,
                    FeeTokenSymbol:  ltCtx.TokenFactoryTokenSymbol,
                    Timestamp:       eventTime,
                    TypeName:        constant.WithSwapName(constant.SwapTypeLaunch),
                    ChainName:       ltCtx.Chain.ChainName,
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
                    Fee:             ltCtx.TokenFactoryLaunchFee,
                    FeeTokenSymbol:  ltCtx.TokenFactoryTokenSymbol,
                    FeeTokenDecimal: ltCtx.TokenFactoryTokenDecimal,
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
