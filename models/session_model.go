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

func SaveSession(sessionID string, userID int) error {
	db := database.GetDB()

	stmt, err := db.Prepare("INSERT INTO sessions (session_id, user_id) VALUES (?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(sessionID, userID)
	if err != nil {
		return err
	}

	return nil
}
