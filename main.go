package main

import (
    "dexpert-event-listener/listener"
    "log/slog"

    "github.com/x1rh/event-listener/logger"
    "github.com/joho/godotenv"
    "dexpert-event-listener/appctx"
    "os"
    "dexpert-event-listener/config"
    "fmt"
    "gopkg.in/yaml.v3"
)

func main() {
    err := godotenv.Load(".env")
    if err != nil {
        panic(fmt.Sprintf("error loading .env file: %v", err))
    }
    file, err := os.ReadFile("./etc/config.yaml")
    if err != nil {
        panic(fmt.Sprintf("error reading config.yaml: %v", err))
    }
    yamlContent := os.ExpandEnv(string(file))
    var c config.Config
    err = yaml.Unmarshal([]byte(yamlContent), &c)
    if err != nil {
        panic(fmt.Sprintf("error parsing config.yaml: %v", err))
    }
    appCtx := appctx.NewContext(&c)
    logger.Init(slog.LevelInfo, false)
    el, err := listener.NewTokenFactoryEventListener(appCtx)
    if err != nil {
        slog.Error("fail to new a token factory listener", slog.Any("err", err))
    }
    el.Start()
}
