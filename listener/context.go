package listener

import (
    el "github.com/x1rh/event-listener"
    "dexpert-event-listener/config"
    "dexpert-event-listener/constant"
)

const (
    defaultStep = 100
)

type Context struct {
    Chain              el.ChainConfig
    TokenFactorAddress string
    ABIStr             string
    BlockNumber        int64
    Step               int64
    TokenSymbol        string
    LaunchFee          string
    TokenDecimal       int32
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

func NewContext(c *config.ChainConfig, opts ...Option) *Context {
    _options := options{
        Step: defaultStep,
    }
    for _, o := range opts {
        o.apply(&_options)
    }
    return &Context{
        Chain: el.ChainConfig{
            ChainId:   c.ChainId,
            ChainName: c.ChainName,
            URL:       c.URL,
        },
        BlockNumber:        c.BlockNumber,
        ABIStr:             constant.GetABIByChainId(c.ChainId),
        TokenFactorAddress: c.TokenFactorAddress,
        TokenSymbol:        c.TokenSymbol,
        LaunchFee:          c.LaunchFee,
        TokenDecimal:       c.TokenDecimal,
        Step:               _options.Step,
    }
}
