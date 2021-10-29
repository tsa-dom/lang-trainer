package main

import (
	//"log"

	//"github.com/joho/godotenv"
	"github.com/tsa-dom/language-trainer/app/db"
	"github.com/tsa-dom/language-trainer/app/http"
)

func main() {
	//err := godotenv.Load()
	//if err != nil {
	//	log.Fatal("Error loading .env file")
	//}
	db.Connect()
	http.Router()
}
