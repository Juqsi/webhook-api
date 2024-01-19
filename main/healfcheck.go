package main

import (
	"database/sql"
	"fmt"
	"time"
)

func dbReachable(db *sql.DB) {

	status := true

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C
		if err := db.Ping(); err != nil && status {
			fmt.Println("❌  database connection -> " + err.Error())
			status = false
		} else if !status {
			fmt.Println("✅  database connected")
			status = true
		}
	}
}
