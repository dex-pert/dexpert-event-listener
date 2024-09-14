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

            _ = query.Q.Transaction(func(tx *query.Query) error {
                block, err := ltCtx.EthClient.HeaderByNumber(ctx, new(big.Int).SetUint64(event.BlockNumber))
                if err != nil {
                    slog.Error("PaymentFee event", "fail to get block,err", err, "block number", event.BlockNumber)
                    return err
                }
                tokenSymbol, tokenName, tokenDecimal, err := erc20.GetSymbolNameDecimalByAddress(l.TokenIn, ltCtx.EthClient)
                if err != nil {
                    slog.Error("PaymentFee event", "fail to get SymbolNameDecimalByAddress,err ", err, "token in", l.TokenIn)
                    return err
                }
                feeTokenSymbol, feeTokenName, feeTokenDecimal, err := erc20.GetSymbolNameDecimalByAddress(l.FeeToken, ltCtx.EthClient)
                if err != nil {
                    slog.Error("PaymentFee event", "erc20 getSymbolNameDecimalByAddress,err", err, "fee token", l.FeeToken)
                    return err
                }
                fee := decimal.NewFromBigInt(l.FeeAmount, -int32(feeTokenDecimal)).String()

                volume := ""
                if l.TokenIn.String() == ltCtx.USDTAddress {
                    volume = decimal.NewFromBigInt(l.AmountIn, -int32(tokenDecimal)).String()
                } else {
                    var swapPath []common.Address
                    if l.TokenIn.String() == ltCtx.EthAddress || l.TokenIn.String() == ltCtx.WethAddress {
                        swapPath = []common.Address{common.HexToAddress(ltCtx.WethAddress), common.HexToAddress(ltCtx.USDTAddress)}
                    } else {
                        swapPath = []common.Address{l.TokenIn, common.HexToAddress(ltCtx.WethAddress), common.HexToAddress(ltCtx.USDTAddress)}
                    }
                    outAmounts, err := uniswapv2router.GetAmountsOutByAddressAndBlockNumber(ltCtx.UniswapV2RouterAddress, new(big.Int).SetUint64(event.BlockNumber), l.AmountIn, ltCtx.EthClient, swapPath)
                    if err != nil {
                        slog.Error("PaymentFee event", "uniswapv2router getAmountsOutByAddressAndBlockNumber,err", err, "uniswapV2RouterAddress", ltCtx.UniswapV2RouterAddress, "block number", event.BlockNumber, "amount in", l.AmountIn, "path", swapPath)
                        return err
                    }
                    volume = decimal.NewFromBigInt(outAmounts[len(outAmounts)-1], -ltCtx.USDTDecimal).String()
                    // volume = _volume.Mul(decimal.NewFromBigInt(l.AmountIn, -int32(tokenDecimal))).String()
                }
                if volume == "0" {
                    volume = "0.00"
                }

                uw := tx.UserWallet
                userWallet, err := uw.WithContext(ctx).Where(uw.Address.Eq(l.Payer.String())).Take()
                if err != nil {
                    userWallet = &model.UserWallet{UID: 0}
                    slog.Error("PaymentFee event", "fail to get user wallet,err is: ", err)
                }

                eventTime := time.Unix(int64(block.Time), 0)
                userSwapTx := model.UserSwapTx{
                    UID:             userWallet.UID,
                    Address:         l.Payer.String(),
                    Tx:              event.TxHash.String(),
                    TokenSymbol:     tokenSymbol,
                    AmountIn:        decimal.NewFromBigInt(l.AmountIn, -int32(tokenDecimal)).String(),
                    TokenIn:         l.TokenIn.String(),
                    TransactionTime: eventTime,
                    Fee:             fee,
                    BlockNumber:     int32(event.BlockNumber),
                    ChainID:         int32(ltCtx.Chain.ChainId),
                    ChainName:       ltCtx.Chain.ChainName,
                    SwapType:        int32(l.SwapType.Int64()),
                    TokenName:       tokenName,
                    Decimal:         int32(tokenDecimal),
                    FeeTokenName:    feeTokenName,
                    FeeDecimal:      int32(feeTokenDecimal),
                    FeeTokenSymbol:  feeTokenSymbol,
                    CreatedAt:       time.Now().UTC(),
                }
                if err = tx.WithContext(ctx).UserSwapTx.Create(&userSwapTx); err != nil {
                    slog.Error("PaymentFee event", "fail to create user swap tx,err is: ", err)
                    return errors.Wrap(err, "fail to create user swap tx")
                }

                if err = tx.WithContext(ctx).UserTransaction.Create(&model.UserTransaction{
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
                    FeeTokenDecimal: int32(feeTokenDecimal),
                    IdentifyAddress: event.TxHash.String(),
                }); err != nil {
                    slog.Error("PaymentFee event", "fail to create user transaction,err is: ", err)
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
