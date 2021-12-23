package models

import (
	"errors"

	"github.com/tsa-dom/lang-trainer/app-v2/utils"
	conn "github.com/tsa-dom/lang-trainer/app/db"
)

// Tested
func (g *Groups) Create(ownerId int, group *Group) error {
	db := conn.GetDbConnection()
	defer db.Close()

	row := db.QueryRow(addGroup(), ownerId, &group.Name, &group.Description)
	if err := row.Scan(&group.Id); err != nil {
		return err
	}
	group.OwnerId = ownerId

	return nil
}

// Tested
func (g Groups) GetAll(ownerId int, groups *[]Group) error {
	db := conn.GetDbConnection()
	defer db.Close()

	rows, err := db.Query(getGroups(), ownerId)
	if err != nil {
		return err
	}

	fetchedGroups := []Group{}
	for rows.Next() {
		group := Group{}
		err := rows.Scan(&group.Id, &group.OwnerId, &group.Name, &group.Description)
		if err != nil {
			return err
		}
		fetchedGroups = append(fetchedGroups, group)
	}

	*groups = fetchedGroups

	return nil
}

// This asterix should be fixed. It doesn't work how I wanted it to work
func (g Groups) RemoveByIds(ownerId int, groupIds *[]int) error {
	db := conn.GetDbConnection()
	defer db.Close()

	tx, err := db.Begin()
	defer tx.Rollback()

	if err != nil {
		return err
	}

	_, err = tx.Exec(deleteGroupLinks(*groupIds))
	if err != nil {
		return err
	}

	rows, err := tx.Query(deleteGroups(*groupIds), ownerId)
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

	if !utils.IntArrayEquality(removed, *groupIds) {
		return errors.New("id's not match")
	}

	if err = tx.Commit(); err != nil {
		return err
	}

	*groupIds = removed

	return nil
}
