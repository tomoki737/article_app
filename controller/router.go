package controller

import (
	"net/http"

	"app/middleware"
)

func RegisterRoutes() {
	http.HandleFunc("/articles", middleware.MakeHandler(ArticleHandler))
	http.HandleFunc("/articles/", middleware.MakeHandler(ArticleHasIdHandler))
	http.HandleFunc("/articles/search", middleware.MakeHandler(SearchArticleHandler))
	http.HandleFunc("/login", middleware.MakeHandler(LoginHandler))
	http.HandleFunc("/register", middleware.MakeHandler(RegisterHandler))
	http.HandleFunc("/check", middleware.RequireLogin(CheckLoginHandler))
}
