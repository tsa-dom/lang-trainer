package models

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func getDbConnection() *sql.DB {
	port := os.Getenv("DB_PORT")
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, username, password, database)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	return db
}

func InitDB(sql string) {
	db := getDbConnection()
	db.Exec(sql)
	db.Close()
}
