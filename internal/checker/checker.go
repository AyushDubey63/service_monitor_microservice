package checker

import (
	"net/http"
	"time"

	"github.com/AyushDubey63/go-monitor/internal/db"
	"github.com/AyushDubey63/go-monitor/internal/incidentmanager"
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

    var success bool
    var statusCode int
    var latency time.Duration
    var lastErr error

    for i := 0; i < service.RetryCount; i++ {
        start := time.Now()
        resp, err := client.Do(req)
        latency = time.Since(start)

        if err != nil {
            lastErr = err
            continue
        }

        statusCode = resp.StatusCode
        success = resp.StatusCode == service.CheckRule.ExpectedCode

        resp.Body.Close()

        if success {
            return true, statusCode, latency, nil
        }
    }

    return success, statusCode, latency, lastErr
}


func RunHealthCheck(service models.MonitorService,pool *pgxpool.Pool,removeService func()){
	_,statusCode,latency,err := checkService(service)
	status := "up" 

	if err!=nil{
		status = "down"
		incidentmanager.HandleIncident(pool,service,statusCode,latency, err)
        removeService()
	}
	serviceHealthlog := models.ServiceHealthLog{
		ServiceID: service.ID,
		Status: status,
		StatusCode: statusCode,
		LatencyMs: int(latency),
		Error: err.Error(),
	}
	db.InsertHealthLog(pool,serviceHealthlog)
}
