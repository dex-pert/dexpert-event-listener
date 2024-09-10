package config

type Config struct {
    Chains []*Chain
    MySQL  *MySQL
}

type MySQL struct {
    User     string
    Pass     string
    Host     string
    Port     string
    Database string
}

type Chain struct {
    ChainId   int
    ChainName string
    URL       string
}
