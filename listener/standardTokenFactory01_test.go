package listener

import (
    "testing"
    "github.com/x1rh/event-listener/logger"
    "log/slog"
    "dexpert-event-listener/gorm/model"
    "dexpert-event-listener/gorm/query"
    "time"
    "context"
)

func TestEventListen(t *testing.T) {
    logger.Init(slog.LevelInfo, false)
    el, err := NewStandardTokenFactory01EventListener(nil)
    if err != nil {
        t.Fatal(err)
    }
    el.Start()

    ult := query.UserLaunchTx
    if err = ult.WithContext(context.Background()).Create(&model.UserLaunchTx{
        UID:             1,
        ContractAddress: "123",
        ChainID:         1,
        ChainName:       "eth",
        PairAddress:     "123",
        Fee:             "123",
        FeeTokenSymbol:  "123",
        Timestamp:       time.Now().UTC(),
        TypeName:        "123",
    }); err != nil {
        slog.Error("user launch tx create error", slog.String("err", err.Error()))
    }
}

func TestGetLen(t *testing.T) {
    s := "0x000000000000000000000000d3952283b16c813c"
    t.Log(len(s))
}
