package types

import "github.com/golang-jwt/jwt"

type User struct {
	Id           int    `json:"id"`
	PasswordHash string `json:"-"`
	Username     string `json:"username"`
	Privileges   string `json:"privileges"`
}

type Claims struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	Privileges string `json:"privileges"`
	jwt.StandardClaims
}

type AuthorizedUser struct {
	Id         int    `json:"id"`
	Username   string `json:"username"`
	Privileges string `json:"privileges"`
	Token      string `json:"token"`
}

type WordItem struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Word struct {
	Id          int        `json:"id"`
	OwnerId     int        `json:"ownerId"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	GroupId     int        `json:"groupId"`
	Items       []WordItem `json:"items"`
}

type Group struct {
	Id          int    `json:"id"`
	OwnerId     int    `json:"ownerId"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Template struct {
	Id           int      `json:"id"`
	OwnerId      int      `json:"ownerId"`
	Name         string   `json:"name"`
	Descriptions []string `json:"descriptions"`
}

type WordKey struct {
	Id          int
	OwnerId     int
	Name        string
	Description string
}

type GroupIds struct {
	Ids []int `json:"groupIds"`
}

type WordIds struct {
	Ids []int `json:"wordIds"`
}

type TemplateIds struct {
	Ids []int `json:"templateIds"`
}
