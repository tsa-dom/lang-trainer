package users

func addNewUser() string {
	return `
		INSERT INTO Users (username, password_hash, priviledges)
		VALUES ($1, $2, $3)
		RETURNING id
	`
}

func userByUsername() string {
	return `
		SELECT id, username, password_hash, priviledges FROM Users WHERE username=$1
	`
}
