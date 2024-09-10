package config

import "dexpert-event-listener/constant"

func GetChainConfigs() []*Chain {
    return []*Chain{
        {
            ChainId:   constant.ChainIDEthereumSepolia,
            ChainName: constant.ChainNameEthereumSepolia,
            URL:       constant.ChainURLEthereumSepolia,
        },
    }
}
