package controller

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"strconv"
	"io/ioutil"

	"app/controller"
	"app/database"
)

func TestMain(m *testing.M) {
	db := database.GetDB()
	tx, _ := db.Begin()

	m.Run()
	defer tx.Rollback()
}

func TestGetArticles(t *testing.T) {
	r, _ := http.NewRequest("GET", "/articles", nil)
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.GetAllArticlesHandler)
	handler.ServeHTTP(w, r)

	if w.Code != 200 {
		t.Errorf("Response code is %v", w.Code)
	}
}

func TestPostArticle(t *testing.T) {
	jsonBytes := createjsonBytes()
	r, _ := http.NewRequest("POST", "/articles", bytes.NewBuffer(jsonBytes))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Content-Length", strconv.Itoa(len(jsonBytes)))
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.SaveArticleHandler)
	handler.ServeHTTP(w, r)

	resBody, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Response Body: ", string(resBody))
	if w.Code != 200 {
		t.Errorf("Response code is %v", w.Code)
	}
}

func TestPutArticle(t *testing.T) {
	jsonBytes := createjsonBytes()
	r, _ := http.NewRequest("PUT", "/articles", bytes.NewBuffer(jsonBytes))
	r.Header.Set("Content-Type", "application/json")
	r.Header.Set("Content-Length", strconv.Itoa(len(jsonBytes)))
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.SaveArticleHandler)
	handler.ServeHTTP(w, r)

	resBody, err := ioutil.ReadAll(w.Body)
	if err != nil {
		t.Fatal(err)
	}

	t.Log("Response Body: ", string(resBody))
	if w.Code != 200 {
		t.Errorf("Response code is %v", w.Code)
	}
}

func createjsonBytes() []byte {
	json := `{"title":"test-title","body":"test-content"}`
	jsonBytes := []byte(json)
	return jsonBytes
}