package main

import (
	"fmt"
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

	item := db.WordItem{}
	item.Word = "slikkesmile"
	item.Description = "do it"
	item2 := db.WordItem{}
	//item2.Word = "slikkesmile second"
	item2.Description = "do it now"

	word := db.Word{}
	word.OwnerId = 2
	word.BaseWordItem.Word = "test"
	word.BaseWordItem.Description = "testdesc"
	word.LinkedItems = []db.WordItem{item, item2}

	err = db.CreateWord(word)
	fmt.Println(err)

	router.Run()
}
