package standardTokenFactory01

import (
    "math/big"
    el "github.com/x1rh/event-listener"

    "github.com/pkg/errors"
    "dexpert-event-listener/gorm/query"
    "gorm.io/gorm"
    "context"
    "log/slog"
)

func newStandardTokenFactory01EventListener(ltCtx *Context) (*el.EventListener, error) {
    c := el.ChainConfig{
        ChainId:   ltCtx.Chain.ChainId,
        ChainName: ltCtx.Chain.ChainName,
        URL:       ltCtx.Chain.URL,
    }

    blockNumber := ltCtx.BlockNumber
    if ltCtx.StandardTokenFactory01IsStartSavedNewestBlockNumber {
        ul := query.UserLaunchTx
        newestUserLaunchTx, err := ul.WithContext(context.Background()).Order(ul.BlockNumber.Desc()).Take()
        if err != nil {
            slog.Error("newStandardTokenFactory01EventListener", "err", err.Error())
            if !errors.Is(err, gorm.ErrRecordNotFound) {
                return nil, errors.Wrap(err, "fail to get newest user launch tx")
            }
        } else {
            if newestUserLaunchTx.BlockNumber > int32(blockNumber) {
                blockNumber = int64(newestUserLaunchTx.BlockNumber)
            }
        }
    }

    tokenFactory, err := el.NewContract(ltCtx.StandardTokenFactory01Address, ltCtx.ABIStr, big.NewInt(blockNumber), big.NewInt(ltCtx.Step))
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
