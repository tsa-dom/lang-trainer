package main

import (
	"log"

	"github.com/joho/godotenv"
	router "github.com/tsa-dom/lang-trainer/app/controller"
	"github.com/tsa-dom/lang-trainer/app/db"
	models "github.com/tsa-dom/lang-trainer/app/models/users"
	g "github.com/tsa-dom/lang-trainer/app/types"
	"github.com/tsa-dom/lang-trainer/app/utils"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	db.InitDB("schema.sql")

	// This shoulb be removed
	hash, _ := utils.HashPassword("salainen")
	user := g.User{Username: "Admin", PasswordHash: hash, Privileges: "admin"}
	models.CreateUser(user)

	router.Run()
}
