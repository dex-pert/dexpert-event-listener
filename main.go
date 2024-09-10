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
    "os/signal"
    "time"
    "syscall"
    "dexpert-event-listener/constant"
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
    c.Chains = appCtx.Chains
    logger.Init(slog.LevelInfo, false)

    if _, ok := c.Chains[constant.ChainIDEthereumSepolia]; !ok {
        panic(fmt.Sprintf("chain id %d not exist", constant.ChainIDEthereumSepolia))
    }
    ethereumSepoliaCtx := listener.NewContext(c.Chains[constant.ChainIDEthereumSepolia])
    ethereumSepoliaTokenFactoryEventListener, err := listener.NewTokenFactoryEventListener(ethereumSepoliaCtx)
    if err != nil {
        slog.Error("ethereum sepolia token factory event listener new failed", slog.Any("err", err))
    }
    go ethereumSepoliaTokenFactoryEventListener.Start()

    signalChan := make(chan os.Signal, 1)
    signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
    done := make(chan bool, 1)
    go func() {
        sig := <-signalChan
        slog.Info("Received signal: %s\n", sig)
        done <- true
    }()
    slog.Info("Dexpert event listener running. Press Ctrl+C to exit...")
    <-done
    slog.Info("Shutting down gracefully...")
    time.Sleep(2 * time.Second)
    slog.Info("ShuShutdown complete.")
}
