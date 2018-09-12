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

func Finds(sql string) []map[string]string {
    rows, err := db.Query(sql)
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
    result := make([]map[string]string, 0)
    record := make(map[string]string)
    for rows.Next() {
        err = rows.Scan(scanArgs...)
        for i, col := range values {
            if col != nil {
                record[columns[i]] = string(col.([]byte))
            }
        }
        result = append(result, record)
    }
    return result
}
