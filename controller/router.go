package controller

import (
	"fmt"
	"net/http"
)

func SetRouter() {
	http.HandleFunc("/articles", makeHandler(articleHandler))
}

func articleHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		index(w, r)
	case "POST":
		store(w, r)
	default:
		fmt.Fprint(w, "Method not allowed.\n")
	}
}
