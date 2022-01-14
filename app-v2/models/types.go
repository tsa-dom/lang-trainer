package models

type Groups struct {
	Values []Group `json:"groups"`
}

type Group struct {
	Id          int    `json:"id"`
	OwnerId     int    `json:"ownerId"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type Words struct {
	Values []Word `json:"words"`
}

type Word struct {
	Id          int        `json:"id"`
	OwnerId     int        `json:"ownerId"`
	Name        string     `json:"name"`
	Description string     `json:"description"`
	GroupId     int        `json:"groupId"`
	Items       []WordItem `json:"items"`
}

type WordItem struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type GroupIds struct {
	Ids []int `json:"groupIds"`
}
