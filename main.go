package main

import (
	"flag"
	"fmt"
    "time"
	"goscript/config"
	"goscript/model"
	"goscript/redisModel"
    "goscript/internal"
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

func recoverData() bool {
    date := internal.ThisMonday(time.Now()).Format("2006-01-02")
	record := model.BattlePlays(date)
	data := make(map[int]map[int]map[int]int)
	for _, item := range record {
		user, ok := data[item.UserId]
		if !ok {
			user = make(map[int]map[int]int)
			data[item.UserId] = user
		}
		game, ok := data[item.UserId][item.GameId]
		if !ok {
			game = make(map[int]int)
			data[item.UserId][item.GameId] = game
		}

		data[item.UserId][item.GameId][item.Score]++
	}
	
    redisModel.InitBattle(data, date)

    return true
}
