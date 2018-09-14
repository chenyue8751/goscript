package main

import (
	"flag"
	"fmt"
	"goscript/config"
	"goscript/internal"
	"goscript/model"
	"goscript/redisModel"
	"time"
)

func main() {
	var command string
	flag.StringVar(&command, "command", "clean_battle", "choose on command.eg: help,clean_battle,recover_battle")
	flag.Parse()

	configs := config.Config()
	db := configs.Database
	model.InitDB(db.Host, db.Port, db.Dbname, db.Username, db.Password)
	redisModel.InitRedis(configs.Redis.Server, configs.Redis.Password)

}

func cleanBattle() {
	count, _ := redisModel.CleanBattle()
	fmt.Println("delete keys, nums: ", count)
}

func recoverBattle() bool {
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
