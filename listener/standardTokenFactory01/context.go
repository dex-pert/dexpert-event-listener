package standardTokenFactory01

import (
    el "github.com/x1rh/event-listener"
    "dexpert-event-listener/config"
    "dexpert-event-listener/constant"
    "github.com/ethereum/go-ethereum/ethclient"
)

const (
    defaultStep = 100
)

type Context struct {
    EthClient                     *ethclient.Client
    Chain                         el.ChainConfig
    UniswapV2RouterAddress        string
    StandardTokenFactory01Address string
    Step                          int64
    FeeSymbol                     string
    FeeDecimal                    int32
    BlockNumber                   int64
    ABIStr                        string
    IsStartSavedNewestBlockNumber bool
    IsUseNewestBlockNumber        bool
    UniswapV2FactoryAddress       string
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
        BlockNumber:                   c.ChainConfig.StandardTokenFactory01.BlockNumber,
        ABIStr:                        constant.GetStandardTokenFactory01ABIByChainId(c.ChainConfig.ChainId),
        StandardTokenFactory01Address: c.ChainConfig.StandardTokenFactory01.Address,
        FeeSymbol:                     c.ChainConfig.StandardTokenFactory01.FeeSymbol,
        FeeDecimal:                    c.ChainConfig.StandardTokenFactory01.FeeDecimal,
        IsStartSavedNewestBlockNumber: c.ChainConfig.StandardTokenFactory01.IsStartSavedNewestBlockNumber,
        IsUseNewestBlockNumber:        c.ChainConfig.StandardTokenFactory01.IsUseNewestBlockNumber,
    }
}
