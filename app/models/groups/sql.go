package models

import (
	"fmt"

	g "github.com/tsa-dom/lang-trainer/app/types"
	"github.com/tsa-dom/lang-trainer/app/utils"
)

func addGroup() string {
	return `
		INSERT INTO Groups (owner_id, name, description)
		VALUES ($1, $2, $3)
		RETURNING id
	`
}

func modifyGroup() string {
	return `
		UPDATE Groups SET name=$3, description=$4
		WHERE id=$1 AND owner_id=$2
		RETURNING id
	`
}

func addWord() string {
	return `
		INSERT INTO Words (owner_id, word, description)
		VALUES ($1, $2, $3)
		RETURNING id
	`
}

func modifyWord() string {
	return `
		UPDATE Words SET word=$3, description=$4
		WHERE id=$1 AND owner_id=$2
		RETURNING id
	`
}

// This may deprecate later
func modifyWordItem() string {
	return `
		UPDATE WordItems SET word=$3, description=$4
		WHERE id=$1 AND word_id=$2
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

func deleteGroupLinks(ids g.GroupIds) string {
	array := utils.ArrayString(ids.Ids)
	sql := fmt.Sprintf(`
		DELETE FROM GroupLinks WHERE group_id IN (%s)
	`, array)
	return sql
}

func deleteGroups(ids g.GroupIds) string {
	// I know, there is a risk for sql infjection attack, but this should be ok for int an array
	array := utils.ArrayString(ids.Ids)
	sql := fmt.Sprintf(`
		DELETE FROM Groups WHERE id IN (%s) AND owner_id=$1
		RETURNING id
	`, array)
	return sql
}

func linkWordToGroup() string {
	return `
		INSERT INTO GroupLinks (group_id, word_id)
		VALUES ($1, $2)
	`
}

func deleteWordLinks(ids []int) string {
	array := utils.ArrayString(ids)
	sql := fmt.Sprintf(`
		DELETE FROM GroupLinks WHERE word_id IN (%s)
	`, array)
	return sql
}

func deleteWords(ids []int) string {
	array := utils.ArrayString(ids)
	sql := fmt.Sprintf(`
		DELETE FROM Words WHERE id IN (%s) AND owner_id=$1
		RETURNING id
	`, array)
	return sql
}

func deleteWordItems(ids []int) string {
	array := utils.ArrayString(ids)
	sql := fmt.Sprintf(`
		DELETE FROM WordItems WHERE word_id IN (%s)
	`, array)
	return sql
}

func getGroups() string {
	return `
		SELECT id, owner_id, name, description 
		FROM Groups WHERE owner_id=$1
	`
}

func wordsInGroup() string {
	return `
		SELECT 
			W.id, W.owner_id, W.word, W.description, 
			I.id, I.word, I.description 
		FROM GroupLinks G 
		LEFT JOIN Words W ON G.group_id=$1 AND G.word_id=W.id 
		LEFT JOIN WordItems I ON W.id=I.word_id
	`
}
