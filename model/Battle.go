package model

import (
    "log"
)

type BattlePlay struct {
    UserId int
    GameId int
    Score int
}

func BattlePlays(date string) []*BattlePlay {
    sql := "SELECT user_id, score, game_id FROM battle_player LEFT JOIN battle ON battle.id = battle_player.battle_id WHERE battle.create_at >= '" + date + "'"
    rows, err := db.Query(sql)
    defer rows.Close()
    if err != nil {
        log.Println("exec error:", err)
    }

    result := make([]*BattlePlay, 0)
    for rows.Next() {
        record := new(BattlePlay)
        err = rows.Scan(&record.UserId, &record.Score, &record.GameId)
        if err == nil {
            result = append(result, record)
        } else {
            log.Println(err)
        }
    }
    return result
}
