package models

import (
	"context"

	_ "github.com/lib/pq"
	conn "github.com/tsa-dom/lang-trainer/app/db"
	g "github.com/tsa-dom/lang-trainer/app/types"
)

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

func CreateWord(word g.Word) (g.Word, error) {
	db := conn.GetDbConnection()
	defer db.Close()

	err := db.QueryRow(addWord(), word.OwnerId, word.Name, word.Description).Scan(&word.Id)

	if err != nil {
		return g.Word{}, err
	}

	return word, nil
}

func GetWordById(wordId int) (g.Word, error) {
	db := conn.GetDbConnection()
	defer db.Close()

	word := g.Word{}

	err := db.QueryRow(wordById(), wordId).Scan(&word.Id, &word.OwnerId, &word.Name, &word.Description)
	if err != nil {
		return g.Word{}, err
	}

	rows, err := db.Query(wordItemsByWordId(), wordId)
	if err != nil {
		return g.Word{}, err
	}

	items := []g.WordItem{}
	for rows.Next() {
		item := g.WordItem{}
		err := rows.Scan(&item.Id, &item.Name, &item.Description)
		if err != nil {
			return g.Word{}, nil
		}
		items = append(items, item)
	}
	word.Items = items

	return word, nil
}

func AddItemsToWord(wordId int, items []g.WordItem) ([]g.WordItem, error) {
	db := conn.GetDbConnection()
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return []g.WordItem{}, nil
	}
	defer tx.Rollback()
	defer db.Close()
	wordItems := []g.WordItem{}

	for _, item := range items {
		err := tx.QueryRow(addWordItem(), wordId, item.Name, item.Description).Scan(&item.Id)
		if err != nil {
			return []g.WordItem{}, err
		}
		wordItems = append(wordItems, item)
	}

	if err = tx.Commit(); err != nil {
		return []g.WordItem{}, err
	}

	return wordItems, nil
}

func RemoveWord(wordId int) error {
	db := conn.GetDbConnection()
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	defer db.Close()

	_, err = db.Exec(deleteItemsByWordId(), wordId)
	if err != nil {
		return err
	}

	_, err = db.Exec(deleteWordById(), wordId)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func AddWordToGroup(groupId, wordId int) error {
	db := conn.GetDbConnection()

	defer db.Close()

	_, err := db.Exec(linkWordToGroup(), groupId, wordId)
	if err != nil {
		return err
	}

	return nil
}

func RemoveWordFromGroup(groupId, wordId int) error {
	db := conn.GetDbConnection()

	defer db.Close()

	_, err := db.Exec(deleteWordLink(), groupId, wordId)
	if err != nil {
		return err
	}

	return nil
}

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

func GetWordsInGroup(groupId int) ([]g.Word, error) {
	db := conn.GetDbConnection()
	defer db.Close()

	words := []g.Word{}

	rows, err := db.Query(wordsInGroup(), groupId)
	if err != nil {
		return nil, err
	}

	wordMap := make(map[g.WordKey][]g.WordItem)
	for rows.Next() {
		base := g.WordKey{}
		item := g.WordItem{}
		err := rows.Scan(&base.Id, &base.OwnerId, &base.Name, &base.Description, &item.Id, &item.Name, &item.Description)
		if err == nil {
			wordMap[base] = append(wordMap[base], item)
		}
	}

	for base, items := range wordMap {
		word := g.Word{Id: base.Id, Name: base.Name, Description: base.Description, OwnerId: base.OwnerId, Items: items, GroupId: groupId}
		words = append(words, word)
	}

	return words, nil
}
