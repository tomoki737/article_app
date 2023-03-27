package controller

import "net/http"

func SetRouter() {
	http.HandleFunc("/articles", MakeHandler(index))
}
