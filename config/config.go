package config

import (
    "github.com/BurntSushi/toml"
    "sync"
)

type tomlConfig struct {
    Environment string
    Database database
    Redis redisConfig
}

type database struct {
    Host string
    Port int
    Dbname string
    Username string
    Password string
}

type redisConfig struct {
    Server string
    Password string
}

var (
    cfg * tomlConfig
    once sync.Once
)

func Config() *tomlConfig {
    once.Do(func () {
        filePath := "./config.toml"
        if _, err := toml.DecodeFile(filePath, &cfg); err != nil {
            panic(err)
        }
    })
    return cfg
}
