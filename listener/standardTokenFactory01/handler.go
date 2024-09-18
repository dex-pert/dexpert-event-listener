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
    "gorm.io/gorm/clause"
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

            block, err := ltCtx.EthClient.HeaderByNumber(ctx, new(big.Int).SetUint64(event.BlockNumber))
            if err != nil {
                slog.Error("TokenCreated event", "fail to get block,err", err, "block number", event.BlockNumber)
                return err
            }

            _fees, err := standardTokenFactory01.GetFees(ltCtx.StandardTokenFactory01Address, l.Level, block.Number, ltCtx.EthClient)
            if err != nil {
                slog.Error("TokenCreated event", "fail to get fees,err", err, "standardTokenFactory01Address", ltCtx.StandardTokenFactory01Address, "level", l.Level, "block number", block.Number)
                return err
            }
            fee := decimal.NewFromBigInt(_fees, -(ltCtx.FeeDecimal)).String()

            return query.Q.Transaction(func(tx *query.Query) error {
                contractAddress := l.Token.String()
                if len(contractAddress) == 0 {
                    slog.Error("TokenCreated event", "contract address is null", contractAddress, "block number", event.BlockNumber, "token", l.Token, "owner", l.Owner.String())
                    return nil
                }

                uw := tx.UserWallet
                userWallet, err := uw.WithContext(ctx).Where(uw.Address.Eq(l.Owner.String())).Take()
                if err != nil {
                    userWallet = &model.UserWallet{UID: 0}
                    slog.Error("TokenCreated event", "fail to get user wallet,err", err)
                }

                now := time.Now().UTC()
                eventTime := time.Unix(int64(block.Time), 0).UTC()
                userLaunchTx := model.UserLaunchTx{
                    UID:             userWallet.UID,
                    ContractAddress: contractAddress,
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
                    CreatedAt:       now,
                    BlockNumber:     int32(event.BlockNumber),
                }
                if err = tx.WithContext(ctx).UserLaunchTx.Clauses(clause.OnConflict{DoNothing: true}).Create(&userLaunchTx); err != nil {
                    slog.Error("TokenCreated event", "fail to create user launch tx,err", err)
                    return errors.Wrap(err, "fail to create user launch tx")
                }
                if userLaunchTx.ID == 0 { // 说明 Clauses(clause.OnConflict{DoNothing: true}) 生效, 有重复数据，下面不再执行
                    err = errors.Wrap(err, "new user launch tx is is 0")
                    slog.Error("PaymentFee event", "userLaunchTx.ID == 0", err, "contract address", contractAddress)
                    return nil
                }

                if err = tx.WithContext(ctx).UserTransaction.Clauses(clause.OnConflict{DoNothing: true}).Create(&model.UserTransaction{
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
                    IdentifyAddress: contractAddress,
                    CreatedAt:       now,
                }); err != nil {
                    slog.Error("TokenCreated event", "fail to create user transaction,err", err)
                    return errors.Wrap(err, "fail to create user transaction")
                }

                if err = tx.ListenerNewestBlocknumber.WithContext(ctx).Clauses(clause.OnConflict{DoUpdates: clause.AssignmentColumns([]string{"block_number", "updated_at"})}).
                    Create(&model.ListenerNewestBlocknumber{
                        ContractAddress: ltCtx.StandardTokenFactory01Address,
                        ChainID:         int32(ltCtx.Chain.ChainId),
                        BlockNumber:     int32(event.BlockNumber),
                        CreatedAt:       time.Now().UTC(),
                        UpdatedAt:       time.Now().UTC(),
                    }); err != nil {
                    slog.Error("PaymentFee event", "fail to refresh listener block number", err)
                    return errors.Wrap(err, "fail to refresh listener block number")
                }

                return nil
            })
        default:
            listenerNewestBlockNumber := query.ListenerNewestBlocknumber
            if err := listenerNewestBlockNumber.WithContext(ctx).Clauses(clause.OnConflict{DoUpdates: clause.AssignmentColumns([]string{"block_number", "updated_at"})}).
                Create(&model.ListenerNewestBlocknumber{
                    ContractAddress: ltCtx.StandardTokenFactory01Address,
                    ChainID:         int32(ltCtx.Chain.ChainId),
                    BlockNumber:     int32(event.BlockNumber),
                    CreatedAt:       time.Now().UTC(),
                    UpdatedAt:       time.Now().UTC(),
                }); err != nil {
                slog.Error("PaymentFee event", "fail to refresh listener block number", err)
                return errors.Wrap(err, "fail to refresh listener block number")
            }
            return nil
        }
    }
}
