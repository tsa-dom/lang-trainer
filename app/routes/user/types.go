package router

type User struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Privileges string `json:"privileges"`
}
