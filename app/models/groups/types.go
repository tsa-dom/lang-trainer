package groups

type WordItem struct {
	Id          int
	Name        string
	Description string
}

type Word struct {
	Id          int
	OwnerId     int
	Name        string
	Description string
	Items       []WordItem
}

type Group struct {
	Id          int    `json:"id"`
	OwnerId     int    `json:"ownerId"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
