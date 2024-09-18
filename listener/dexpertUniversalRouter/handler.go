package dexpertUniversalRouter

import (
    "time"
    "github.com/ethereum/go-ethereum/common"
    "math/big"
    "context"
    "dexpert-event-listener/abi/erc20"
    "github.com/shopspring/decimal"
    "dexpert-event-listener/abi/uniswapv2router"
    "dexpert-event-listener/gorm/model"
    "dexpert-event-listener/constant"
    el "github.com/x1rh/event-listener"
    "log/slog"
    "github.com/pkg/errors"
    "dexpert-event-listener/gorm/query"
    "gorm.io/gorm/clause"
)

// LogPaymentFee event PaymentFee(address payer, address tokenIn, uint256 amountIn, address feeToken, uint256 feeAmount, uint256 level, uint256 swapType, uint256 feeBps, uint256 feeBaseBps)
type LogPaymentFee struct {
    Payer      common.Address // payer address
    TokenIn    common.Address // token in address
    AmountIn   *big.Int
    FeeToken   common.Address
    FeeAmount  *big.Int
    Level      *big.Int
    SwapType   *big.Int
    FeeBps     *big.Int
    FeeBaseBps *big.Int
}

func dexpertUniversalRouterLogHandler(ltCtx *Context, c *el.Contract) el.LogHandleFunc {
    return func(ctx context.Context, event *el.Event) error {
        switch event.Name {
        case "Payment", "PaymentFee":
            var l LogPaymentFee
            if err := c.Abi.UnpackIntoInterface(&l, event.Name, event.Data); err != nil {
                return errors.Wrap(err, "fail to unpack log")
            }
            slog.Info("PaymentFee event", slog.Any("event", l))

            block, err := ltCtx.EthClient.HeaderByNumber(ctx, new(big.Int).SetUint64(event.BlockNumber))
            if err != nil {
                slog.Error("PaymentFee event", "fail to get block", err, "block number", event.BlockNumber)
                return err
            }

            var (
                tokenSymbol, tokenName string
                tokenDecimal           int32
            )
            if l.TokenIn.String() == ltCtx.EthAddress || l.TokenIn.String() == ltCtx.WethAddress {
                tokenSymbol = ltCtx.EthSymbol
                tokenName = ltCtx.EthName
                tokenDecimal = ltCtx.EthDecimal
            } else {
                tokenSymbol, tokenName, tokenDecimal, err = erc20.GetSymbolNameDecimalByAddress(l.TokenIn, ltCtx.EthClient)
                if err != nil {
                    slog.Error("PaymentFee event", "fail to get SymbolNameDecimalByAddress", err, "token in", l.TokenIn)
                    return err
                }
            }

            var (
                feeTokenSymbol, feeTokenName string
                feeTokenDecimal              int32
            )
            if l.FeeToken.String() == ltCtx.EthAddress || l.FeeToken.String() == ltCtx.WethAddress {
                feeTokenSymbol = ltCtx.EthSymbol
                feeTokenName = ltCtx.EthName
                feeTokenDecimal = ltCtx.EthDecimal
            } else {
                feeTokenSymbol, feeTokenName, feeTokenDecimal, err = erc20.GetSymbolNameDecimalByAddress(l.FeeToken, ltCtx.EthClient)
                if err != nil {
                    slog.Error("PaymentFee event", "erc20 getSymbolNameDecimalByAddress", err, "fee token", l.FeeToken)
                    return err
                }
            }

            fee := decimal.NewFromBigInt(l.FeeAmount, -feeTokenDecimal).String()
            volume := ""
            if l.TokenIn.String() == ltCtx.USDTAddress {
                volume = decimal.NewFromBigInt(l.AmountIn, -tokenDecimal).String()
            } else {
                var swapPath []common.Address
                if l.TokenIn.String() == ltCtx.EthAddress || l.TokenIn.String() == ltCtx.WethAddress {
                    swapPath = []common.Address{common.HexToAddress(ltCtx.WethAddress), common.HexToAddress(ltCtx.USDTAddress)}
                } else {
                    swapPath = []common.Address{l.TokenIn, common.HexToAddress(ltCtx.WethAddress), common.HexToAddress(ltCtx.USDTAddress)}
                }
                outAmounts, err := uniswapv2router.GetAmountsOutByAddressAndBlockNumber(ltCtx.UniswapV2RouterAddress, new(big.Int).SetUint64(event.BlockNumber), l.AmountIn, ltCtx.EthClient, swapPath)
                if err != nil {
                    slog.Error("PaymentFee event", "uniswapv2router getAmountsOutByAddressAndBlockNumber", err, "uniswapV2RouterAddress", ltCtx.UniswapV2RouterAddress, "block number", event.BlockNumber, "amount in", l.AmountIn, "path", swapPath)
                    return err
                }
                volume = decimal.NewFromBigInt(outAmounts[len(outAmounts)-1], -ltCtx.USDTDecimal).String()
                // volume = _volume.Mul(decimal.NewFromBigInt(l.AmountIn, -int32(tokenDecimal))).String()
            }
            if volume == "0" {
                volume = "0.00"
            }

            return query.Q.Transaction(func(tx *query.Query) error {
                txStr := event.TxHash.String()
                if len(txStr) == 0 {
                    slog.Error("PaymentFee event", "tx hash is null", txStr, "block number", event.BlockNumber, "token in", l.TokenIn)
                    return nil
                }

                uw := tx.UserWallet
                userWallet, err := uw.WithContext(ctx).Where(uw.Address.Eq(l.Payer.String())).Take()
                if err != nil {
                    userWallet = &model.UserWallet{UID: 0}
                    slog.Error("PaymentFee event", "fail to get user wallet", err)
                }

                now := time.Now().UTC()
                eventTime := time.Unix(int64(block.Time), 0)
                userSwapTx := model.UserSwapTx{
                    UID:             userWallet.UID,
                    Address:         l.Payer.String(),
                    Tx:              txStr,
                    TokenSymbol:     tokenSymbol,
                    AmountIn:        decimal.NewFromBigInt(l.AmountIn, tokenDecimal).String(),
                    TokenIn:         l.TokenIn.String(),
                    TransactionTime: eventTime,
                    Fee:             fee,
                    BlockNumber:     int32(event.BlockNumber),
                    ChainID:         int32(ltCtx.Chain.ChainId),
                    ChainName:       ltCtx.Chain.ChainName,
                    SwapType:        int32(l.SwapType.Int64()),
                    TokenName:       tokenName,
                    Decimal:         tokenDecimal,
                    FeeTokenName:    feeTokenName,
                    FeeDecimal:      feeTokenDecimal,
                    FeeTokenSymbol:  feeTokenSymbol,
                    CreatedAt:       now,
                }
                if err = tx.WithContext(ctx).UserSwapTx.Clauses(clause.OnConflict{DoNothing: true}).Create(&userSwapTx); err != nil {
                    slog.Error("PaymentFee event", "fail to create user swap tx", err)
                    return errors.Wrap(err, "fail to create user swap tx")
                }
                if userSwapTx.ID == 0 {
                    err = errors.Wrap(err, "new user swap tx is is 0")
                    slog.Error("PaymentFee event", "userSwapTx.ID == 0", err, "tx", txStr)
                    return err
                }

                if err = tx.WithContext(ctx).UserTransaction.Clauses(clause.OnConflict{DoNothing: true}).Create(&model.UserTransaction{
                    UID:             userWallet.UID,
                    Tid:             userSwapTx.ID,
                    SwapType:        userSwapTx.SwapType,
                    Volume:          volume,
                    Timestamp:       eventTime,
                    SwapName:        constant.WithSwapName(l.SwapType.Int64()),
                    ChainName:       ltCtx.Chain.ChainName,
                    ChainID:         int32(ltCtx.Chain.ChainId),
                    Fee:             fee,
                    FeeTokenSymbol:  feeTokenSymbol,
                    FeeTokenDecimal: feeTokenDecimal,
                    IdentifyAddress: txStr,
                    CreatedAt:       now,
                }); err != nil {
                    slog.Error("PaymentFee event", "fail to create user transaction", err)
                    return errors.Wrap(err, "fail to create user transaction")
                }

                if err = tx.ListenerNewestBlocknumber.WithContext(ctx).Clauses(clause.OnConflict{DoUpdates: clause.AssignmentColumns([]string{"block_number", "updated_at"})}).
                    Create(&model.ListenerNewestBlocknumber{
                        ContractAddress: ltCtx.DexpertUniversalRouterAddress,
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
                    ContractAddress: ltCtx.DexpertUniversalRouterAddress,
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
