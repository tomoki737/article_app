package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"app/models"
	"app/utils"
)

var db *sql.DB

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetAllArticlesHandler(w, r)
	case "POST":
		SaveArticleHandler(w, r)
	default:
		fmt.Fprint(w, "Method not allowed.\n")
	}
}

func ArticleHasIdHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetArticleHandler(w, r)
	case "DELETE":
		DeleteArticleHandler(w, r)
	case "PUT":
		EditArticleHandler(w, r)
	default:
		fmt.Fprint(w, "Method not allowed.\n")
	}
}

func GetArticleHandler(w http.ResponseWriter, r *http.Request) {
	sub := strings.TrimPrefix(r.URL.Path, "/articles")
	_, id := filepath.Split(sub)

	if id != "" {
		GetSingleArticleHandler(w, r, id)
		return
	}
	GetAllArticlesHandler(w, r)
}

func GetAllArticlesHandler(w http.ResponseWriter, r *http.Request) {
	articles, err := models.GetAllArticles()
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
	article := &models.Article{Id: id}
	err := article.GetSingleArticle()
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
	jsonBody, err := utils.GetJsonBody(w, r)
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
	article := &models.Article{Title: article_title, Body: article_body}

	err = article.CreateArticle()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func EditArticleHandler(w http.ResponseWriter, r *http.Request) {
	sub := strings.TrimPrefix(r.URL.Path, "/articles")
	_, id := filepath.Split(sub)
	jsonBody, err := utils.GetJsonBody(w, r)
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
	article := &models.Article{Id: id, Title: article_title, Body: article_body}

	err = article.EditArticle()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func DeleteArticleHandler(w http.ResponseWriter, r *http.Request) {
	sub := strings.TrimPrefix(r.URL.Path, "/articles")
	_, id := filepath.Split(sub)
	article := &models.Article{Id: id}

	err := article.DeleteArticle()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}

func SearchArticleHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	title := query.Get("title")
	body := query.Get("body")
	articles, err := models.SearchArticles(title, body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	data, err := json.Marshal(articles)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

// func ArticleGetCommentsHanler(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "GET" {
// 		jsonComments, err := json.Marshal(comments)
// 		if err != nil {
// 			http.Error(w, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		w.Header().Set("Content-Type", "application/json")
// 		w.Write(jsonComments)
// 		return
// 	}

// 	http.Error(w, "", http.StatusBadRequest)
// }
