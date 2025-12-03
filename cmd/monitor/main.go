package main

import (
	"context"
	"fmt"
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
	services , err := db.GetActiveServices(pool)
	if err != nil{
		log.Fatalf("Error fetching the services: %v",err)
	}
	fmt.Println("Active services:")
	for _, s := range services {
		fmt.Printf("ID: %s | Endpoint: %s | Timeout: %d | Retry: %d | Interval: %d\n",
			s.ID, s.Endpoint, s.TimeoutMS, s.RetryCount, s.IntervalSeconds)
	}
	log.Println("Connected to PostgreSQL successfully")
}
