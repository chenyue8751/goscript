package redisModel

import (
	"errors"
	"fmt"
	"goscript/internal"
	"time"
)

func CleanBattle() (int, error) {
	lastMonday := internal.LastMonday(time.Now()).Format("2006-01-02")
	keyPattern := "minigame:battlePlayer:" + lastMonday + ":*"
	keys := getKeys(keyPattern)
	count := deleteMulti(keys)
	if count == len(keys) {
		return count, nil
	} else {
		return count, errors.New("deleted count not right")
	}
}

func InitBattle(data map[int]map[int]map[int]int, date string) {
	conn := pool.Get()
	defer conn.Close()
	for userId, game := range data {
		for gameId, score := range game {
			key := fmt.Sprintf("minigame:battlePlayer:%s:%d:%d", date, userId, gameId)
			count := map[string]int{"win": 0, "tie": 0, "lose": 0}

			winCount, ok := score[1]
			if ok {
				count["win"] = winCount
			}
			tieCount, ok := score[0]
			if ok {
				count["tie"] = tieCount
			}
			loseCount, ok := score[-1]
			if ok {
				count["lose"] = loseCount
			}

			conn.Do("HMSET", key, "win", count["win"], "tie", count["tie"], "lose", count["lose"])
		}
	}
}
