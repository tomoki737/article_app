package main

import (
	"net/http"
	"database/sql"
	"fmt"
	"log"

	"app/database"
)

type Article struct {
	Title string
	Body  string
}

func index(w http.ResponseWriter, r *http.Request)  {
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
		fmt.Println(article)
	}
}

func main() {
	db := database.GetDB()
	http.HandleFunc("/aricles", index)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
