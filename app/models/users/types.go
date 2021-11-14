package users

type User struct {
	Id           int    `json:"id"`
	PasswordHash string `json:"-"`
	Username     string `json:"username"`
	Privileges   string `json:"priviledges"`
}
