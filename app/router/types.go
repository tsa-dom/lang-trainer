package router

type User struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Privileges string `json:"priviledges"`
}

type AuthorizedUser struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	Privileges string `json:"priviledges"`
	Token      string `json:"token"`
}

type authKey struct{}
