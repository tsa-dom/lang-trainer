package main

import (
	"log"

	"github.com/joho/godotenv"
	router "github.com/tsa-dom/lang-trainer/app/controller"
	"github.com/tsa-dom/lang-trainer/app/db"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	db.InitDB("schema.sql")

	router.Run()
}
