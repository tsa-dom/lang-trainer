package models

import (
	_ "github.com/lib/pq"
	conn "github.com/tsa-dom/lang-trainer/app/db"
	g "github.com/tsa-dom/lang-trainer/app/types"
)

// Deletes groups and their links to words

// Connects an existing word to a group
func AddWordToGroup(groupId, wordId int) error {
	db := conn.GetDbConnection()

	defer db.Close()

	_, err := db.Exec(linkWordToGroup(), groupId, wordId)
	if err != nil {
		return err
	}

	return nil
}

// Returns groups which belong to a specific user
/* func GetGroups(ownerId int) ([]g.Group, error) {

} */

// Returns words that belong to a specific group
func GetWordsInGroup(groupId int) ([]g.Word, error) {
	db := conn.GetDbConnection()
	defer db.Close()

	words := []g.Word{}

	rows, err := db.Query(wordsInGroup(), groupId)
	if err != nil {
		return nil, err
	}

	wordMap := make(map[g.WordKey][]g.WordItem)
	emptyBases := []g.WordKey{}
	for rows.Next() {
		base := g.WordKey{}
		item := g.WordItem{}
		err := rows.Scan(&base.Id, &base.OwnerId, &base.Name, &base.Description, &item.Id, &item.Name, &item.Description)

		if err == nil {
			wordMap[base] = append(wordMap[base], item)
		} else {
			emptyBases = append(emptyBases, base)
		}
	}

	for base, items := range wordMap {
		word := g.Word{Id: base.Id, Name: base.Name, Description: base.Description, OwnerId: base.OwnerId, Items: items, GroupId: groupId}
		words = append(words, word)
	}

	for _, base := range emptyBases {
		if base.Id != 0 {
			word := g.Word{Id: base.Id, Name: base.Name, Description: base.Description, OwnerId: base.OwnerId, Items: []g.WordItem{}, GroupId: groupId}
			words = append(words, word)
		}
	}

	return words, nil
}
