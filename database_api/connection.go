package database_api

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func OpenConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"postgres", 5432, "user", "mpassword", "db")

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	log.Println("server started and listening on http://127.0.0.1:8080")

	return db
}
