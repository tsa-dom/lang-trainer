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
		INSERT INTO users (username, passwordHash, priviledges)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	db := getDbConnection()
	id := -1
	db.QueryRow(sql, username, passwordHash, priviledges).Scan(&id)
	db.Close()

	return id
}

func UserAuthInfo(username string) (AuthInfo, error) {
	sql := `
		SELECT username, passwordHash, priviledges FROM users WHERE username=$1
	`
	db := getDbConnection()
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
