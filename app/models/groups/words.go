package groups

import (
	"context"

	_ "github.com/lib/pq"
	"github.com/tsa-dom/lang-trainer/app/models"
)

func CreateGroup(group Group) (Group, error) {
	db := models.GetDbConnection()
	defer db.Close()

	row := db.QueryRow(addGroup(), group.OwnerId, group.Name, group.Description)
	err := row.Scan(&group.Id)
	if err != nil {
		return Group{}, err
	}

	return group, nil
}

func CreateWord(word Word) (Word, error) {
	db := models.GetDbConnection()
	defer db.Close()

	err := db.QueryRow(addWord(), word.OwnerId, word.Name, word.Description).Scan(&word.Id)

	if err != nil {
		return Word{}, err
	}

	return word, nil
}

func GetWordById(wordId int) (Word, error) {
	db := models.GetDbConnection()
	defer db.Close()

	word := Word{}

	err := db.QueryRow(wordById(), wordId).Scan(&word.Id, &word.OwnerId, &word.Name, &word.Description)
	if err != nil {
		return Word{}, err
	}

	rows, err := db.Query(wordItemsByWordId(), wordId)
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

	for _, item := range items {
		_, err := tx.Exec(addWordItem(), wordId, item.Name, item.Description)
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
	db := models.GetDbConnection()

	defer db.Close()

	_, err := db.Exec(linkWordToGroup(), groupId, wordId)
	if err != nil {
		return err
	}

	return nil
}

func RemoveWordFromGroup(groupId, wordId int) error {
	db := models.GetDbConnection()

	defer db.Close()

	_, err := db.Exec(deleteWordLink(), groupId, wordId)
	if err != nil {
		return err
	}

	return nil
}

func GetGroups(ownerId int) ([]Group, error) {
	db := models.GetDbConnection()

	defer db.Close()

	rows, err := db.Query(getGroups(), ownerId)
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
