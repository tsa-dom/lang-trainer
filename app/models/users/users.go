package models

import (
	_ "github.com/lib/pq"
	conn "github.com/tsa-dom/lang-trainer/app/db"
	g "github.com/tsa-dom/lang-trainer/app/types"
)

func CreateUser(user g.User) (g.User, error) {
	db := conn.GetDbConnection()
	defer db.Close()

	err := db.QueryRow(addNewUser(), user.Username, user.PasswordHash, user.Privileges).Scan(&user.Id)
	if err != nil {
		return g.User{}, err
	}

	return user, nil
}

func GetUserByUsername(username string) (g.User, error) {
	db := conn.GetDbConnection()
	defer db.Close()

	user := g.User{}
	err := db.QueryRow(userByUsername(), username).Scan(&user.Id, &user.Username, &user.PasswordHash, &user.Privileges)
	if err != nil {
		return g.User{}, err
	}

	return user, nil
}

func RemoveUser(userId int) error {
	db := conn.GetDbConnection()
	defer db.Close()

	sql := `
		DELETE FROM Users WHERE id=$1
	`
	_, err := db.Exec(sql, userId)

	return err
}
