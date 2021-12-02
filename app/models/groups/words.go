package models

import (
	"context"

	_ "github.com/lib/pq"
	conn "github.com/tsa-dom/lang-trainer/app/db"
	g "github.com/tsa-dom/lang-trainer/app/types"
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
