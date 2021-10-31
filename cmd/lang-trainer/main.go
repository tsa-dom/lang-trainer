package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/tsa-dom/lang-trainer/app/db"
	"github.com/tsa-dom/lang-trainer/app/router"
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

	group := db.WordGroup{}
	group.OwnerId = 2
	group.Name = "testter"
	group.Description = "awesome desc"
	//db.CreateWordGroup(group)

	db.AddWordToGroup(1, 17)
	db.AddWordToGroup(1, 18)
	//db.RemoveWordFromGroup(1, 17)

	groups, _ := db.GetGroups(2)
	group = groups[0]
	fmt.Println(group.Id)
	fmt.Println(group.Name)
	fmt.Println(group.OwnerId)
	fmt.Println(group.Description)

	words, _ := db.GetWordsInGroup(1)
	fmt.Println(words[0].BaseWordItem)
	fmt.Println(words[0].LinkedItems)
	fmt.Println(words[1].BaseWordItem)
	fmt.Println(words[1].LinkedItems)
	fmt.Println(words[2].BaseWordItem)
	fmt.Println(words[2].LinkedItems)

	router.Run()
}
