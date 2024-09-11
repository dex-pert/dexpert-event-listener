package config

type Config struct {
    IsCloseTokenFactoryListener    bool          `yaml:"is_close_token_factory_listener"`
    IsCloseUniversalRouterListener bool          `yaml:"is_close_universal_router_listener"`
    ChainConfig                    []ChainConfig `yaml:"chain_config"`
    MySQL                          *MySQL        `yaml:"mysql"`
    Chains                         map[int]*ChainConfig
}

type MySQL struct {
    User string `yaml:"user"`
    Pass string `yaml:"pass"`
    Host string `yaml:"host"`
    Port string `yaml:"port"`
    DB   string `yaml:"db"`
}

type ChainConfig struct {
    ChainId                    int    `yaml:"chain_id"`
    ChainName                  string `yaml:"chain_name"`
    URL                        string `yaml:"url"`
    UniswapV2RouterAddress     string `yaml:"uniswap_v2_router_address"`
    TokenFactoryTokenSymbol    string `yaml:"token_factory_token_symbol"`
    TokenFactoryLaunchFee      string `yaml:"token_factory_launch_fee"`
    TokenFactoryTokenDecimal   int32  `yaml:"token_factory_token_decimal"`
    TokenFactoryBlockNumber    int64  `yaml:"token_factory_block_number"`
    TokenFactoryAddress        string `yaml:"token_factory_address"`
    UniversalRouterAddress     string `yaml:"universal_router_address"`
    UniversalRouterBlockNumber int64  `yaml:"universal_router_block_number"`
    UniversalRouterUSDTAddress string `yaml:"universal_router_usdt_address"`
    UniversalRouterWethAddress string `yaml:"universal_router_weth_address"`
    UniversalRouterEthAddress  string `yaml:"universal_router_eth_address"`
    UniversalRouterUSDTDecimal int32  `yaml:"universal_router_usdt_decimal"`
}
