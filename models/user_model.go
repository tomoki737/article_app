package models

import (
	"net/http"

	"app/database"
	"app/utils"
)

type User struct {
	Id   string
	Name string
}

type UserError struct {
	Err  error
	Code int
}

type AuthenticatedUser struct {
	Id        int
	SessionID string
	Name      string
}

func (e *UserError) Error() string {
	return e.Err.Error()
}

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

func (u *AuthenticatedUser) GetUserBySessionID() *UserError {
	db := database.GetDB()

	getUserIDStmt, err := db.Prepare("SELECT user_id FROM sessions WHERE session_id=?")
	if err != nil {
		return &UserError{Err: err, Code: http.StatusInternalServerError}
	}
	defer getUserIDStmt.Close()

	err = getUserIDStmt.QueryRow(u.SessionID).Scan(&u.Id)

	if err != nil {
		return &UserError{Err: err, Code: http.StatusUnauthorized}
	}

	getUserNameStmt, err := db.Prepare("SELECT name FROM users WHERE id=?")
	if err != nil {
		return &UserError{Err: err, Code: http.StatusInternalServerError}
	}
	defer getUserNameStmt.Close()

	err = getUserNameStmt.QueryRow(u.Id).Scan(&u.Name)
	if err != nil {
		return &UserError{Err: err, Code: http.StatusInternalServerError}
	}
	return nil
}
