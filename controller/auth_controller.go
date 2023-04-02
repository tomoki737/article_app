package controller

import (
	"net/http"
	"time"
	"fmt"

	"github.com/google/uuid"

	"app/database"
	"app/models"
	"app/utils"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed.", http.StatusBadRequest)
		return
	}

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
		http.SetCookie(w, &http.Cookie{
			Name:    "session",
			Value:   sessionID,
			Expires: time.Now().Add(5 * time.Minute),
		})

		err = saveSession(sessionID, userID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "成功しました")
}

// func authenticate(name, password string) bool {
// 	db := database.GetDB()

// 	var count int
// 	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE name=? AND password=?", name, password).Scan(&count)
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	if count == 1 {
// 		return true
// 	} else {
// 		return false
// 	}
// }

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
