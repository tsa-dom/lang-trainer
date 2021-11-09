package groups

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

type WordKey struct {
	Id          int
	OwnerId     int
	Name        string
	Description string
}
