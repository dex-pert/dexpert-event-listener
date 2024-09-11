package listener

import (
    el "github.com/x1rh/event-listener"
    "github.com/ethereum/go-ethereum/ethclient"
    "math/big"
    "github.com/pkg/errors"
    "github.com/ethereum/go-ethereum/common"
    "context"
    "log/slog"
    "dexpert-event-listener/gorm/model"
    "time"
    "dexpert-event-listener/constant"
    "dexpert-event-listener/gorm/query"
    "dexpert-event-listener/abi/erc20"
    "github.com/shopspring/decimal"
    "dexpert-event-listener/abi/uniswapv2router"
)

func NewUniversalRouterEventListener(ltCtx *Context) (*el.EventListener, error) {
    c := el.ChainConfig{
        ChainId:   ltCtx.Chain.ChainId,
        ChainName: ltCtx.Chain.ChainName,
        URL:       ltCtx.Chain.URL,
    }

    client, err := ethclient.Dial(c.URL)
    if err != nil {
        panic(err)
    }

    universalRouter, err := el.NewContract(ltCtx.UniversalRouterAddress, ltCtx.UniversalRouterABIStr, big.NewInt(ltCtx.UniversalRouterBlockNumber), big.NewInt(ltCtx.UniversalRouterStep))
    if err != nil {
        panic(err)
    }
    universalRouter.SetLogHandler(universalRouterLogHandler(ltCtx, universalRouter))

    _el, err := el.New(
        c,
        el.WithClient(client),
        el.WithContract(*universalRouter),
    )
    if err != nil {
        return nil, errors.Wrap(err, "fail to new an EventListener object")
    }
    return _el, nil
}

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

func universalRouterLogHandler(ltCtx *Context, c *el.Contract) el.LogHandleFunc {
    return func(ctx context.Context, event *el.Event) error {
        switch event.Name {
        case "Payment", "PaymentFee":
            var l LogPaymentFee
            if err := c.Abi.UnpackIntoInterface(&l, event.Name, event.Data); err != nil {
                return errors.Wrap(err, "fail to unpack log")
            }
            slog.Info("PaymentFee event", slog.Any("event", l))

            _ = query.Q.Transaction(func(tx *query.Query) error {
                abiProxy := ltCtx.AbiProxy.WithChainID(ltCtx.Chain.ChainId)
                block, err := abiProxy.BlockByNumber(ctx, new(big.Int).SetUint64(event.BlockNumber))
                if err != nil {
                    slog.Error("TokenCreated event", "fail to get block,err is: ", err)
                    return err
                }
                tokenSymbol, tokenName, tokenDecimal, err := erc20.GetSymbolNameDecimalByAddress(l.TokenIn, abiProxy.Client)
                if err != nil {
                    slog.Error("PaymentFee event", "fail to get block,err is: ", err)
                    return err
                }
                feeTokenSymbol, feeTokenName, feeTokenDecimal, err := erc20.GetSymbolNameDecimalByAddress(l.FeeToken, abiProxy.Client)
                if err != nil {
                    slog.Error("PaymentFee event", "erc20 getSymbolNameDecimalByAddress,err is: ", err)
                    return err
                }
                fee := decimal.NewFromBigInt(l.FeeAmount, -int32(feeTokenDecimal)).String()

                volume := ""
                if l.TokenIn.String() == ltCtx.UniversalRouterUSDTAddress {
                    volume = decimal.NewFromBigInt(l.FeeAmount, -int32(feeTokenDecimal)).StringFixed(int32(feeTokenDecimal))
                } else {
                    var swapPath []common.Address
                    if l.TokenIn.String() == ltCtx.UniversalRouterEthAddress || l.TokenIn.String() == ltCtx.UniversalRouterWethAddress {
                        swapPath = []common.Address{common.HexToAddress(ltCtx.UniversalRouterWethAddress), common.HexToAddress(ltCtx.UniversalRouterUSDTAddress)}
                    } else {
                        swapPath = []common.Address{l.TokenIn, common.HexToAddress(ltCtx.UniversalRouterWethAddress), common.HexToAddress(ltCtx.UniversalRouterUSDTAddress)}
                    }
                    outAmounts, err := uniswapv2router.GetAmountsOutByAddressAndBlockNumber(ltCtx.UniswapV2RouterAddress, new(big.Int).SetUint64(event.BlockNumber), l.AmountIn, abiProxy.Client, swapPath)
                    if err != nil {
                        slog.Error("PaymentFee event", "uniswapv2router getAmountsOutByAddressAndBlockNumber,err is: ", err)
                        return err
                    }
                    _volume := decimal.NewFromBigInt(outAmounts[len(outAmounts)-1], -ltCtx.UniversalRouterUSDTDecimal)
                    volume = _volume.Mul(decimal.NewFromBigInt(l.AmountIn, -int32(tokenDecimal))).String()
                }
                if volume == "0" {
                    volume = "0.00"
                }

                uw := tx.UserWallet
                userWallet, err := uw.WithContext(ctx).Where(uw.Address.Eq(l.Payer.String()), uw.ChainID.Eq(int32(ltCtx.Chain.ChainId))).Take()
                if err != nil {
                    userWallet = &model.UserWallet{UID: 0}
                    slog.Error("PaymentFee event", "fail to get user wallet,err is: ", err)
                }

                eventTime := time.Unix(int64(block.Time()), 0).UTC()
                userSwapTx := model.UserSwapTx{
                    UID:             userWallet.UID,
                    Address:         l.Payer.String(),
                    Tx:              "",
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
