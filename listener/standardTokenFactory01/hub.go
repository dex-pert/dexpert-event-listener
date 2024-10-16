package standardTokenFactory01

import (
    el "github.com/x1rh/event-listener"
    "dexpert-event-listener/abi"
    "github.com/pkg/errors"
    "fmt"
)

type Hub struct {
    EventListenerMap map[int]*el.EventListener
}

func NewHub(p *abi.Proxy) (*Hub, error) {
    eventListenerMap := make(map[int]*el.EventListener)
    for i, v := range p.Chains() {
        if v.StandardTokenFactory01.IsClose { // 关闭
            continue
        }
        var context *Context
        if p.Chains()[i].StandardTokenFactory01.Step > 0 {
            context = NewContext(&ContextParam{EthClient: p.WithChainID(v.ChainId).Client, ChainConfig: p.Chains()[i]}, WithStep(p.Chains()[i].StandardTokenFactory01.Step))
        } else {
            context = NewContext(&ContextParam{EthClient: p.WithChainID(v.ChainId).Client, ChainConfig: p.Chains()[i]})
        }
        eventListener, err := newStandardTokenFactory01EventListener(context)
        if err != nil {
            return nil, errors.Wrap(err, fmt.Sprintf("(chain_id:%d) new standard token factory01 event listener failed", v.ChainId))
        }
        eventListenerMap[v.ChainId] = eventListener
    }
    return &Hub{EventListenerMap: eventListenerMap}, nil
}
