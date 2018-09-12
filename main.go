package main

import (
    "flag"
    "fmt"
    "github.com/garyburd/redigo/redis"
    "goscript/config"
    "goscript/model"
    "goscript/redisModel"
)

func main() {
    flag.Parse()

    configs := config.Config()
    db := configs.Database

    model.InitDB(db.Host, db.Port, db.Dbname, db.Username, db.Password)
    pool := redisModel.InitRedis(configs.Redis.Server, configs.Redis.Password)

    redisModel.CleanBattle()

    sql := "select * from user limit 3"
    record := model.Finds(sql)
    fmt.Println(record)

    c := pool.Get()
    defer c.Close()

    _, err := c.Do("SET", "username", "nick", "EX", "60")
    if err != nil {
        fmt.Println("redis set failed:", err)
    }

    username, err := redis.String(c.Do("GET", "username"))
    if err != nil {
        fmt.Println("redis get failed:", err)
    } else {
        fmt.Printf("Got username %v \n", username)
    }
}
