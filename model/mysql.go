package model

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
    "fmt"
)

var db *sql.DB

func InitDB(host string, port int, dbname, username, password string) bool {
    server := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset-utf8", username, password, host, port, dbname)
    var err error
    db, err = sql.Open("mysql", server)
    if err != nil {
        log.Println("connect error:",err)
        return false
    }
    return true
}
