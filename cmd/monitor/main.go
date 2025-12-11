package main

import (
	"context"
	"log"

	"github.com/AyushDubey63/go-monitor/internal/db"
	"github.com/AyushDubey63/go-monitor/internal/listeners"
	"github.com/AyushDubey63/go-monitor/internal/scheduler"
)

func main() {

	// Connect to DB
	pool, err := db.ConnectDB(context.Background())
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	scheduler.StartScheduler(pool);
	listeners.ListenToChannel(context.Background(),pool,"service_changes")
	defer pool.Close()
	log.Println("Connected to PostgreSQL successfully")
}
