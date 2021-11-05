package models

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	_ "github.com/lib/pq"
)

func GetDbConnection() *sql.DB {
	port := os.Getenv("DB_PORT")
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")
	sslmode := os.Getenv("DB_SSLMODE")

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=%s",
		host, port, username, password, database, sslmode)
	db, err := sql.Open("postgres", psqlInfo)

	if err != nil {
		panic(err)
	}
	return db
}

func InitDB(file string) {
	path := filepath.Join(file)

	c, ioErr := ioutil.ReadFile(path)
	if ioErr != nil {
		log.Fatal("Error loading schema.sql file")
	}

	sql := string(c)
	db := GetDbConnection()
	db.Exec(sql)
	db.Close()
}
