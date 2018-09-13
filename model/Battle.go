package model

type battlePlay struct {
    user_id int
    game_id int
    score int
}

func BattlePlays() []battlePlay {
    sql := "SELECT user_id, score, game_id, battle.create_at FROM battle_player left join battle on battle.id = battle_player.battle_id;"
    rows, err := db.Query(sql)
    defer rows.Close()
    if err != nil {
        log.Println("exec error:", err)
    }

    result := make([]battlePlay, 0)
    record := new(battlePlay)
    for rows.Next() {
        err = rows.Scan(&record.user_id, &record.score, &record.game_id)
        result = append(result, record)
    }
    return result
}
