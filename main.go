package main

import (
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
    "dexpert-event-listener/listener"
    "path/filepath"
)

func main() {
    dir, err := os.Getwd()
    if err != nil {
        slog.Error(fmt.Sprintf("os getwd failed: %v", err))
    }
    envPath := filepath.Join(dir, ".env")
    slog.Info("filepath.Join", "envPath", envPath)
    err = godotenv.Load(envPath)
    if err != nil {
        slog.Error(fmt.Sprintf("error loading .env file: %v", err))
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

    // 启动监听器
    listener.Init(&c, appCtx)

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
