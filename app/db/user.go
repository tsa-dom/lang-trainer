package db

import (
	_ "github.com/lib/pq"
)

type AuthInfo struct {
	PasswordHash string
	Username     string
	Priviledges  string
}

func CreateUser(username, passwordHash, priviledges string) int {
	sql := `
		INSERT INTO Users (username, passwordHash, priviledges)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	db, _ := getDbConnection()
	id := -1
	db.QueryRow(sql, username, passwordHash, priviledges).Scan(&id)
	db.Close()

	return id
}

func UserAuthInfo(username string) (AuthInfo, error) {
	sql := `
		SELECT username, passwordHash, priviledges FROM Users WHERE username=$1
	`
	db, _ := getDbConnection()
	user := AuthInfo{}
	row := db.QueryRow(sql, username)
	err := row.Scan(&user.Username, &user.PasswordHash, &user.Priviledges)
	db.Close()

	if err != nil {
		nilInfo := AuthInfo{}
		return nilInfo, err
	}

	return user, nil
}
