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
    sql := "SELECT user_id, score, game_id, battle.create_at FROM battle_player left join battle on battle.id = battle_player.battle_id;"
    record := model.Finds(sql)
    data := make(map[int]map[int]map[int]int)
    for _, item  := range record {
        data[item["user_id"]][item["game_id"]][item["score"]]++
    }
    return data
}
