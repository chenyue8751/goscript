package main

import (
    "flag"
    "fmt"
    "goscript/config"
    "goscript/model"
    "goscript/redisModel"
)

func main() {
    flag.Parse()

    configs := config.Config()
    db := configs.Database
    model.InitDB(db.Host, db.Port, db.Dbname, db.Username, db.Password)
    redisModel.InitRedis(configs.Redis.Server, configs.Redis.Password)

    count, _ := redisModel.CleanBattle()
    fmt.Println("delete keys, nums: ", count)

    data := recoverData()
    fmt.Println(data)
}

func recoverData() map[int]map[int]map[int]int {
    record := model.BattlePlays()
    data := make(map[int]map[int]map[int]int)
    for _, item  := range record {
        fmt.Printf("%#v", item.UserId)
        data[item.UserId][item.GameId][item.Score]++
    }
    return data
}
