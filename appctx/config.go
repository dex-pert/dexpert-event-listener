package appctx

type Config struct {
    MySQL *MySQL
}

type MySQL struct {
    User     string
    Pass     string
    Host     string
    Port     string
    Database string
}
