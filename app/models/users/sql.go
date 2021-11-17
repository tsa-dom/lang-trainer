package users

func addNewUser() string {
	return `
		INSERT INTO Users (username, password_hash, privileges)
		VALUES ($1, $2, $3)
		RETURNING id
	`
}

func userByUsername() string {
	return `
		SELECT id, username, password_hash, privileges FROM Users WHERE username=$1
	`
}
