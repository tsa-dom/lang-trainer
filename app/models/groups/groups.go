package models

import (
	"errors"

	_ "github.com/lib/pq"
	conn "github.com/tsa-dom/lang-trainer/app/db"
	g "github.com/tsa-dom/lang-trainer/app/types"
	"github.com/tsa-dom/lang-trainer/app/utils"
)

// Creates a new group
func CreateGroup(group g.Group) (g.Group, error) {
	db := conn.GetDbConnection()
	defer db.Close()

	row := db.QueryRow(addGroup(), group.OwnerId, group.Name, group.Description)
	err := row.Scan(&group.Id)
	if err != nil {
		return g.Group{}, err
	}

	return group, nil
}

func ModifyGroup(ownerId int, group g.Group) error {
	db := conn.GetDbConnection()
	defer db.Close()

	id := 0
	db.QueryRow(modifyGroup(), group.Id, ownerId, group.Name, group.Description).Scan(&id)

	if id == 0 {
		return errors.New("group modifications failed, are groupId and ownerId valid ?")
	}

	return nil
}

// Deletes groups and their links to words
func RemoveGroups(ownerId int, groupIds g.GroupIds) error {
	db := conn.GetDbConnection()
	defer db.Close()

	tx, err := db.Begin()
	defer tx.Rollback()
	if err != nil {
		return err
	}

	_, err = tx.Exec(deleteGroupLinks(groupIds))
	if err != nil {
		return err
	}

	rows, err := tx.Query(deleteGroups(groupIds), ownerId)
	if err != nil {
		return err
	}

	removed := []int{}
	for rows.Next() {
		id := 0
		err := rows.Scan(&id)
		if err == nil && id != 0 {
			removed = append(removed, id)
		}
	}

	if !utils.IntArrayEquality(removed, groupIds.Ids) {
		return errors.New("id's not match")
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

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
func GetGroups(ownerId int) ([]g.Group, error) {
	db := conn.GetDbConnection()

	defer db.Close()

	rows, err := db.Query(getGroups(), ownerId)
	if err != nil {
		return nil, err
	}

	groups := []g.Group{}
	for rows.Next() {
		group := g.Group{}
		err := rows.Scan(&group.Id, &group.OwnerId, &group.Name, &group.Description)
		if err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}

	return groups, nil
}

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
