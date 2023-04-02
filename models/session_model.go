package models

import (
	"fmt"
	"time"

	"app/database"
)

func StartSessionClear() {
	for {
		fmt.Println("Running session cleaner...")
		cleanSessions()
		time.Sleep(24 * time.Hour)
	}
}

func cleanSessions() {
	db := database.GetDB()

	_, err := db.Exec("DELETE FROM sessions WHERE created_at < DATE_SUB(NOW(), INTERVAL 1 WEEK)")
	if err != nil {
		panic(err)
	}
	fmt.Println("Session cleaner finished.")
}
