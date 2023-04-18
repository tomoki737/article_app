package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"regexp"
	"strings"

	"app/middleware"
	"app/models"
	"app/utils"
)

var db *sql.DB

func HandleArticleRequest(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	var articleIDPattern = regexp.MustCompile(`^/articles/\d+$`)
	var isCommentPath = strings.HasPrefix(path, "/articles/") && strings.HasSuffix(path, "/comment")
	var isLikePath = strings.HasPrefix(path, "/articles/") && strings.HasSuffix(path, "/like")
	var isArticlePath = articleIDPattern.MatchString(path) || path == "/articles"

	switch {
	case isArticlePath:
		ArticleHandler(w, r)
	case isCommentPath:
		CommentHandler(w, r)
	case isLikePath:
		ArticleLikeHandler(w, r)
	default:
		http.NotFound(w, r)
	}
}

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetArticleHandler(w, r)
	case "POST":
		SaveArticleHandler(w, r)
	case "DELETE":
		DeleteArticleHandler(w, r)
	case "PUT":
		EditArticleHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func CommentHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		GetCommentHandler(w, r)
	case "POST":
		SaveCommentHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func ArticleLikeHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "POST":
		ArticleAddLikeHandler(w, r)
	case "DELETE":
		ArticleUnLikeHandler(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func GetArticleHandler(w http.ResponseWriter, r *http.Request) {
	id := utils.GetURLID(r, "/articles")

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
	id := utils.GetURLID(r, "/articles")
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
	id := utils.GetURLID(r, "/articles")
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

func SaveCommentHandler(w http.ResponseWriter, r *http.Request) {
	jsonBody, err := utils.GetJsonBody(w, r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	text, ok := jsonBody["text"].(string)
	if !ok {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = middleware.ValidateComment(text)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userID, err := middleware.GetAuthenticatedUser(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	articleIdInt, err := utils.GetURLSubID(r, 1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	articleID, err := utils.IntToUint64(articleIdInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	comment := &models.Comment{
		ArticleId: articleID,
		UserId:    userID.Id,
		Text:      text,
	}

	err = comment.SaveComment()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func GetCommentHandler(w http.ResponseWriter, r *http.Request) {
	articleIdInt, err := utils.GetURLSubID(r, 1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	articleID, err := utils.IntToUint64(articleIdInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	comments, err := models.GetComments(articleID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	data, err := json.Marshal(comments)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}

func ArticleAddLikeHandler(w http.ResponseWriter, r *http.Request) {
	articleIdInt, err := utils.GetURLSubID(r, 1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	articleID, err := utils.IntToUint64(articleIdInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := middleware.GetAuthenticatedUser(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	Like := &models.Like{UserId: user.Id, ArticleId: articleID}
	err = Like.UnLike()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = Like.AddLike()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func ArticleUnLikeHandler(w http.ResponseWriter, r *http.Request) {
	articleIdInt, err := utils.GetURLSubID(r, 1)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	articleID, err := utils.IntToUint64(articleIdInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := middleware.GetAuthenticatedUser(r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	Like := &models.Like{UserId: user.Id, ArticleId: articleID}
	err = Like.UnLike()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
