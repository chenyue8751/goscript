package main

import (
    "log"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "flag"
    "fmt"
    "time"
    "github.com/garyburd/redigo/redis"
)

func main() {
    flag.Parse()

    db,err := sql.Open("mysql", "Config().Username:Config().Password@tcp(Config().Host:Config().Port)/Config().Dbname?charset-utf8")
    if err != nil {
        log.Println("connect error:",err)
    }
    var user string
    db.QueryRow("select username from user limit 1").Scan(&user)
    if err != nil {
        log.Println("exec error:",err)
    }
    log.Println("result:",user)

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
