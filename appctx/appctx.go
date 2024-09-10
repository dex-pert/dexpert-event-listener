package appctx

import (
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "fmt"
    "gorm.io/gorm/logger"
    "time"
    "log/slog"
    "dexpert-event-listener/gorm/query"
    "dexpert-event-listener/config"
    "dexpert-event-listener/abi/tokenfactory"
)

type Context struct {
    ChainConfig       *config.Chain
    TokenFactoryProxy *tokenfactory.Proxy
}

func NewContext(c *config.Config) *Context {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", c.MySQL.User, c.MySQL.Pass, c.MySQL.Host, c.MySQL.Port, c.MySQL.DB)
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
        SkipDefaultTransaction: true,
        PrepareStmt:            true,
        Logger:                 settingLogConfig(),
    })
    if err != nil {
        panic(err)
    }
    query.SetDefault(db)

    chains := make(map[int]*config.Chain, len(c.Chains))
    for _, v := range c.Chains {
        chains[v.ChainId] = &config.Chain{
            ChainId:   v.ChainId,
            ChainName: v.ChainName,
            URL:       v.URL,
        }
    }
    tokenFactoryProxy, err := tokenfactory.NewProxy(chains)
    if err != nil {
        panic(err)
    }
    return &Context{
        TokenFactoryProxy: tokenFactoryProxy,
    }
}

type Writer struct {
}

func (w Writer) Printf(format string, args ...interface{}) {
    slog.Info(format, args)
}

func settingLogConfig() logger.Interface {
    newLogger := logger.New(
        Writer{},
        logger.Config{
            SlowThreshold:             200 * time.Millisecond, // Slow SQL threshold
            LogLevel:                  logger.Info,            // Log level
            IgnoreRecordNotFoundError: true,                   // Ignore ErrRecordNotFound error for logger
            Colorful:                  true,                   // Disable color
        },
    )
    return newLogger
}
