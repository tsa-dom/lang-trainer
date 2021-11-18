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

	/* hash, _ := utils.HashPassword("test")
	user := types.User{Username: "Student", PasswordHash: hash, Privileges: "student"}
	models.CreateUser(user) */

	router.Run()
}
