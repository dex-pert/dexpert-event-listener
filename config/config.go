package config

type Config struct {
    ChainConfig []ChainConfig        `yaml:"chain_config"`
    MySQL       *MySQL               `yaml:"mysql"`
    Chains      map[int]*ChainConfig `yaml:"chains"`
}

type MySQL struct {
    User string `yaml:"user"`
    Pass string `yaml:"pass"`
    Host string `yaml:"host"`
    Port string `yaml:"port"`
    DB   string `yaml:"db"`
}

type ChainConfig struct {
    ChainId   int    `yaml:"chain_id"`
    ChainName string `yaml:"chain_name"`
    URL       string `yaml:"url"`
}
