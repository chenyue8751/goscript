package redisModel

import (
    "fmt"
    "time"
)

func CleanBattle() {
    lastMonday := lastMonday(time.Now()).Format("2006-01-02")
    keyPattern := "minigame:battlePlayer:" + lastMonday + ":*"
    keys := getKeys(keyPattern)
    fmt.Println(keys)
}

func lastMonday(now time.Time) time.Time {
    weekday := int(now.Weekday())
    if weekday == 0 {
        weekday = 7
    }
    year, month, day := now.Date()
    today := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
    lastMonday := today.AddDate(0, 0, -(6 + weekday))
    return lastMonday
}
