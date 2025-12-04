package main

import (
	"context"
	"log"

	"github.com/AyushDubey63/go-monitor/internal/db"
)

func main() {

	// Connect to DB
	pool, err := db.ConnectDB(context.Background())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer pool.Close()
	log.Println("Connected to PostgreSQL successfully")
}
