package db

import (
	"database/sql"
	"io/ioutil"
	"log"
	"path/filepath"

	_ "github.com/lib/pq"
	"github.com/tsa-dom/lang-trainer/app/utils"
)

func GetDbConnection() *sql.DB {

	db, err := sql.Open("postgres", utils.DBConnectionString())

	if err != nil {
		panic(err)
	}
	return db
}

func InitDB(file string) {
	path := filepath.Join(file)

	c, ioErr := ioutil.ReadFile(path)
	if ioErr != nil {
		log.Fatal("Error loading schema.sql file")
	}

	sql := string(c)

	db := GetDbConnection()
	defer db.Close()

	_, err := db.Exec(sql)
	if err != nil {
		log.Panic(err)
	}
}

func InitTestDb() {
	db := GetDbConnection()
	defer db.Close()

	_, err := db.Exec(userTable())
	if err != nil {
		log.Panic(err)
	}

	_, err = db.Exec(groupsTable())
	if err != nil {
		log.Panic(err)
	}

	_, err = db.Exec(wordsTable())
	if err != nil {
		log.Panic(err)
	}

	_, err = db.Exec(wordItemsTable())
	if err != nil {
		log.Panic(err)
	}

	_, err = db.Exec(groupLinksTable())
	if err != nil {
		log.Panic(err)
	}
}
