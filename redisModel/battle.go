package redisModel

import (
    "time"
    "errors"
    "goscript/internal"
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
