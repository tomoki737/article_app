package controller

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"app/controller"
	"app/database"
)

func TestGetArticles(t *testing.T) {
	db := database.GetDB()
	q, err := db.Prepare("INSERT INTO articles (title, body) VALUES (?, ?)")
	q.Exec("Test title", "Test body")

	if err != nil {
		t.Fatal(err)
	}
	r, _ := http.NewRequest("GET", "/articles", nil)
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.ArticleHandler)
	handler.ServeHTTP(w, r)

	if w.Code != 200 {
		t.Errorf("Response code is %v", w.Code)
	}
}

func TestPostArticle(t *testing.T) {

	json := `{title:"test-title",body:"test-content"}`
	jsonBytes := []byte(json)
	r, _ := http.NewRequest("POST", "/articles", bytes.NewBuffer(jsonBytes))
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.ArticleHandler)
	handler.ServeHTTP(w, r)

	if w.Code != 200 {
		t.Errorf("Response code is %v", w.Code)
	}
}
