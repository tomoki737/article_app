package models

import (
	"app/database"
	"app/utils"
)

func GetUserId(name, password string) (int, error) {
	password = utils.HashString(password)
	db := database.GetDB()
	var userID int
	err := db.QueryRow("SELECT id FROM users WHERE name=? AND password=?", name, password).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}
