package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"

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

	err = models.SaveSession(sessionID, userID)
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

	userID, err := models.CreateUser(name, password)
	id := int(userID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	sessionID := uuid.NewString()

	setSession(w, sessionID)

	err = models.SaveSession(sessionID, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Successfully registered")
}

func setSession(w http.ResponseWriter, sessionID string) {
	http.SetCookie(w, &http.Cookie{
		Name:    "session",
		Value:   sessionID,
		Expires: time.Now().Add(5 * time.Minute),
	})
}

func CheckLoginHandler(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user")
	user, ok := user.(*models.AuthenticatedUser)
	if !ok {
		http.Error(w, "token not found", http.StatusInternalServerError)
		return
	}
	json, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "user not found in context", http.StatusInternalServerError)
		return
	}
	w.Write(json)
}
