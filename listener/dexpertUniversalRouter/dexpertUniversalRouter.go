package dexpertUniversalRouter

import (
    el "github.com/x1rh/event-listener"
    "math/big"
    "github.com/pkg/errors"
)

func newDexpertUniversalRouterEventListener(ltCtx *Context) (*el.EventListener, error) {
    c := el.ChainConfig{
        ChainId:   ltCtx.Chain.ChainId,
        ChainName: ltCtx.Chain.ChainName,
        URL:       ltCtx.Chain.URL,
    }

    universalRouter, err := el.NewContract(ltCtx.DexpertUniversalRouterAddress, ltCtx.ABIStr, big.NewInt(ltCtx.BlockNumber), big.NewInt(ltCtx.Step))
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
