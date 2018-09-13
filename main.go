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

    sql := "select * from user limit 3"
    record := model.Finds(sql)
    fmt.Println(record)
}
