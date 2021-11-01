package testutils

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func InitTestDb() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Println("Error loading .env file")
	}
	db := TestDbConnection()
	defer db.Close()

	path := filepath.Join("../../schema.sql")
	c, ioErr := ioutil.ReadFile(path)
	if ioErr != nil {
		log.Fatal("Error loading schema.sql file")
	}
	sql := string(c)
	db.Exec(sql)

	path = filepath.Join("../../testdata.sql")
	c, ioErr = ioutil.ReadFile(path)
	if ioErr != nil {
		log.Fatal("Error loading schema.sql file")
	}
	sql = string(c)
	db.Exec(sql)
}

func TestDbConnection() *sql.DB {
	psql := os.Getenv("TEST_DB_STR")
	db, err := sql.Open("postgres", psql)

	if err != nil {
		panic(err)
	}
	return db
}

func ClearDb() {
	db := TestDbConnection()
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
		log.Println(err)
	}
}
