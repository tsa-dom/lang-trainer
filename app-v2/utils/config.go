package utils

import (
	"fmt"
	"os"
)

func DBConnectionString() string {
	port := os.Getenv("DB_PORT")
	host := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")
	sslmode := os.Getenv("DB_SSLMODE")

	connString := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=%s",
		host, port, username, password, database, sslmode)

	return connString
}
