package main

import (
	"net/http"
	"log"

	"app/controller"
)


func main() {
	controller.RegisterRoutes()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
