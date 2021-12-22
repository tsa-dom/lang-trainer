package models

import (
	conn "github.com/tsa-dom/lang-trainer/app/db"
)

func (w *Words) Create(ownerId int, word *Word) (Word, error) {
	db := conn.GetDbConnection()
	defer db.Close()

	tx, err := db.Begin()
	defer tx.Rollback()

	if err != nil {
		return Word{}, err
	}

	err = db.QueryRow(addWord(), ownerId, &word.Name, &word.Description).Scan(&word.Id)
	if err != nil {
		return Word{}, err
	}

	wordItems := []WordItem{}
	for _, item := range word.Items {
		err := tx.QueryRow(addWordItem(), word.Id, item.Name, item.Description).Scan(&item.Id)
		if err != nil {
			return Word{}, err
		}
		wordItems = append(wordItems, item)
	}
	word.Items = wordItems

	if err = tx.Commit(); err != nil {
		return Word{}, err
	}

	return *word, err
}
