package users

import (
	"log"

	_ "github.com/lib/pq"
	"github.com/tsa-dom/lang-trainer/app/models"
)

type User struct {
	Id           int
	PasswordHash string
	Username     string
	Priviledges  string
}

func CreateUser(user User) error {
	db := models.GetDbConnection()
	defer db.Close()
	defer log.Println("Connection closed")

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
	db := models.GetDbConnection()
	defer db.Close()
	defer log.Println("Connection closed")

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
	db := models.GetDbConnection()
	defer db.Close()
	defer log.Panicln("Connection closed")

	sql := `
		DELETE FROM Users WHERE id=$1
	`
	_, err := db.Exec(sql, userId)

	return err
}
