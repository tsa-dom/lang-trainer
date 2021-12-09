package templates

import (
	"errors"

	"github.com/lib/pq"
	conn "github.com/tsa-dom/lang-trainer/app/db"
	g "github.com/tsa-dom/lang-trainer/app/types"
	"github.com/tsa-dom/lang-trainer/app/utils"
)

func CreateTemplate(ownerId int, template g.Template) (g.Template, error) {
	db := conn.GetDbConnection()
	defer db.Close()

	err := db.QueryRow(addTemplate(), ownerId, template.Name, pq.Array(template.Descriptions)).Scan(&template.Id)
	template.OwnerId = ownerId
	if err != nil {
		return g.Template{}, err
	}

	return template, nil
}

func ModifyTemplate(ownerId int, template g.Template) error {
	db := conn.GetDbConnection()
	defer db.Close()

	id := 0
	db.QueryRow(modifyTemplate(), template.Id, ownerId, template.Name, pq.Array(template.Descriptions)).Scan(&id)
	if id == 0 {
		return errors.New("template modification failed, are templateId and ownerId valid ?")
	}

	return nil
}

func RemoveTemplates(ownerId int, templateIds g.TemplateIds) error {
	db := conn.GetDbConnection()
	defer db.Close()

	tx, err := db.Begin()
	defer tx.Rollback()
	if err != nil {
		return err
	}

	rows, err := tx.Query(deleteTemplates(templateIds.Ids), ownerId)
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

	if !utils.IntArrayEquality(removed, templateIds.Ids) {
		return errors.New("id's not match")
	}

	return nil
}
