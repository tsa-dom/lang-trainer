package db

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

func getDbConnection() (*sql.DB, context.Context) {
	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		panic(err)
	}

	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, username, password, database)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	ctx := context.Background()
	return db, ctx
}

func InitDB(sql string) {
	db, _ := getDbConnection()
	db.Exec(sql)
	db.Close()
}
