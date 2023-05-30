package models

import (
	"carilokak/configs"
)

func CreateObject(object Objects) (Objects, error) {
	db := configs.GetDB()
	sqlStatement := `INSERT INTO objects (object_name,gender_id) VALUES ($1, $2) RETURNING object_id`
	err := db.QueryRow(sqlStatement, object.ObjectName, object.ObjectGender).Scan(&object.ObjectId)
	if err != nil {
		return object, err
	}
	return object, nil
}
