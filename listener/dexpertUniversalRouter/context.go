package dexpertUniversalRouter

import (
    "github.com/ethereum/go-ethereum/ethclient"
    el "github.com/x1rh/event-listener"
    "dexpert-event-listener/config"
    "dexpert-event-listener/constant"
)

const (
    defaultStep = 10
)

type Context struct {
    EthClient                                           *ethclient.Client
    Chain                                               el.ChainConfig
    DexpertUniversalRouterAddress                       string
    Step                                                int64
    ABIStr                                              string
    BlockNumber                                         int64
    USDTAddress                                         string
    WethAddress                                         string
    EthAddress                                          string
    EthName                                             string
    EthSymbol                                           string
    EthDecimal                                          int32
    USDTDecimal                                         int32
    UniswapV2RouterAddress                              string
    DexpertUniversalRouterIsStartSavedNewestBlockNumber bool
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
        EthClient:                     c.EthClient,
        Step:                          _options.Step,
        DexpertUniversalRouterAddress: c.ChainConfig.DexpertUniversalRouterAddress,
        ABIStr:                        constant.GetDexpertUniversalRouterABIByChainId(c.ChainConfig.ChainId),
        BlockNumber:                   c.ChainConfig.DexpertUniversalRouterBlockNumber,
        USDTAddress:                   c.ChainConfig.DexpertUniversalRouterUSDTAddress,
        WethAddress:                   c.ChainConfig.DexpertUniversalRouterWethAddress,
        EthAddress:                    c.ChainConfig.DexpertUniversalRouterEthAddress,
        EthName:                       c.ChainConfig.DexpertUniversalRouterEthName,
        EthSymbol:                     c.ChainConfig.DexpertUniversalRouterEthSymbol,
        EthDecimal:                    c.ChainConfig.DexpertUniversalRouterEthDecimal,
        USDTDecimal:                   c.ChainConfig.DexpertUniversalRouterUSDTDecimal,
        DexpertUniversalRouterIsStartSavedNewestBlockNumber: c.ChainConfig.DexpertUniversalRouterIsStartSavedNewestBlockNumber,
    }
}
