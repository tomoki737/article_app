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

func store(w http.ResponseWriter, r *http.Request) {
	db = database.GetDB()

	title := r.FormValue("title")
	body := r.FormValue("body")

	stmt, err := db.Prepare("INSERT INTO articles(title, body) VALUES (?, ?)")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer stmt.Close()

	_, err = stmt.Exec(title, body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	fmt.Fprintf(w,"成功です")
}
