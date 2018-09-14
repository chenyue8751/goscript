package internal

import (
    "time"
)

func ThisMonday(now time.Time) time.Time {
    weekday := weekday(now)
    today := todayDate(now)
    thisMonday := today.AddDate(0, 0, 1 - weekday)
    return thisMonday
}

func LastMonday(now time.Time) time.Time {
    weekday := weekday(now)
    today := todayDate(now)
    lastMonday := today.AddDate(0, 0, -(6 + weekday))
    return lastMonday
}

func weekday(now time.Time) int {
    weekday := int(now.Weekday())
    if weekday == 0 {
        weekday = 7
    }
    return weekday
}

func todayDate(now time.Time) time.Time {
    year, month, day := now.Date()
    today := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
    return today
}
