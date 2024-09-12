package listener

import (
    "dexpert-event-listener/listener/standardTokenFactory01"
    "dexpert-event-listener/listener/dexpertUniversalRouter"
    "dexpert-event-listener/config"
    "dexpert-event-listener/appctx"
    "log/slog"
    "fmt"
)

func Init(c *config.Config, a *appctx.Context) {
    if !c.IsCloseStandardTokenFactory01EventListener {
        hub, err := standardTokenFactory01.NewHub(a.AbiProxy)
        if err != nil {
            panic(err)
        }
        for i := range hub.EventListenerMap {
            slog.Info("=====================================================================================")
            slog.Info(fmt.Sprintf("chain_id:%d standard token factory01 event listener start...", i))
            slog.Info("=====================================================================================")
            go hub.EventListenerMap[i].Start()
        }
    }
    if !c.IsCloseDexpertUniversalRouterListener {
        hub, err := dexpertUniversalRouter.NewHub(a.AbiProxy)
        if err != nil {
            panic(err)
        }
        for i := range hub.EventListenerMap {
            slog.Info("=====================================================================================")
            slog.Info(fmt.Sprintf("chain_id:%d dexpert universal router event listener start...", i))
            slog.Info("=====================================================================================")
            go hub.EventListenerMap[i].Start()
        }
    }
}
