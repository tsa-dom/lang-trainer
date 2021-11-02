package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/tsa-dom/lang-trainer/app/models"
	"github.com/tsa-dom/lang-trainer/app/models/users"
	"github.com/tsa-dom/lang-trainer/app/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	models.InitDB("schema.sql")

	user := users.User{Username: "Admin2", PasswordHash: "salainen", Priviledges: "admin"}
	users.CreateUser(user)
	router.Run()
}

/* func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	db.InitPsql()
	err = db.Ping()
	if err != nil {
		log.Fatal("Failed to ping database")
	}

	path := filepath.Join("schema.sql")

	c, ioErr := ioutil.ReadFile(path)
	if ioErr != nil {
		log.Fatal("Error loading schema.sql file")
	}
	sql := string(c)
	err = db.InitDB(sql)
	if err != nil {
		log.Fatal(err)
	}

	db.CreateUser("Admin", "salainen", "admin")

  group := db.Group{}
	group.Name = "awesome"
	group.Description = "somenthing different"
	group.OwnerId = 1
	err = db.CreateGroup(group)
	log.Println(err) */

/* word := db.Word{}
word.Description = "worddesc"
word.Name = "wordname"
word.OwnerId = 1

item := db.WordItem{}
item.Name = "slikkesmile"
item.Description = "do it"
item2 := db.WordItem{}
item2.Name = "slikkesmile second"
item2.Description = "do it now"

word.Items = []db.WordItem{item, item2}
err = db.CreateWord(word)
log.Println(err)
err = db.AddItemsToWord(4, word.Items)
log.Println(err) */
