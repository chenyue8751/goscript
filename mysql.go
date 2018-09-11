package main

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
    "fmt"
)

var db *sql.DB

func init() {
    server := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset-utf8", Config().Database.Username, Config().Database.Password, Config().Database.Host, Config().Database.Port, Config().Database.Dbname)
    var err error
    db, err = sql.Open("mysql", server)
    if err != nil {
        log.Println("connect error:",err)
    }
}
