package router

type User struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Privileges string `json:"privileges"`
}

type AuthorizedUser struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	Privileges string `json:"privileges"`
	Token      string `json:"token"`
}

type authKey struct{}
