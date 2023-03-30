package controller

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"app/middleware"
)

func RegisterRoutes() {
	http.HandleFunc("/articles", middleware.MakeHandler(ArticleHandler))
	http.HandleFunc("/articles/", middleware.MakeHandler(ArticleHasIdHandler))
	http.HandleFunc("/articles/search", middleware.MakeHandler(SearchArticleHandler))
}

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