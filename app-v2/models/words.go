package models

import (
	"errors"

	"github.com/tsa-dom/lang-trainer/app-v2/utils"
	conn "github.com/tsa-dom/lang-trainer/app/db"
)

// Tested
func (w *Words) Create(ownerId int, word *Word) error {
	db := conn.GetDbConnection()
	defer db.Close()

	tx, err := db.Begin()
	defer tx.Rollback()

	if err != nil {
		return err
	}

	err = tx.QueryRow(addWord(), ownerId, &word.Name, &word.Description).Scan(&word.Id)
	if err != nil {
		return err
	}

	wordItems := []WordItem{}
	for _, item := range word.Items {
		err := tx.QueryRow(addWordItem(), word.Id, item.Name, item.Description).Scan(&item.Id)
		if err != nil {
			return err
		}
		wordItems = append(wordItems, item)
	}
	word.Items = wordItems
	word.OwnerId = ownerId

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}

func (w *Words) Modify(ownerId int, word *Word) error {
	db := conn.GetDbConnection()
	defer db.Close()

	tx, err := db.Begin()
	defer tx.Rollback()
	if err != nil {
		return err
	}

	id := 0
	db.QueryRow(modifyWord(), &word.Id, ownerId, &word.Name, &word.Description).Scan(&id)
	if id == 0 {
		return errors.New("word modification failed, are wordId and ownerId valid ?")
	}

	modifiedItems := []WordItem{}
	for _, item := range word.Items {
		db.QueryRow(modifyWordItem(), &item.Id, &word.Id, &item.Name, &item.Description).Scan(&item.Id)
		if item.Id == 0 {
			db.QueryRow(addWordItem(), &word.Id, &item.Name, &item.Description).Scan(&item.Id)
			if err != nil {
				return err
			}
		}
		modifiedItems = append(modifiedItems, item)
	}

	word.Items = modifiedItems
	word.OwnerId = ownerId

	return nil
}

// These asterixes need some fixing
func RemoveWords(ownerId int, wordIds *[]int) error {
	db := conn.GetDbConnection()
	defer db.Close()

	tx, err := db.Begin()
	defer tx.Rollback()
	if err != nil {
		return err
	}

	_, err = db.Exec(deleteWordLinks(*wordIds))
	if err != nil {
		return err
	}

	_, err = db.Exec(deleteWordItems(*wordIds))
	if err != nil {
		return err
	}

	rows, err := db.Query(deleteWords(*wordIds), ownerId)
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

	if !utils.IntArrayEquality(removed, *wordIds) {
		return errors.New("id's not match")
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	return nil
}
