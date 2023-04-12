package controller

import (
	"net/http"

	"app/middleware"
)

func RegisterRoutes() {
	http.HandleFunc("/articles", ArticleHandler)
	http.HandleFunc("/articles/", ArticleHandler)
	http.HandleFunc("/articles/search", SearchArticleHandler)
	http.HandleFunc("/articles/comment", CommentSaveHandler)
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/register", RegisterHandler)
	http.HandleFunc("/check", middleware.RequireLogin(CheckLoginHandler))
}
