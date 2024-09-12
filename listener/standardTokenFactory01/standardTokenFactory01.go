package standardTokenFactory01

import (
    "math/big"
    el "github.com/x1rh/event-listener"

    "github.com/pkg/errors"
)

func newStandardTokenFactory01EventListener(ltCtx *Context) (*el.EventListener, error) {
    c := el.ChainConfig{
        ChainId:   ltCtx.Chain.ChainId,
        ChainName: ltCtx.Chain.ChainName,
        URL:       ltCtx.Chain.URL,
    }

    tokenFactory, err := el.NewContract(ltCtx.StandardTokenFactory01Address, ltCtx.ABIStr, big.NewInt(ltCtx.BlockNumber), big.NewInt(ltCtx.Step))
    if err != nil {
        panic(err)
    }
    tokenFactory.SetLogHandler(standardTokenFactory01EventLogHandler(ltCtx, tokenFactory))

    _el, err := el.New(
        c,
        el.WithClient(ltCtx.EthClient),
        el.WithContract(*tokenFactory),
    )
    if err != nil {
        return nil, errors.Wrap(err, "fail to new an EventListener object")
    }
    return _el, nil
}
