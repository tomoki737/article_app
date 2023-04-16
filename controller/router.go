package controller

import (
	"net/http"

	"app/middleware"
)

func RegisterRoutes() {
	http.HandleFunc("/articles", HandleArticleRequest)
	http.HandleFunc("/articles/", HandleArticleRequest)
	http.HandleFunc("/articles/search", SearchArticleHandler)
	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/register", RegisterHandler)
	http.HandleFunc("/check", middleware.RequireLogin(CheckLoginHandler))
}
