package listener

import (
    "dexpert-event-listener/config"
    el "github.com/x1rh/event-listener"
    "dexpert-event-listener/constant"
    "github.com/pkg/errors"
    "fmt"
)

type IListener interface {
    Start()
}

func InitListeners(c *config.Config) ([]IListener, error) {
    if _, ok := c.Chains[constant.ChainIDEthereumSepolia]; !ok {
        return nil, errors.New(fmt.Sprintf("chain id %d not exist", constant.ChainIDEthereumSepolia))
    }
    context1 := NewContext(el.ChainConfig{
        ChainId:   c.Chains[constant.ChainIDEthereumSepolia].ChainId,
        ChainName: c.Chains[constant.ChainIDEthereumSepolia].ChainName,
        URL:       c.Chains[constant.ChainIDEthereumSepolia].URL,
    })
    ethereumSepoliaTokenFactoryEventListener, err := NewTokenFactoryEventListener(context1)
    if err != nil {
        return nil, err
    }

    return []IListener{
        ethereumSepoliaTokenFactoryEventListener,
    }, nil
}
