package scheduler

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/AyushDubey63/go-monitor/internal/checker"
	"github.com/AyushDubey63/go-monitor/internal/db"
	"github.com/AyushDubey63/go-monitor/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Scheduler struct{
    tickers map[string]*ServiceTicker
    mu sync.Mutex
    DB *pgxpool.Pool
}

type ServiceTicker struct{
    Ticker *time.Ticker
    Cancel context.CancelFunc
}

var S = &Scheduler{
    tickers: make(map[string]*ServiceTicker),
}

func StartScheduler(pool *pgxpool.Pool) {
    
    S.DB = pool

    services, err := db.GetActiveServices(pool)
    if err != nil {
        log.Fatalf("Error while fetching active services: %v", err)
    }

    for _, svc := range services {
        S.AddOrUpdateService(svc)
    }
}

func (s *Scheduler) AddOrUpdateService(service models.MonitorService,){
    s.mu.Lock()
    defer s.mu.Unlock()

    if existing, ok := s.tickers[service.ID.String()]; ok{
        existing.Cancel()
        existing.Ticker.Stop()
        delete(s.tickers,service.ID.String())
    }

    ctx, cancel := context.WithCancel(context.Background())
    ticker := time.NewTicker(time.Duration(service.IntervalSeconds) * time.Second)

    s.tickers[service.ID.String()] = &ServiceTicker{
        Ticker: ticker,
        Cancel: cancel,
    }
    go func(svc models.MonitorService){
        for {
            select{
            case <-ticker.C:
                checker.RunHealthCheck(svc,S.DB,func ()  {
                    S.RemoveService(svc.ID.String())
                })
            case <-ctx.Done():
                return
            }
        }
    }(service)
}

func (s *Scheduler) RemoveService(id string){
    s.mu.Lock()
    defer s.mu.Unlock()

    if existing, ok := s.tickers[id];ok{
        existing.Cancel()
        existing.Ticker.Stop()
        delete(s.tickers,id)
        log.Printf("Sevice %s removed from schdeuler\n",id)
    }
}