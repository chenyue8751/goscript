package config

import (
    "github.com/BurntSushi/toml"
    "sync"
    "path/filepath"
    "os"
    "os/exec"
    "strings"
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

func Config(configFilePath string) *tomlConfig {
    once.Do(func () {
        if _, err := toml.DecodeFile(configFilePath, &cfg); err != nil {
            panic(err)
        }
    })
    return cfg
}

func getAppPath() string {
    file, _ := exec.LookPath(os.Args[0])
    path, _ := filepath.Abs(file)
    index := strings.LastIndex(path, string(os.PathSeparator))

    return path[:index]
}
