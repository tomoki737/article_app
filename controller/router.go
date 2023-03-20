package controller

import "net/http"

func SetRouter() {
	http.HandleFunc("/", MakeHandler(index))
}
