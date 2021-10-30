package db

import (
	_ "github.com/lib/pq"
)

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
