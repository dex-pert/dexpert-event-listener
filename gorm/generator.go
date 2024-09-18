package main

import (
    "gorm.io/gen"
    "gorm.io/gorm"
    "gorm.io/driver/mysql"
    "github.com/joho/godotenv"
    "fmt"
    "os"
)

func main() {
    g := gen.NewGenerator(gen.Config{
        OutPath: "./query",
        // Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
        Mode: gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
    })

    err := godotenv.Load("../.env")
    if err != nil {
        panic(err)
    }

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
        os.Getenv("MYSQL_USER"),
        os.Getenv("MYSQL_PASS"),
        os.Getenv("MYSQL_HOST"),
        os.Getenv("MYSQL_PORT"),
        os.Getenv("DB_NAME"),
    )
    db, _ := gorm.Open(mysql.Open(dsn))
    g.UseDB(db)
    // Generate basic type-safe DAO API for struct `model.User` following conventions
    g.ApplyBasic(g.GenerateModel("user_transaction"),
        g.GenerateModel("user_launch_tx"),
        g.GenerateModel("user_wallet"),
        g.GenerateModel("user_swap_tx"),
        g.GenerateModel("listener_newest_blocknumber"),
    )
    // Generate the code
    g.Execute()
}
