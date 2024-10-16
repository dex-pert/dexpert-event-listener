package dexpertUniversalRouter

import (
    "github.com/ethereum/go-ethereum/ethclient"
    el "github.com/x1rh/event-listener"
    "dexpert-event-listener/config"
    "dexpert-event-listener/constant"
)

const (
    defaultStep = 100
)

type Context struct {
    EthClient                     *ethclient.Client
    Chain                         el.ChainConfig
    DexpertUniversalRouterAddress string
    Step                          int64
    ABIStr                        string
    BlockNumber                   int64
    USDTAddress                   string
    WethAddress                   string
    EthAddress                    string
    EthName                       string
    EthSymbol                     string
    EthDecimal                    int32
    USDTDecimal                   int32
    UniswapV2RouterAddress        string
    UniswapV2FactoryAddress       string
    IsStartSavedNewestBlockNumber bool
    IsUseNewestBlockNumber        bool
}

type ContextParam struct {
    EthClient   *ethclient.Client
    ChainConfig *config.ChainConfig
}

type options struct {
    Step int64
}

type Option interface {
    apply(o *options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) { f(o) }

func WithStep(p int64) Option {
    return optionFunc(func(o *options) {
        o.Step = p
    })
}

func NewContext(c *ContextParam, opts ...Option) *Context {
    _options := options{
        Step: defaultStep,
    }
    for _, o := range opts {
        o.apply(&_options)
    }
    return &Context{
        Chain: el.ChainConfig{
            ChainId:   c.ChainConfig.ChainId,
            ChainName: c.ChainConfig.ChainName,
            URL:       c.ChainConfig.URL,
        },
        UniswapV2RouterAddress:        c.ChainConfig.UniswapV2RouterAddress,
        UniswapV2FactoryAddress:       c.ChainConfig.UniswapV2FactoryAddress,
        EthClient:                     c.EthClient,
        Step:                          _options.Step,
        DexpertUniversalRouterAddress: c.ChainConfig.DexpertUniversalRouter.Address,
        ABIStr:                        constant.GetDexpertUniversalRouterABIByChainId(c.ChainConfig.ChainId),
        BlockNumber:                   c.ChainConfig.DexpertUniversalRouter.BlockNumber,
        USDTAddress:                   c.ChainConfig.DexpertUniversalRouter.USDTAddress,
        WethAddress:                   c.ChainConfig.DexpertUniversalRouter.WethAddress,
        EthAddress:                    c.ChainConfig.DexpertUniversalRouter.EthAddress,
        EthName:                       c.ChainConfig.DexpertUniversalRouter.EthName,
        EthSymbol:                     c.ChainConfig.DexpertUniversalRouter.EthSymbol,
        EthDecimal:                    c.ChainConfig.DexpertUniversalRouter.EthDecimal,
        USDTDecimal:                   c.ChainConfig.DexpertUniversalRouter.USDTDecimal,
        IsStartSavedNewestBlockNumber: c.ChainConfig.DexpertUniversalRouter.IsStartSavedNewestBlockNumber,
        IsUseNewestBlockNumber:        c.ChainConfig.DexpertUniversalRouter.IsUseNewestBlockNumber,
    }
}
