package scheduler

import (
	"context"
	"log"
	"time"

	"github.com/AyushDubey63/go-monitor/internal/db"
	"github.com/AyushDubey63/go-monitor/internal/models"
)

func StartScheduler() {
    ctx := context.Background()

    pool, err := db.ConnectDB(ctx)
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    // defer pool.Close()

    services, err := db.GetActiveServices(pool)
    if err != nil {
        log.Fatalf("Error while fetching active services: %v", err)
    }

    for _, svc := range services {
        go func(service models.MonitorService) {
            ticker := time.NewTicker(time.Duration(service.IntervalSeconds) * time.Second)
            defer ticker.Stop()

            for {
                <-ticker.C
                // call your health check function
                // runHealthCheck(service, pool)
            }
        }(svc)
    }
}
