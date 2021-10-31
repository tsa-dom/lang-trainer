package db

import (
	"errors"
	"log"

	_ "github.com/lib/pq"
)

type WordItem struct {
	Id          int
	Word        string
	Description string
}

type Word struct {
	OwnerId      int
	BaseWordItem WordItem
	LinkedItems  []WordItem
}

type WordGroup struct {
	Id          int
	OwnerId     int
	Name        string
	Description string
}

func CreateWordGroup(group WordGroup) int {
	sql := `
		INSERT INTO WordGroups (ownerId, name, description)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	db, _ := getDbConnection()
	defer db.Close()

	id := -1
	db.QueryRow(sql, group.OwnerId, group.Name, group.Description).Scan(&id)

	return id
}

func CreateWord(word Word) error {
	db, ctx := getDbConnection()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()
	defer db.Close()

	sqlWordItem := `
		INSERT INTO WordItems (word, description)
		VALUES ($1, $2)
		RETURNING id
	`
	sqlWord := `
		INSERT INTO Words (ownerId, wordItemId, targetItemId)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	id := -1
	base := word.BaseWordItem
	tx.QueryRow(sqlWordItem, base.Word, base.Description).Scan(&id)
	if id == -1 {
		return errors.New("failed to create word item")
	}

	for _, item := range word.LinkedItems {
		itemId := -1
		tx.QueryRow(sqlWordItem, item.Word, item.Description).Scan(&itemId)
		if itemId == -1 {
			return errors.New("failed to create word item")
		}
		_, err = tx.ExecContext(ctx, sqlWord, word.OwnerId, id, itemId)
		if err != nil {
			return err
		}
	}

	if err = tx.Commit(); err != nil {
		return errors.New("failed to commit queries")
	}

	return nil
}

func RemoveWord(wordId int) error {
	db, ctx := getDbConnection()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()
	defer db.Close()

	sqlWordItems := `
		SELECT DISTINCT UNNEST(ARRAY[wordItemId, targetItemId]) FROM Words WHERE wordItemId=$1;
	`
	sqlDeleteItems := `
		DELETE FROM WordItems WHERE id=$1;
	`
	sqlDeleteWord := `
		DELETE FROM Words WHERE wordItemId=$1;
	`

	rows, err := db.Query(sqlWordItems, wordId)
	if err != nil {
		return err
	}
	defer rows.Close()

	var id int
	for rows.Next() {
		err := rows.Scan(&id)
		if err != nil {
			return err
		}

		_, err = db.ExecContext(ctx, sqlDeleteItems, id)
		if err != nil {
			return err
		}
	}
	_, err = db.ExecContext(ctx, sqlDeleteWord, wordId)
	if err != nil {
		return err
	}

	if err = tx.Commit(); err != nil {
		return errors.New("failed to commit queries")
	}

	return nil
}

func AddWordToGroup(groupId, wordId int) int {
	db, _ := getDbConnection()

	defer db.Close()

	sql := `
		INSERT INTO WordGroupLinks (wordGroupId, wordId)
		VALUES ($1, $2)
		RETURNING id
	`

	id := -1
	db.QueryRow(sql, groupId, wordId).Scan(&id)

	return id
}

func RemoveWordFromGroup(groupId, wordId int) error {
	db, ctx := getDbConnection()

	defer db.Close()

	sql := `
		DELETE FROM WordGroupLinks WHERE wordGroupId=$1 AND wordId=$2
	`

	_, err := db.ExecContext(ctx, sql, groupId, wordId)

	if err != nil {
		return err
	}

	return nil
}

func GetGroups(ownerId int) ([]WordGroup, error) {
	db, _ := getDbConnection()

	defer db.Close()

	sql := `
		SELECT id, ownerId, name, description 
		FROM WordGroups WHERE ownerId=$1
	`
	groups := []WordGroup{}

	rows, err := db.Query(sql, ownerId)
	if err != nil {
		return groups, err
	}

	for rows.Next() {
		group := WordGroup{}
		err := rows.Scan(&group.Id, &group.OwnerId, &group.Name, &group.Description)
		if err != nil {
			return nil, err
		}
		groups = append(groups, group)
	}

	return groups, nil
}

func GetWordsInGroup(groupId int) ([]Word, error) {
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
}
