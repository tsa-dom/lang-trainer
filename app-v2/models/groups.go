package models

import (
	"errors"

	"github.com/tsa-dom/lang-trainer/app-v2/utils"
	conn "github.com/tsa-dom/lang-trainer/app/db"
)

func (g *Groups) Create(ownerId int, group *Group) (Group, error) {
	db := conn.GetDbConnection()
	defer db.Close()

	row := db.QueryRow(addGroup(), ownerId, &group.Name, &group.Description)
	if err := row.Scan(&group.Id); err != nil {
		return Group{}, err
	}
	group.OwnerId = ownerId

	return *group, nil
}

func (g Groups) GetAll(ownerId int) (interface{}, error) {
	db := conn.GetDbConnection()
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

// This asterix should be fixed. It doesn't work how I wanted it to work
func (g Groups) RemoveByIds(ownerId int, groupIds *[]int) ([]int, error) {
	db := conn.GetDbConnection()
	defer db.Close()

	tx, err := db.Begin()
	defer tx.Rollback()

	if err != nil {
		return []int{}, err
	}

	_, err = tx.Exec(deleteGroupLinks(*groupIds))
	if err != nil {
		return []int{}, err
	}

	rows, err := tx.Query(deleteGroups(*groupIds), ownerId)
	if err != nil {
		return []int{}, err
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
		return []int{}, errors.New("id's not match")
	}

	if err = tx.Commit(); err != nil {
		return []int{}, err
	}

	*groupIds = removed

	return *groupIds, nil
}
