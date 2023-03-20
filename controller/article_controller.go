package controller

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"app/database"
)

var db *sql.DB

type Article struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func index(w http.ResponseWriter, r *http.Request) {
	db = database.GetDB()

	var buf bytes.Buffer
	var articles []Article

	enc := json.NewEncoder(&buf)
	rows, err := db.Query("SELECT title, body FROM articles")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	for rows.Next() {
		article := &Article{}
		if err := rows.Scan(&article.Title, &article.Body); err != nil {
			log.Fatal(err)
		}
		articles = append(articles, Article{
			Title: article.Title,
			Body:  article.Body,
		})
	}
	enc.Encode(&articles)
	fmt.Fprintf(w, buf.String())
}
