package controller

import (
	"database/sql"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"

	"app/database"
)

var db *sql.DB

type Article struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func createArticle(title string, body string) error {
	db = database.GetDB()
	stmt, err := db.Prepare("INSERT INTO articles(title, body) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(title, body)
	if err != nil {
		return err
	}
	return nil
}

func editArticle(title string, body string, id string) error {
	db = database.GetDB()
	stmt, err := db.Prepare("UPDATE articles SET title = ?, body = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(title, body, id)
	if err != nil {
		return err
	}
	return nil
}

func deleteArticle(id string) error {
	db = database.GetDB()
	stmt, err := db.Prepare("DELETE FROM articles WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func getAllArticles() ([]Article, error) {
	db = database.GetDB()
	var articles []Article

	rows, err := db.Query("SELECT id, title, body FROM articles")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		article := &Article{}
		err := rows.Scan(&article.Id, &article.Title, &article.Body)
		if err != nil {
			return nil, err
		}

		articles = append(articles, Article{
			Id:    article.Id,
			Title: article.Title,
			Body:  article.Body,
		})
	}
	return articles, nil
}

func getSingleArticle(id string) (*Article, error) {
	db = database.GetDB()
	row := db.QueryRow("SELECT id, title, body FROM articles where id = ?", id)

	article := &Article{}
	err := row.Scan(&article.Id, &article.Title, &article.Body)

	if err != nil {
		return nil, err
	}
	return article, nil
}

func GetJsonBody(w http.ResponseWriter, r *http.Request) (map[string]interface{}, error) {
	if r.Header.Get("Content-Type") != "application/json" {
		return nil, errors.New("Invalid Content-Type")
	}

	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		return nil, err
	}

	body := make([]byte, length)
	length, err = r.Body.Read(body)

	if err != nil && err != io.EOF {
		return nil, err
	}

	var jsonBody map[string]interface{}
	err = json.Unmarshal(body[:length], &jsonBody)
	if err != nil {
		return nil, err
	}

	if len(jsonBody) == 0 {
		return nil, errors.New("Empty request body")
	}

	return jsonBody, nil
}

func SearchArticles(title, body string)([]Article, error) {
	db = database.GetDB()
	var articles []Article
	query := "SELECT id, title, body FROM articles WHERE 1 = 1"
	if title != "" {
		query += " AND title Like '%" + title + "%'"
	}
	if title != "" {
		query += " AND body Like '%" + body + "%'"
	}

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		article := &Article{}
		err := rows.Scan(&article.Id, &article.Title, &article.Body)
		if err != nil {
			return nil, err
		}

		articles = append(articles, Article{
			Id:    article.Id,
			Title: article.Title,
			Body:  article.Body,
		})
	}
	return articles, nil
}

func GetAllArticlesHandler(w http.ResponseWriter, r *http.Request) {
	articles, err := getAllArticles()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(articles)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func GetSingleArticleHandler(w http.ResponseWriter, r *http.Request, id string) {
	article, err := getSingleArticle(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(article)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func SaveArticleHandler(w http.ResponseWriter, r *http.Request) {
	jsonBody, err := GetJsonBody(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	article_title, ok1 := jsonBody["title"].(string)
	article_body, ok2 := jsonBody["body"].(string)

	if !ok1 || !ok2 {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = createArticle(article_title, article_body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func EditArticleHandler(w http.ResponseWriter, r *http.Request) {
	sub := strings.TrimPrefix(r.URL.Path, "/articles")
	_, id := filepath.Split(sub)
	jsonBody, err := GetJsonBody(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	article_title, ok1 := jsonBody["title"].(string)
	article_body, ok2 := jsonBody["body"].(string)

	if !ok1 || !ok2 {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = editArticle(article_title, article_body, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteArticleHandler(w http.ResponseWriter, r *http.Request) {
	sub := strings.TrimPrefix(r.URL.Path, "/articles")
	_, id := filepath.Split(sub)
	err := deleteArticle(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func SearchArticleHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	title := query.Get("title")
	body := query.Get("body")
	articles, err := SearchArticles(title, body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	data, err := json.Marshal(articles)
	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
