package users

import (
	_ "github.com/lib/pq"
	"github.com/tsa-dom/lang-trainer/app/models"
)

func CreateUser(user User) (User, error) {
	db := models.GetDbConnection()
	defer db.Close()

	err := db.QueryRow(addNewUser(), user.Username, user.PasswordHash, user.Priviledges).Scan(&user.Id)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func GetUserByUsername(username string) (User, error) {
	db := models.GetDbConnection()
	defer db.Close()

	user := User{}
	err := db.QueryRow(userByUsername(), username).Scan(&user.Id, &user.Username, &user.PasswordHash, &user.Priviledges)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func RemoveUser(userId int) error {
	db := models.GetDbConnection()
	defer db.Close()

	sql := `
		DELETE FROM Users WHERE id=$1
	`
	_, err := db.Exec(sql, userId)

	return err
}
