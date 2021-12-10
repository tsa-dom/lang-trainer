package models

import (
	"context"
	"errors"

	_ "github.com/lib/pq"
	conn "github.com/tsa-dom/lang-trainer/app/db"
	g "github.com/tsa-dom/lang-trainer/app/types"
	"github.com/tsa-dom/lang-trainer/app/utils"
)

// Created a new word
func CreateWord(word g.Word) (g.Word, error) {
	db := conn.GetDbConnection()
	defer db.Close()

	err := db.QueryRow(addWord(), word.OwnerId, word.Name, word.Description).Scan(&word.Id)

	if err != nil {
		return g.Word{}, err
	}

	return word, nil
}

func ModifyWord(ownerId int, word g.Word) error {
	db := conn.GetDbConnection()
	defer db.Close()

	tx, err := db.Begin()
	defer tx.Rollback()
	if err != nil {
		return err
	}

	id := 0
	db.QueryRow(modifyWord(), word.Id, ownerId, word.Name, word.Description).Scan(&id)
	if id == 0 {
		return errors.New("word modification failed, are wordId and ownerId valid ?")
	}

	for _, item := range word.Items {
		id = 0
		db.QueryRow(modifyWordItem(), item.Id, word.Id, item.Name, item.Description).Scan(&id)
		if id == 0 {
			_, err = db.Exec(addWordItem(), word.Id, item.Name, item.Description)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func RemoveWords(ownerId int, wordIds g.WordIds) error {
	db := conn.GetDbConnection()
	defer db.Close()

	tx, err := db.Begin()
	defer tx.Rollback()
	if err != nil {
		return err
	}

	_, err = db.Exec(deleteWordLinks(wordIds.Ids))
	if err != nil {
		return err
	}

	_, err = db.Exec(deleteWordItems(wordIds.Ids))
	if err != nil {
		return err
	}

	rows, err := db.Query(deleteWords(wordIds.Ids), ownerId)
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

	if !utils.IntArrayEquality(removed, wordIds.Ids) {
		return errors.New("id's not match")
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
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
