package models

import (
	"app/database"
)

func GetUserId(name, password string) (int, error) {
	db := database.GetDB()
	var userID int
	err := db.QueryRow("SELECT id FROM users WHERE name=? AND password=?", name, password).Scan(&userID)
	if err != nil {
		return 0, err
	}
	return userID, nil
}
