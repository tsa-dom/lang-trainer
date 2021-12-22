package models

import (
	"fmt"

	"github.com/tsa-dom/lang-trainer/app-v2/utils"
)

// Group sql queries

func getGroups() string {
	return `
		SELECT id, owner_id, name, description 
		FROM Groups WHERE owner_id=$1
	`
}

func addGroup() string {
	return `
		INSERT INTO Groups (owner_id, name, description)
		VALUES ($1, $2, $3)
		RETURNING id
	`
}

func deleteGroupLinks(ids []int) string {
	array := utils.ArrayString(ids)
	sql := fmt.Sprintf(`
		DELETE FROM GroupLinks WHERE group_id IN (%s)
	`, array)
	return sql
}

func deleteGroups(ids []int) string {
	// I know, there is a risk for sql infjection attack, but this should be ok for int an array
	array := utils.ArrayString(ids)
	sql := fmt.Sprintf(`
		DELETE FROM Groups WHERE id IN (%s) AND owner_id=$1
		RETURNING id
	`, array)
	return sql
}

// Word sql queries

func addWord() string {
	return `
		INSERT INTO Words (owner_id, word, description)
		VALUES ($1, $2, $3)
		RETURNING id
	`
}

func addWordItem() string {
	return `
		INSERT INTO WordItems (word_id, word, description)
		VALUES ($1, $2, $3)
		RETURNING id
	`
}
