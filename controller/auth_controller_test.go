package controller_test

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
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

	_, err := db.Exec("INSERT INTO users (name, password) VALUES (?, ?)", "test_user", "test_password")
	if err != nil {
		panic(err)
	}
	m.Run()
	defer tx.Rollback()
}

func TestLoginHandler(t *testing.T) {
	body := map[string]interface{}{
		"name":     "test_user",
		"password": "test_password",
	}
	jsonBody, _ := json.Marshal(body)
	r, err := http.NewRequest("POST", "/login", bytes.NewBuffer(jsonBody))
	if err != nil {
		t.Errorf(err.Error())
	}

	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Content-Length", strconv.Itoa(len(jsonBody)))

	w := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.LoginHandler)
	handler.ServeHTTP(w, r)

	if status := w.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	if cookie := w.Result().Header.Get("Set-Cookie"); cookie == "" {
		t.Errorf("session cookie not set in response")
	}
}

func TestRegisterHandler(t *testing.T) {
	body := map[string]interface{}{
		"name":     "test_user",
		"password": "test_password",
	}
	jsonBody, _ := json.Marshal(body)

	r := httptest.NewRequest("POST", "/register", bytes.NewBuffer(jsonBody))

	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Content-Length", strconv.Itoa(len(jsonBody)))
	w := httptest.NewRecorder()

	controller.RegisterHandler(w, r)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d, but got %d", http.StatusOK, resp.StatusCode)
	}

	expectedBody := "Successfully registered"
	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	if string(bodyBytes) != expectedBody {
		t.Errorf("Expected response body '%s', but got '%s'", expectedBody, string(bodyBytes))
	}
}
