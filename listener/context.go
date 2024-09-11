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
    AbiProxy                   *abi.Proxy
    Chain                      el.ChainConfig
    TokenFactoryStep           int64
    TokenFactoryTokenSymbol    string
    TokenFactoryLaunchFee      string
    TokenFactoryTokenDecimal   int32
    TokenFactoryBlockNumber    int64
    TokenFactoryAddress        string
    TokenFactoryABIStr         string
    UniversalRouterStep        int64
    UniversalRouterAddress     string
    UniversalRouterABIStr      string
    UniversalRouterBlockNumber int64
    UniversalRouterUSDTAddress string
    UniversalRouterWethAddress string
    UniswapV2RouterAddress     string
    UniversalRouterEthAddress  string
    UniversalRouterUSDTDecimal int32
}

type ContextParam struct {
    AbiProxy    *abi.Proxy
    ChainConfig *config.ChainConfig
}

type options struct {
    TokenFactoryStep    int64
    UniversalRouterStep int64
}

type Option interface {
    apply(o *options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) { f(o) }

func WithTokenFactoryStep(p int64) Option {
    return optionFunc(func(o *options) {
        o.TokenFactoryStep = p
    })
}

func WithUniversalRouterStep(p int64) Option {
    return optionFunc(func(o *options) {
        o.UniversalRouterStep = p
    })
}

func NewContext(c *ContextParam, opts ...Option) *Context {
    _options := options{
        TokenFactoryStep:    defaultStep,
        UniversalRouterStep: defaultStep,
    }
    for _, o := range opts {
        o.apply(&_options)
    }
    return &Context{
        TokenFactoryStep:       _options.TokenFactoryStep,
        UniversalRouterStep:    _options.UniversalRouterStep,
        UniswapV2RouterAddress: c.ChainConfig.UniswapV2RouterAddress,
        Chain: el.ChainConfig{
            ChainId:   c.ChainConfig.ChainId,
            ChainName: c.ChainConfig.ChainName,
            URL:       c.ChainConfig.URL,
        },
        AbiProxy:                   c.AbiProxy,
        TokenFactoryBlockNumber:    c.ChainConfig.TokenFactoryBlockNumber,
        TokenFactoryABIStr:         constant.GetTokenFactoryABIByChainId(c.ChainConfig.ChainId),
        TokenFactoryAddress:        c.ChainConfig.TokenFactoryAddress,
        TokenFactoryTokenSymbol:    c.ChainConfig.TokenFactoryTokenSymbol,
        TokenFactoryLaunchFee:      c.ChainConfig.TokenFactoryLaunchFee,
        TokenFactoryTokenDecimal:   c.ChainConfig.TokenFactoryTokenDecimal,
        UniversalRouterAddress:     c.ChainConfig.UniversalRouterAddress,
        UniversalRouterABIStr:      constant.GetUniversalRouterABIByChainId(c.ChainConfig.ChainId),
        UniversalRouterBlockNumber: c.ChainConfig.UniversalRouterBlockNumber,
        UniversalRouterUSDTAddress: c.ChainConfig.UniversalRouterUSDTAddress,
        UniversalRouterWethAddress: c.ChainConfig.UniversalRouterWethAddress,
        UniversalRouterEthAddress:  c.ChainConfig.UniversalRouterEthAddress,
        UniversalRouterUSDTDecimal: c.ChainConfig.UniversalRouterUSDTDecimal,
    }
}
