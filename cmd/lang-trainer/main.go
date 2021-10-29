package main

import (
	"github.com/tsa-dom/language-trainer/app/server"
	"github.com/tsa-dom/language-trainer/app/db"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
  if err != nil {
    log.Fatal("Error loading .env file")
  }
	db.Connect()
	server.Router()
}