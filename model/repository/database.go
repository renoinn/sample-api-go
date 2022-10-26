package repository

import (
    "database/sql"
    "fmt"
)

var Db *sql.DB

func init() {
    var err error
    dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", "root", "root", "localhost:3306", "bookmarks_db")
    Db, err = sql.Open("mysql", dataSourceName)
    if err != nil {
        panic(err)
    }
}
