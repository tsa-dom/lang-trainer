package db

import (
	"errors"
	"log"

	_ "github.com/lib/pq"
)

type WordItem struct {
	Word        string
	Description string
}

type Word struct {
	OwnerId      int
	BaseWordItem WordItem
	LinkedItems  []WordItem
}

type WordGroup struct {
	OwnerId     int
	Name        string
	Description string
}

func CreateWordGroup(group WordGroup) int {
	sql := `
		INSERT INTO Groups (ownerId, name, description)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	db, _ := getDbConnection()
	id := -1
	db.QueryRow(sql, group.OwnerId, group.Name, group.Description).Scan(&id)
	db.Close()

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

func DeleteWord(wordId int) error {
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
