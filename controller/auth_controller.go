package controller

import (
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"

	"app/database"
	"app/models"
	"app/utils"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	jsonBody, err := utils.GetJsonBody(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	name, ok1 := jsonBody["name"].(string)
	password, ok2 := jsonBody["password"].(string)

	if !ok1 || !ok2 {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userID, err := models.GetUserId(name, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	sessionID := uuid.NewString()
	setSession(w, sessionID)

	err = saveSession(sessionID, userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Successfully logged in")
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	jsonBody, err := utils.GetJsonBody(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	name, ok1 := jsonBody["name"].(string)
	password, ok2 := jsonBody["password"].(string)

	if !ok1 || !ok2 {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	userID, err := createUser(name, password)
	id := int(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sessionID := uuid.NewString()

	setSession(w, sessionID)

	err = saveSession(sessionID, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Successfully registered")
}

func saveSession(sessionID string, userID int) error {
	db := database.GetDB()

	stmt, err := db.Prepare("INSERT INTO sessions (session_id, user_id) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(sessionID, userID)
	if err != nil {
		return err
	}

	return nil
}

func setSession(w http.ResponseWriter, sessionID string) {
	http.SetCookie(w, &http.Cookie{
		Name:    "session",
		Value:   sessionID,
		Expires: time.Now().Add(5 * time.Minute),
	})
}

func createUser(name, password string) (int64, error) {
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
