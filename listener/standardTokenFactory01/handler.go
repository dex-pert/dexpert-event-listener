package standardTokenFactory01

import (
    "github.com/ethereum/go-ethereum/common"
    "math/big"
    "context"
    el "github.com/x1rh/event-listener"
    "dexpert-event-listener/abi/standardTokenFactory01"
    "github.com/shopspring/decimal"
    "dexpert-event-listener/gorm/model"
    "time"
    "dexpert-event-listener/constant"
    "log/slog"
    "github.com/pkg/errors"
    "dexpert-event-listener/gorm/query"
)

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
                block, err := ltCtx.EthClient.BlockByNumber(ctx, new(big.Int).SetUint64(event.BlockNumber))
                if err != nil {
                    slog.Error("TokenCreated event", "fail to get block,err is: ", err)
                    return err
                }

                _fees, err := standardTokenFactory01.GetFees(ltCtx.StandardTokenFactory01Address, l.Level, block.Number(), ltCtx.EthClient)
                if err != nil {
                    slog.Error("TokenCreated event", "fail to get fees,err is: ", err)
                    return err
                }
                fee := decimal.NewFromBigInt(_fees, -(ltCtx.FeeDecimal)).String()

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
                    FeeTokenSymbol:  ltCtx.FeeSymbol,
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
                    FeeTokenSymbol:  ltCtx.FeeSymbol,
                    FeeTokenDecimal: ltCtx.FeeDecimal,
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
