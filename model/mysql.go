package mysql

import (
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "log"
    "fmt"
    "github"
)

var db *sql.DB

func initDB(*config Database) {
    server := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset-utf8", config.Username, config.Password, config.Host, config.Port, config.Dbname)
    var err error
    db, err = sql.Open("mysql", server)
    if err != nil {
        log.Println("connect error:",err)
    }
}
