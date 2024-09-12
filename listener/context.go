package listener

import (
    el "github.com/x1rh/event-listener"
    "dexpert-event-listener/config"
    "dexpert-event-listener/constant"
    "dexpert-event-listener/abi"
)

const (
    defaultStep = 100
)

type Context struct {
    AbiProxy                          *abi.Proxy
    Chain                             el.ChainConfig
    StandardTokenFactory01Address     string
    StandardTokenFactory01Step        int64
    StandardTokenFactory01FeeSymbol   string
    StandardTokenFactory01FeeDecimal  int32
    StandardTokenFactory01BlockNumber int64
    StandardTokenFactory01ABIStr      string

    DexpertUniversalRouterAddress     string
    DexpertUniversalRouterStep        int64
    DexpertUniversalRouterABIStr      string
    DexpertUniversalRouterBlockNumber int64
    DexpertUniversalRouterUSDTAddress string
    DexpertUniversalRouterWethAddress string
    DexpertUniversalRouterEthAddress  string
    DexpertUniversalRouterUSDTDecimal int32

    UniswapV2RouterAddress string
}

type ContextParam struct {
    AbiProxy    *abi.Proxy
    ChainConfig *config.ChainConfig
}

type options struct {
    StandardTokenFactory01Step int64
    UniversalRouterStep        int64
}

type Option interface {
    apply(o *options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) { f(o) }

func WithStandardTokenFactory01Step(p int64) Option {
    return optionFunc(func(o *options) {
        o.StandardTokenFactory01Step = p
    })
}

func WithUniversalRouterStep(p int64) Option {
    return optionFunc(func(o *options) {
        o.UniversalRouterStep = p
    })
}

func NewContext(c *ContextParam, opts ...Option) *Context {
    _options := options{
        StandardTokenFactory01Step: defaultStep,
        UniversalRouterStep:        defaultStep,
    }
    for _, o := range opts {
        o.apply(&_options)
    }
    return &Context{
        UniswapV2RouterAddress: c.ChainConfig.UniswapV2RouterAddress,
        Chain: el.ChainConfig{
            ChainId:   c.ChainConfig.ChainId,
            ChainName: c.ChainConfig.ChainName,
            URL:       c.ChainConfig.URL,
        },
        AbiProxy:                          c.AbiProxy,
        StandardTokenFactory01Step:        _options.StandardTokenFactory01Step,
        StandardTokenFactory01BlockNumber: c.ChainConfig.StandardTokenFactory01BlockNumber,
        StandardTokenFactory01ABIStr:      constant.GetStandardTokenFactory01ABIByChainId(c.ChainConfig.ChainId),
        StandardTokenFactory01Address:     c.ChainConfig.StandardTokenFactory01Address,
        StandardTokenFactory01FeeSymbol:   c.ChainConfig.StandardTokenFactory01FeeSymbol,
        StandardTokenFactory01FeeDecimal:  c.ChainConfig.StandardTokenFactory01FeeDecimal,

        DexpertUniversalRouterStep:        _options.UniversalRouterStep,
        DexpertUniversalRouterAddress:     c.ChainConfig.DexpertUniversalRouterAddress,
        DexpertUniversalRouterABIStr:      constant.GetDexpertUniversalRouterABIByChainId(c.ChainConfig.ChainId),
        DexpertUniversalRouterBlockNumber: c.ChainConfig.DexpertUniversalRouterBlockNumber,
        DexpertUniversalRouterUSDTAddress: c.ChainConfig.DexpertUniversalRouterUSDTAddress,
        DexpertUniversalRouterWethAddress: c.ChainConfig.DexpertUniversalRouterWethAddress,
        DexpertUniversalRouterEthAddress:  c.ChainConfig.DexpertUniversalRouterEthAddress,
        DexpertUniversalRouterUSDTDecimal: c.ChainConfig.DexpertUniversalRouterUSDTDecimal,
    }
}
