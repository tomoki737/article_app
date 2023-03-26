package controller

import (
	"database/sql"
	"encoding/json"
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

func GetSingleArticleHandler(w http.ResponseWriter, r *http.Request) {
	sub := strings.TrimPrefix(r.URL.Path, "/articles")
	_, id := filepath.Split(sub)
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
	if r.Header.Get("Content-Type") != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
	}

	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	body := make([]byte, length)
	length, err = r.Body.Read(body)

	if err != nil && err != io.EOF {
		w.WriteHeader(http.StatusInternalServerError)
	}

	var jsonBody map[string]interface{}
	err = json.Unmarshal(body[:length], &jsonBody)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	article_title := jsonBody["title"].(string)
	article_body := jsonBody["body"].(string)
	err = createArticle(article_title, article_body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
