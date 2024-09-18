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
    if ltCtx.IsStartSavedNewestBlockNumber {
        ul := query.ListenerNewestBlocknumber
        record, err := ul.WithContext(context.Background()).Where(ul.ContractAddress.Eq(ltCtx.DexpertUniversalRouterAddress), ul.ChainID.Eq(int32(ltCtx.Chain.ChainId))).Take()
        if err != nil {
            slog.Error("newDexpertUniversalRouterEventListener", "failed to select ListenerNewestBlocknumber", err.Error(),
                "contract address", ltCtx.DexpertUniversalRouterAddress, "chain_id", ltCtx.Chain.ChainId)
            if !errors.Is(err, gorm.ErrRecordNotFound) {
                return nil, errors.Wrap(err, "fail to get newest user swap tx")
            }
        } else {
            if record.BlockNumber > int32(blockNumber) {
                blockNumber = int64(record.BlockNumber)
            }
        }
    }
    if ltCtx.IsUseNewestBlockNumber {
        _newestBlockNumber, err := ltCtx.EthClient.BlockNumber(context.Background())
        if err != nil {
            slog.Error("newDexpertUniversalRouterEventListener", "failed to get newest block number", err, "chain_id", ltCtx.Chain.ChainId)
            return nil, errors.Wrap(err, "failed to get newest block number")
        }
        blockNumber = int64(_newestBlockNumber)
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
