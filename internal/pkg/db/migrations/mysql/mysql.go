package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func InitDb() {
	conn, err := sql.Open("mysql", "root:root@tcp(localhost)/Auth")

	if err != nil {
		log.Fatal(err)
	}

	err = conn.Ping()

	if err != nil {
		log.Fatal(err)
	}

	Db = conn

	log.Println("Connected to DB")

}

func CloseDB() error {
	return Db.Close()
}
