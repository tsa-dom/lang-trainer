package models

import (
	"context"

	_ "github.com/lib/pq"
)

type WordItem struct {
	Id          int
	Name        string
	Description string
}

type Word struct {
	OwnerId     int
	Name        string
	Description string
	Items       []WordItem
}

type Group struct {
	Id          int
	OwnerId     int
	Name        string
	Description string
}

func CreateGroup(group Group) error {
	sql := `
		INSERT INTO Groups (owner_id, name, description)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	db := GetDbConnection()
	defer db.Close()

	_, err := db.Exec(sql, group.OwnerId, group.Name, group.Description)
	if err != nil {
		return err
	}

	return nil
}

func CreateWord(word Word) error {
	db := GetDbConnection()
	defer db.Close()

	sql := `
		INSERT INTO Words (owner_id, word, description)
		VALUES ($1, $2, $3)
	`

	_, err := db.Exec(sql, word.OwnerId, word.Name, word.Description)
	if err != nil {
		return err
	}

	return nil
}

func AddItemsToWord(wordId int, items []WordItem) error {
	db := GetDbConnection()
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return nil
	}
	defer tx.Rollback()
	defer db.Close()

	sql := `
		INSERT INTO WordItems (word_id, word, description)
		VALUES ($1, $2, $3)
	`

	for _, item := range items {
		_, err := tx.Exec(sql, wordId, item.Name, item.Description)
		if err != nil {
			return err
		}
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func RemoveWord(wordId int) error {
	db := GetDbConnection()
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	defer db.Close()

	sqlItems := `
		DELETE FROM WordItems WHERE word_id=$1
	`
	sql := `
		DELETE FROM Words WHERE id=$1;
	`

	_, err = db.Exec(sqlItems, wordId)
	if err != nil {
		return err
	}

	_, err = db.Exec(sql, wordId)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func AddWordToGroup(groupId, wordId int) error {
	db := GetDbConnection()

	defer db.Close()

	sql := `
		INSERT INTO GroupLinks (group_id, word_id)
		VALUES ($1, $2)
		RETURNING id
	`

	_, err := db.Exec(sql, groupId, wordId)
	if err != nil {
		return err
	}

	return nil
}

func RemoveWordFromGroup(groupId, wordId int) error {
	db := GetDbConnection()

	defer db.Close()

	sql := `
		DELETE FROM GroupLinks WHERE group_id=$1 AND word_id=$2
	`

	_, err := db.Exec(sql, groupId, wordId)
	if err != nil {
		return err
	}

	return nil
}

func GetGroups(ownerId int) ([]Group, error) {
	db := GetDbConnection()

	defer db.Close()

	sql := `
		SELECT id, owner_id, name, description 
		FROM WordGroups WHERE owner_id=$1
	`

	rows, err := db.Query(sql, ownerId)
	if err != nil {
		return nil, err
	}

	groups := []Group{}
	for rows.Next() {
		group := Group{}
		err := rows.Scan(&group.Id, &group.OwnerId, &group.Name, &group.Description)
		if err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}

	return groups, nil
}

/* func GetWordsInGroup(groupId int) ([]Word, error) {
	db, _ := getDbConnection()

	defer db.Close()

	sql := `
		SELECT A.id, A.word, A.description, B.id, B.word, B.description
		FROM (
			SELECT DISTINCT W.wordItemId, W.targetItemId
			FROM WordGroupLinks L, words W
			WHERE L.wordGroupId=$1
		) W
		LEFT JOIN WordItems A
		ON W.wordItemId=A.id
		LEFT JOIN WordItems B
		ON W.targetItemId=B.id
		ORDER BY A.id, B.id;
	`

	words := []Word{}

	rows, err := db.Query(sql, groupId)
	if err != nil {
		return nil, err
	}

	wordMap := make(map[WordItem][]WordItem)
	for rows.Next() {
		base := WordItem{}
		item := WordItem{}
		rows.Scan(&base.Id, &base.Word, &base.Description, &item.Id, &item.Word, &item.Description)
		if err != nil {
			return nil, err
		}
		wordMap[base] = append(wordMap[base], item)
	}

	for base, items := range wordMap {
		word := Word{}
		word.BaseWordItem = base
		word.LinkedItems = items
		words = append(words, word)
	}

	return words, nil
} */
