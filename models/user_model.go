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

func CreateUser(name, password string) (int64, error) {
	password = utils.HashString(password)
	db := database.GetDB()
	stmt, err := db.Prepare("INSERT INTO users (name, password) VALUES (?, ?)")
	defer stmt.Close()
	result, err := stmt.Exec(name, password)
	if err != nil {
		return 0, err
	}

	userID, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return userID, err
}
