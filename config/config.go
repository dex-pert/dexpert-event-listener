package config

type Config struct {
    IsCloseStandardTokenFactory01EventListener bool          `yaml:"is_close_standard_token_factory01_listener"`
    IsCloseDexpertUniversalRouterListener      bool          `yaml:"is_close_dexpert_universal_router_listener"`
    ChainConfig                                []ChainConfig `yaml:"chain_config"`
    MySQL                                      *MySQL        `yaml:"mysql"`
    Chains                                     map[int]*ChainConfig
}

type MySQL struct {
    User string `yaml:"user"`
    Pass string `yaml:"pass"`
    Host string `yaml:"host"`
    Port string `yaml:"port"`
    DB   string `yaml:"db"`
}

type ChainConfig struct {
    ChainId                 int                    `yaml:"chain_id" json:"chain_id,omitempty"`
    ChainName               string                 `yaml:"chain_name" json:"chain_name,omitempty"`
    URL                     string                 `yaml:"url" json:"url,omitempty"`
    UniswapV2RouterAddress  string                 `yaml:"uniswap_v2_router_address" json:"uniswap_v_2_router_address,omitempty"`
    UniswapV2FactoryAddress string                 `yaml:"uniswap_v2_factory_address"`
    StandardTokenFactory01  StandardTokenFactory01 `yaml:"standard_token_factory01"`
    DexpertUniversalRouter  DexpertUniversalRouter `yaml:"dexpert_universal_router"`
}

type StandardTokenFactory01 struct {
    Step                          int64  `yaml:"step"`
    Address                       string `yaml:"address"`
    FeeSymbol                     string `yaml:"fee_symbol"`
    FeeDecimal                    int32  `yaml:"fee_decimal"`
    BlockNumber                   int64  `yaml:"block_number"`
    IsClose                       bool   `yaml:"is_close"`
    IsStartSavedNewestBlockNumber bool   `yaml:"is_start_saved_newest_block_number"`
    IsUseNewestBlockNumber        bool   `yaml:"is_use_newest_block_number"`
}

type DexpertUniversalRouter struct {
    Step                          int64  `yaml:"step"`
    Address                       string `yaml:"address"`
    BlockNumber                   int64  `yaml:"block_number"`
    USDTAddress                   string `yaml:"usdt_address"`
    WethAddress                   string `yaml:"weth_address"`
    EthAddress                    string `yaml:"eth_address"`
    EthName                       string `yaml:"eth_name"`
    EthSymbol                     string `yaml:"eth_symbol"`
    EthDecimal                    int32  `yaml:"eth_decimal"`
    USDTDecimal                   int32  `yaml:"usdt_decimal"`
    IsClose                       bool   `yaml:"is_close"`
    IsStartSavedNewestBlockNumber bool   `yaml:"is_start_saved_newest_block_number"`
    IsUseNewestBlockNumber        bool   `yaml:"is_use_newest_block_number"`
}
