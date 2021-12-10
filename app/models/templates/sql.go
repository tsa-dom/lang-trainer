package templates

import (
	"fmt"

	"github.com/tsa-dom/lang-trainer/app/utils"
)

func addTemplate() string {
	return `
		INSERT INTO Templates (owner_id, name, descriptions)
		VALUES ($1, $2, $3)
		RETURNING id
	`
}

func modifyTemplate() string {
	return `
		UPDATE Templates SET name=$3, descriptions=$4
		WHERE id=$1 AND owner_id=$2
		RETURNING id
	`
}

func deleteTemplates(ids []int) string {
	array := utils.ArrayString(ids)
	sql := fmt.Sprintf(`
		DELETE FROM Templates WHERE id IN (%s) AND owner_id=$1
		RETURNING id
	`, array)
	return sql
}

func getTemplates() string {
	return `
		SELECT id, name, descriptions FROM Templates WHERE owner_id=$1
	`
}
