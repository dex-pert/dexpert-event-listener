package dexpertUniversalRouter

import (
    el "github.com/x1rh/event-listener"
    "math/big"
    "github.com/pkg/errors"
    "dexpert-event-listener/gorm/query"
    "gorm.io/gorm"
    "context"
    "log/slog"
)

func newDexpertUniversalRouterEventListener(ltCtx *Context) (*el.EventListener, error) {
    c := el.ChainConfig{
        ChainId:   ltCtx.Chain.ChainId,
        ChainName: ltCtx.Chain.ChainName,
        URL:       ltCtx.Chain.URL,
    }

    blockNumber := ltCtx.BlockNumber
    if ltCtx.DexpertUniversalRouterIsStartSavedNewestBlockNumber {
        ul := query.UserSwapTx
        newestUserSwapTx, err := ul.WithContext(context.Background()).Order(ul.BlockNumber.Desc()).Take()
        if err != nil {
            slog.Error("newDexpertUniversalRouterEventListener", "err", err.Error())
            if !errors.Is(err, gorm.ErrRecordNotFound) {
                return nil, errors.Wrap(err, "fail to get newest user swap tx")
            }
        } else {
            if newestUserSwapTx.BlockNumber > int32(blockNumber) {
                blockNumber = int64(newestUserSwapTx.BlockNumber)
            }
        }
    }

    universalRouter, err := el.NewContract(ltCtx.DexpertUniversalRouterAddress, ltCtx.ABIStr, big.NewInt(blockNumber), big.NewInt(ltCtx.Step))
    if err != nil {
        panic(err)
    }
    universalRouter.SetLogHandler(dexpertUniversalRouterLogHandler(ltCtx, universalRouter))

    _el, err := el.New(
        c,
        el.WithClient(ltCtx.EthClient),
        el.WithContract(*universalRouter),
    )
    if err != nil {
        return nil, errors.Wrap(err, "fail to new an EventListener object")
    }
    return _el, nil
}
