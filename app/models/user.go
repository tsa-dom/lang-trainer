package models

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

type User struct {
	Id           int
	PasswordHash string
	Username     string
	Priviledges  string
}

func CreateUser(user User) error {
	db := getDbConnection()
	defer db.Close()
	defer log.Println("Connection closed")
	err := createUser(db, user)
	if err != nil {
		return err
	}
	return nil
}

func createUser(db *sql.DB, user User) error {
	sql := `
		INSERT INTO Users (username, password_hash, priviledges)
		VALUES ($1, $2, $3)
	`
	_, err := db.Exec(sql, user.Username, user.PasswordHash, user.Priviledges)
	if err != nil {
		return err
	}

	return nil
}

func GetUserByUsername(username string) (User, error) {
	db := getDbConnection()
	defer db.Close()
	defer log.Println("Connection closed")
	user, err := getUserByUsername(db, username)
	if err != nil {
		return User{}, err
	}
	return user, nil
}

func getUserByUsername(db *sql.DB, username string) (User, error) {
	sql := `
		SELECT id, username, password_hash, priviledges FROM Users WHERE username=$1
	`
	user := User{}
	row := db.QueryRow(sql, username)
	err := row.Scan(&user.Id, &user.Username, &user.PasswordHash, &user.Priviledges)

	if err != nil {
		return User{}, err
	}
	return user, nil
}

func RemoveUser(userId int) error {
	db := getDbConnection()
	defer db.Close()
	defer log.Panicln("Connection closed")
	err := removeUser(db, userId)
	return err
}

func removeUser(db *sql.DB, userId int) error {
	sql := `
		DELETE FROM Users WHERE id=$1
	`
	_, err := db.Exec(sql, userId)

	return err
}
