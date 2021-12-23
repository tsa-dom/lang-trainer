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
	InitDB("../../schema.sql")

	db := GetDbConnection()
	path := filepath.Join("../../testdata.sql")
	c, ioErr := ioutil.ReadFile(path)
	if ioErr != nil {
		log.Fatal("Error loading testdata.sql file")
	}
	sql := string(c)
	db.Exec(sql)
}

func ClearTestDb() {
	db := GetDbConnection()
	defer db.Close()
	clear := `
		DROP TABLE Users CASCADE;
		DROP TABLE Words CASCADE;
		DROP TABLE WordItems CASCADE;
		DROP TABLE Groups CASCADE;
		DROP TABLE GroupLinks CASCADE;
	`
	_, err := db.Exec(clear)
	if err != nil {
		log.Panic(err)
	}
}
