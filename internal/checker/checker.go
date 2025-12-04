package checker

import (
	"net/http"
	"time"

	"github.com/AyushDubey63/go-monitor/internal/db"
	"github.com/AyushDubey63/go-monitor/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

func checkService(service models.MonitorService) (bool, int, time.Duration, error) {
    req, err := http.NewRequest(service.HttpMethod, service.Endpoint, nil)
    if err != nil {
        return false, 0, 0, err
    }

    client := &http.Client{
        Timeout: time.Duration(service.TimeoutMS) * time.Millisecond,
    }

    start := time.Now()
    resp, err := client.Do(req)
    latency := time.Since(start)

    if err != nil {
        return false, 0, latency, err
    }
    defer resp.Body.Close()

    success := resp.StatusCode == service.CheckRule.ExpectedCode

    return success, resp.StatusCode, latency, nil
}

func RunHealthCheck(service models.MonitorService,pool *pgxpool.Pool){
	success,statusCode,latency,err := checkService(service)
	status := "up" 
	if err!=nil{
		status = "down"
	}
	serviceHealthlog := models.ServiceHealthLog{
		ServiceID: service.ID,
		Status: status,
		LatencyMs: la,
	}
	db.InsertHealthLog(pool,)


}
