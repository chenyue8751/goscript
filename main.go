package main

import (
    "log"
    "flag"
    "fmt"
    "time"
    "github.com/garyburd/redigo/redis"
    "redis/model"
)

func main() {
    flag.Parse()

    model.initDB();

    rows, err := db.Query("select * from user limit 3")
    defer rows.Close()
    if err != nil {
        log.Println("exec error:",err)
    }
    columns, _ := rows.Columns()
    scanArgs := make([]interface{}, len(columns))
    values := make([]interface{}, len(columns))
    for j := range values {
        scanArgs[j] = &values[j]
    }
    record := make(map[string]string)
    for rows.Next() {
        err = rows.Scan(scanArgs...)
        for i, col := range values {
            if col != nil {
                record[columns[i]] = string(col.([]byte))
            }
        }
        fmt.Println(record)
    }

    c := pool.Get()
    defer c.Close()

    _, err = c.Do("SET", "username", "nick")
    if err != nil {
        fmt.Println("redis set failed:", err)
    }

    username, err := redis.String(c.Do("GET", "username"))
    if err != nil {
        fmt.Println("redis get failed:", err)
    } else {
        fmt.Printf("Got username %v \n", username)
    }

    _, err = c.Do("SET", "password", "123456", "EX", "2")
    if err != nil {
        fmt.Println("redis set failed:", err)
    }

    time.Sleep(3 * time.Second)
    password, err := redis.String(c.Do("GET", "password"))
    if err != nil {
        fmt.Println("redis get failed:", err)
    } else {
        fmt.Printf("Got password %v \n", password)
    }
}
