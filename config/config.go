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
    ChainId                                             int    `yaml:"chain_id" json:"chain_id,omitempty"`
    ChainName                                           string `yaml:"chain_name" json:"chain_name,omitempty"`
    URL                                                 string `yaml:"url" json:"url,omitempty"`
    UniswapV2RouterAddress                              string `yaml:"uniswap_v2_router_address" json:"uniswap_v_2_router_address,omitempty"`
    StandardTokenFactory01Address                       string `yaml:"standard_token_factory01_address" json:"standard_token_factory_01_address,omitempty"`
    StandardTokenFactory01FeeSymbol                     string `yaml:"standard_token_factory01_fee_symbol" json:"standard_token_factory01_fee_symbol,omitempty"`
    StandardTokenFactory01FeeDecimal                    int32  `yaml:"standard_token_factory01_fee_decimal" json:"standard_token_factory01_fee_decimal,omitempty"`
    StandardTokenFactory01BlockNumber                   int64  `yaml:"standard_token_factory01_block_number" json:"standard_token_factory01_block_number,omitempty"`
    StandardTokenFactory01IsClose                       bool   `yaml:"standard_token_factory01_is_close" json:"standard_token_factory01_is_close,omitempty"`
    StandardTokenFactory01IsStartSavedNewestBlockNumber bool   `yaml:"standard_token_factory01_is_start_saved_newest_block_number"`
    DexpertUniversalRouterAddress                       string `yaml:"dexpert_universal_router_address" json:"dexpert_universal_router_address,omitempty"`
    DexpertUniversalRouterBlockNumber                   int64  `yaml:"dexpert_universal_router_block_number" json:"dexpert_universal_router_block_number,omitempty"`
    DexpertUniversalRouterUSDTAddress                   string `yaml:"dexpert_universal_router_usdt_address" json:"dexpert_universal_router_usdt_address,omitempty"`
    DexpertUniversalRouterWethAddress                   string `yaml:"dexpert_universal_router_weth_address" json:"dexpert_universal_router_weth_address,omitempty"`
    DexpertUniversalRouterEthAddress                    string `yaml:"dexpert_universal_router_eth_address" json:"dexpert_universal_router_eth_address,omitempty"`
    DexpertUniversalRouterUSDTDecimal                   int32  `yaml:"dexpert_universal_router_usdt_decimal" json:"dexpert_universal_router_usdt_decimal,omitempty"`
    DexpertUniversalRouterIsClose                       bool   `yaml:"dexpert_universal_router_is_close" json:"dexpert_universal_router_is_close,omitempty"`
    DexpertUniversalRouterIsStartSavedNewestBlockNumber bool   `yaml:"dexpert_universal_router_is_start_saved_newest_block_number"`
}
