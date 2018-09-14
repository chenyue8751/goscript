package model

import (
    "log"
)

type BattlePlay struct {
    UserId int
    GameId int
    Score int
}

func BattlePlays() []*BattlePlay {
    sql := "SELECT user_id, score, game_id FROM battle_player left join battle on battle.id = battle_player.battle_id;"
    rows, err := db.Query(sql)
    defer rows.Close()
    if err != nil {
        log.Println("exec error:", err)
    }

    result := make([]*BattlePlay, 0)
    record := new(BattlePlay)
    for rows.Next() {
        err = rows.Scan(&record.UserId, &record.Score, &record.GameId)
        if err == nil {
            result = append(result, record)
        } else {
            log.Println(err)
        }
    }
    return result
}
