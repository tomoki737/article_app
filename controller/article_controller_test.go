package controller

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

	db.Exec("INSERT INTO articles (title, body) VALUES (?, ?)", "test titles", "test bodys")
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

func TestSearchArticleHandler(t *testing.T) {
	r, _ := http.NewRequest(http.MethodGet, "/articles/search?title=test&body=test", nil)
	w := httptest.NewRecorder()
	handler := http.HandlerFunc(controller.SearchArticleHandler)
	handler.ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("unexpected status code: got %v want %v", w.Code, http.StatusOK)
		t.Log(w)
	}

	type Article struct {
		Id    string `json:"id"`
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	type Articles struct {
		ArticleList []Article
	}

	var articles []Article

	if err := json.NewDecoder(w.Body).Decode(&articles); err != nil {
		t.Errorf("failed to decode response body: %v", err)
	}

	if len(articles) == 0 {
		t.Errorf("no articles found")
	}

	// expected := Article{Id: "1", Title: "test title", Body: "test body"}
	// if articles[0] != expected {
	// 		t.Errorf("unexpected response: got %v want %v", articles[0], expected)
	// }
}
