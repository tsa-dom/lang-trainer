package main

import (
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/tsa-dom/language-trainer/app/db"
	"github.com/tsa-dom/language-trainer/app/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	path := filepath.Join("schema.sql")
	c, ioErr := ioutil.ReadFile(path)
	if ioErr != nil {
		log.Fatal("Error loading schema.sql file")
	}
	sql := string(c)
	db.InitDB(sql)
	if err != nil {
		log.Println(err)
	}

	router.Run()
}
