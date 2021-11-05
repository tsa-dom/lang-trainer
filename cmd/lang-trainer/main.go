package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/tsa-dom/lang-trainer/app/models"
	"github.com/tsa-dom/lang-trainer/app/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	models.InitDB("schema.sql")

	router.Run()
}
