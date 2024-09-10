package config

type Config struct {
    Chains []Chain
    MySQL  *MySQL `yaml:"mysql"`
}

type MySQL struct {
    User string `yaml:"user"`
    Pass string `yaml:"pass"`
    Host string `yaml:"host"`
    Port string `yaml:"port"`
    DB   string `yaml:"db"`
}

type Chain struct {
    ChainId   int
    ChainName string
    URL       string
}
