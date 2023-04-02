package controller_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"app/controller"
	"app/database"
)

func TestMain(m *testing.M) {
	db := database.GetDB()
	tx, _ := db.Begin()

	_ , err := db.Exec("INSERT INTO users (name, password) VALUES (?, ?)", "test_user", "test_password")
	if err != nil {
		panic(err)
	}
	m.Run()
	defer tx.Rollback()
}

func TestLoginHandler(t *testing.T) {
	json := `{"name":"test_user","password":"test_password"}`
	jsonBytes := []byte(json)
	r, err := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonBytes))
	if err != nil {
		t.Errorf(err.Error())
	}

	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Content-Length", strconv.Itoa(len(jsonBytes)))

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.LoginHandler)
	handler.ServeHTTP(w, r)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// if _, err := w.Result().Header.Cookie("session"); err != nil {
	// 	t.Errorf("session cookie not set in response")
	// }
}
