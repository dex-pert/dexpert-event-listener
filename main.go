package main

import (
    "dexpert-event-listener/listener"
    "log/slog"

    "github.com/x1rh/event-listener/logger"
    "github.com/joho/godotenv"
    "log"
    "dexpert-event-listener/appctx"
    "os"
    "dexpert-event-listener/config"
)

func main() {
    err := godotenv.Load(".env")
    if err != nil {
        log.Println(err)
    }
    chains := config.GetChainConfigs()
    appCtx := appctx.NewContext(&config.Config{
        MySQL: &config.MySQL{
            User:     os.Getenv("MYSQL_USER"),
            Pass:     os.Getenv("MYSQL_PASS"),
            Host:     os.Getenv("MYSQL_HOST"),
            Port:     os.Getenv("MYSQL_PORT"),
            Database: os.Getenv("DB_NAME"),
        },
        Chains: chains,
    })
    logger.Init(slog.LevelInfo, false)
    el, err := listener.NewTokenFactoryEventListener(appCtx)
    if err != nil {
        slog.Error("fail to new a token factory listener", slog.Any("err", err))
    }
    el.Start()
}
