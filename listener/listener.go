package listener

import (
    "dexpert-event-listener/listener/standardTokenFactory01"
    "dexpert-event-listener/listener/dexpertUniversalRouter"
    "dexpert-event-listener/config"
    "dexpert-event-listener/appctx"
)

func Init(c *config.Config, a *appctx.Context) {
    if !c.IsCloseStandardTokenFactory01EventListener {
        hub, err := standardTokenFactory01.NewHub(a.AbiProxy)
        if err != nil {
            panic(err)
        }
        for i := range hub.EventListenerMap {
            go hub.EventListenerMap[i].Start()
        }
    }
    if !c.IsCloseUniversalRouterListener {
        hub, err := dexpertUniversalRouter.NewHub(a.AbiProxy)
        if err != nil {
            panic(err)
        }
        for i := range hub.EventListenerMap {
            go hub.EventListenerMap[i].Start()
        }
    }
}
