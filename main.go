package main

import (
	"flag"
	"fmt"
    "os"
	"goscript/config"
	"goscript/internal"
	"goscript/model"
	"goscript/redisModel"
	"time"
)

var (
    h bool
    command string
)

func init() {
    flag.BoolVar(&h, "h", false, "this script's usage")
	flag.StringVar(&command, "command", "", "set command: clean_battle,recover_battle")
    flag.Usage = Usage
}

func main() {
	flag.Parse()

    if h {
        flag.Usage()
    }

	configs := config.Config()
	db := configs.Database
	model.InitDB(db.Host, db.Port, db.Dbname, db.Username, db.Password)
	redisModel.InitRedis(configs.Redis.Server, configs.Redis.Password)

    switch command {
    case "clean_battle":
        cleanBattle()
    case "recover_battle":
        recoverBattle()
    default:
        flag.Usage()
    }
}

func Usage() {
    fmt.Fprintf(os.Stderr, `Minigame system script
Version: v1.0
Usage: goscript [-h] [-command=clean_battle]
Options:
`)
    flag.PrintDefaults()
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
