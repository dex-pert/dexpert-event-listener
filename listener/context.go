package listener

import (
    el "github.com/x1rh/event-listener"
)

const (
    defaultStep               = 100
    defaultBlockNumber        = 6641775
    defaultTokenFactorAddress = "0x822935C2240E6A0b5C96E3eA355446a83ed12C03"
    defaultAbiStr             = `[{"inputs":[{"internalType":"address","name":"factoryManager_","type":"address"},{"internalType":"address","name":"implementation_","type":"address"},{"internalType":"address","name":"feeTo_","type":"address"},{"internalType":"uint256","name":"maxFee_","type":"uint256"}],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[{"internalType":"uint256","name":"fee","type":"uint256"}],"name":"InsufficientFee","type":"error"},{"inputs":[{"internalType":"address","name":"implementation","type":"address"}],"name":"InvalidFactoryManager","type":"error"},{"inputs":[{"internalType":"uint256","name":"fee","type":"uint256"}],"name":"InvalidFee","type":"error"},{"inputs":[{"internalType":"address","name":"receiver","type":"address"}],"name":"InvalidFeeReceiver","type":"error"},{"inputs":[{"internalType":"address","name":"factoryManager","type":"address"}],"name":"InvalidImplementation","type":"error"},{"inputs":[{"internalType":"uint256","name":"level","type":"uint256"}],"name":"InvalidLevel","type":"error"},{"inputs":[{"internalType":"uint256","name":"maxFee","type":"uint256"}],"name":"InvalidMaxFee","type":"error"},{"inputs":[],"name":"OnlyOwner","type":"error"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"address","name":"newFeeTo","type":"address"}],"name":"FeeToUpdated","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"uint256","name":"level","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"newFee","type":"uint256"}],"name":"FeeUpdated","type":"event"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"uint256[]","name":"newLevels","type":"uint256[]"}],"name":"LevelsUpdated","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"previousOwner","type":"address"},{"indexed":true,"internalType":"address","name":"newOwner","type":"address"}],"name":"OwnershipTransferred","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"owner","type":"address"},{"indexed":true,"internalType":"address","name":"token","type":"address"},{"indexed":false,"internalType":"uint8","name":"tokenType","type":"uint8"},{"indexed":false,"internalType":"uint96","name":"tokenVersion","type":"uint96"},{"indexed":false,"internalType":"uint256","name":"level","type":"uint256"}],"name":"TokenCreated","type":"event"},{"anonymous":false,"inputs":[{"components":[{"internalType":"string","name":"description","type":"string"},{"internalType":"string","name":"logoLink","type":"string"},{"internalType":"string","name":"twitterLink","type":"string"},{"internalType":"string","name":"telegramLink","type":"string"},{"internalType":"string","name":"discordLink","type":"string"},{"internalType":"string","name":"websiteLink","type":"string"}],"indexed":false,"internalType":"struct TokenMetaData","name":"tokenMetaData","type":"tuple"}],"name":"TokenMetaDataUpdated","type":"event"},{"inputs":[],"name":"FACTORY_MANAGER","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"MAX_FEE","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"level","type":"uint256"},{"components":[{"internalType":"string","name":"name","type":"string"},{"internalType":"string","name":"symbol","type":"string"},{"internalType":"uint8","name":"decimals","type":"uint8"},{"internalType":"uint256","name":"totalSupply","type":"uint256"},{"internalType":"string","name":"description","type":"string"},{"internalType":"string","name":"logoLink","type":"string"},{"internalType":"string","name":"twitterLink","type":"string"},{"internalType":"string","name":"telegramLink","type":"string"},{"internalType":"string","name":"discordLink","type":"string"},{"internalType":"string","name":"websiteLink","type":"string"}],"internalType":"struct TokenInitializeParams","name":"tokenInitializeParams","type":"tuple"}],"name":"create","outputs":[{"internalType":"address","name":"token","type":"address"}],"stateMutability":"payable","type":"function"},{"inputs":[],"name":"feeTo","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"uint256","name":"","type":"uint256"}],"name":"fees","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"getLevels","outputs":[{"internalType":"uint256[]","name":"","type":"uint256[]"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"implementation","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"implementationVersion","outputs":[{"internalType":"uint96","name":"","type":"uint96"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"owner","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"renounceOwnership","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"level","type":"uint256"},{"internalType":"uint256","name":"fee","type":"uint256"}],"name":"setFee","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"feeTo_","type":"address"}],"name":"setFeeTo","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"implementation_","type":"address"}],"name":"setImplementation","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256[]","name":"_levels","type":"uint256[]"}],"name":"setLevels","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"newOwner","type":"address"}],"name":"transferOwnership","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"uint256","name":"level","type":"uint256"},{"internalType":"address","name":"token","type":"address"},{"components":[{"internalType":"string","name":"description","type":"string"},{"internalType":"string","name":"logoLink","type":"string"},{"internalType":"string","name":"twitterLink","type":"string"},{"internalType":"string","name":"telegramLink","type":"string"},{"internalType":"string","name":"discordLink","type":"string"},{"internalType":"string","name":"websiteLink","type":"string"}],"internalType":"struct TokenMetaData","name":"tokenMetaData_","type":"tuple"}],"name":"updateTokenMetaData","outputs":[],"stateMutability":"payable","type":"function"}]`
    defaultTokenSymbol        = "ETH"
    defaultLaunchFee          = "0.08"
    defaultTokenDecimal       = 18
)

type LTContext struct {
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
    Chain              el.ChainConfig
    TokenFactorAddress string
    ABIStr             string
    BlockNumber        int64
    Step               int64
    TokenSymbol        string
    LaunchFee          string
    TokenDecimal       int32
}

type Option interface {
    apply(o *options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) { f(o) }

func WithABIStr(p string) Option {
    return optionFunc(func(o *options) {
        o.ABIStr = p
    })
}

func WithBlockNumber(p int64) Option {
    return optionFunc(func(o *options) {
        o.BlockNumber = p
    })
}

func WithStep(p int64) Option {
    return optionFunc(func(o *options) {
        o.Step = p
    })
}

func WithTokenFactorAddress(p string) Option {
    return optionFunc(func(o *options) {
        o.TokenFactorAddress = p
    })
}

func WithTokenSymbol(p string) Option {
    return optionFunc(func(o *options) {
        o.TokenSymbol = p
    })
}

func WithLaunchFee(p string) Option {
    return optionFunc(func(o *options) {
        o.LaunchFee = p
    })
}

func WithTokenDecimal(p int32) Option {
    return optionFunc(func(o *options) {
        o.TokenDecimal = p
    })
}

func NewContext(c el.ChainConfig, opts ...Option) *LTContext {
    _options := options{
        Chain: el.ChainConfig{
            ChainId:   c.ChainId,
            ChainName: c.ChainName,
            URL:       c.URL,
        },
        Step:               defaultStep,
        BlockNumber:        defaultBlockNumber,
        ABIStr:             defaultAbiStr,
        TokenFactorAddress: defaultTokenFactorAddress,
        TokenSymbol:        defaultTokenSymbol,
        LaunchFee:          defaultLaunchFee,
        TokenDecimal:       defaultTokenDecimal,
    }
    for _, o := range opts {
        o.apply(&_options)
    }
    return &LTContext{
        Chain:              c,
        Step:               _options.Step,
        BlockNumber:        _options.BlockNumber,
        ABIStr:             _options.ABIStr,
        TokenFactorAddress: _options.TokenFactorAddress,
        TokenSymbol:        _options.TokenSymbol,
        LaunchFee:          _options.LaunchFee,
        TokenDecimal:       _options.TokenDecimal,
    }
}
