package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"app/database"
)

type Article struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

var db *sql.DB

func index(w http.ResponseWriter, r *http.Request) {
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

func makeHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Headers", "*")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		fn(w, r)
	}
}

func main() {
	db = database.GetDB()
	http.HandleFunc("/", makeHandler(index))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
