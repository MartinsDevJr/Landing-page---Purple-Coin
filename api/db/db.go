package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error

	DB, err = sql.Open("mysql",  "evelyn:1612@tcp(localhost:3306)/users")
	if err != nil {
		log.Fatal("Error connecting to database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Error checking database connection:", err)
	}

	log.Println("Successful connection to the database.")
}
