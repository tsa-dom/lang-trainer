package types

type Group struct {
	Id          int    `json:"id"`
	OwnerId     int    `json:"ownerId"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (g *Group) setName(name string) {
	g.Name = name
}

func (g Group) getName() string {
	return g.Name
}
