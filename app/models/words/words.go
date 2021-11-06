package words

import (
	"context"

	_ "github.com/lib/pq"
	"github.com/tsa-dom/lang-trainer/app/models"
)

type WordItem struct {
	Id          int
	Name        string
	Description string
}

type Word struct {
	Id          int
	OwnerId     int
	Name        string
	Description string
	Items       []WordItem
}

type Group struct {
	Id          int    `json:"id"`
	OwnerId     int    `json:"ownerId"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func CreateGroup(group Group) (Group, error) {
	sql := `
		INSERT INTO Groups (owner_id, name, description)
		VALUES ($1, $2, $3)
		RETURNING id
	`
	db := models.GetDbConnection()
	defer db.Close()

	row := db.QueryRow(sql, group.OwnerId, group.Name, group.Description)
	err := row.Scan(&group.Id)
	if err != nil {
		return Group{}, err
	}

	return group, nil
}

func CreateWord(word Word) (Word, error) {
	db := models.GetDbConnection()
	defer db.Close()

	sql := `
		INSERT INTO Words (owner_id, word, description)
		VALUES ($1, $2, $3)
		RETURNING id
	`

	err := db.QueryRow(sql, word.OwnerId, word.Name, word.Description).Scan(&word.Id)

	if err != nil {
		return Word{}, err
	}

	return word, nil
}

func GetWordById(wordId int) (Word, error) {
	db := models.GetDbConnection()
	defer db.Close()

	sqlWord := `
		SELECT id, owner_id, word, description FROM Words WHERE id=$1;
	`

	sqlItems := `
		SELECT I.id, I.word, I.description FROM Words W, WordItems I Where W.id=$1 AND W.id=I.word_id;
	`
	word := Word{}

	err := db.QueryRow(sqlWord, wordId).Scan(&word.Id, &word.OwnerId, &word.Name, &word.Description)
	if err != nil {
		return Word{}, err
	}

	rows, err := db.Query(sqlItems, wordId)
	if err != nil {
		return Word{}, err
	}

	items := []WordItem{}

	for rows.Next() {
		item := WordItem{}
		err := rows.Scan(&item.Id, &item.Name, &item.Description)
		if err != nil {
			return Word{}, nil
		}
		items = append(items, item)
	}
	word.Items = items

	return word, nil
}

func AddItemsToWord(wordId int, items []WordItem) error {
	db := models.GetDbConnection()
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
	db := models.GetDbConnection()
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
	db := models.GetDbConnection()

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
	db := models.GetDbConnection()

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
	db := models.GetDbConnection()

	defer db.Close()

	sql := `
		SELECT id, owner_id, name, description 
		FROM Groups WHERE owner_id=$1
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
	db := models.GetDbConnection()
	defer db.Close()

	sql := `
		SELECT
	`
} */

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
