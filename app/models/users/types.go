package users

type User struct {
	Id           int    `json:"id"`
	PasswordHash string `json:"-"`
	Username     string `json:"username"`
	Priviledges  string `json:"priviledges"`
}
