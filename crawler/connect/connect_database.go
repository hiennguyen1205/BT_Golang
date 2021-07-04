package connect

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

const (
	username = "root"
	password = "root"
	hostname = "127.0.0.1:3306"
	dbname   = "crawl"
)

func dsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
}

func Connect() *sql.DB {
	db, err := sql.Open("mysql", dsn())
	if err != nil {
		panic(err)
	}
	fmt.Println("Connect successed")
	return db
}
