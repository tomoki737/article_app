package main

import (
	"log"
	"net/http"

	"app/controller"
	// "app/models"
)

func main() {
	controller.RegisterRoutes()
	// models.StartSessionClear()
	log.Fatal(http.ListenAndServe(":8080", nil))
}
