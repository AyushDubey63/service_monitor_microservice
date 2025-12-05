package incidentmanager

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/AyushDubey63/go-monitor/internal/db"
	"github.com/AyushDubey63/go-monitor/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

func reportIncident(incident models.Incident){
	var serverUrl = os.Getenv("SERVER_URL")
	serverUrl += `/api/v1/incident/report-incident/`
	serverUrl += incident.ID.String()
	client := &http.Client{}
	_, err := client.Get(serverUrl)
	if err!= nil {
		fmt.Print("Error while reporting ", err)
	}
}

func HandleIncident(pool *pgxpool.Pool,service models.MonitorService,statusCode int,latency time.Duration,err error){
	incidentLog := models.Incident{
			ServiceID:  service.ID,
			StartedAt: time.Now(),
			ErrorMessage: err.Error(),
			TriggerStatusCode: statusCode,
			TriggerLatencyMS: int(latency),
		}
	db.InsertIncidentLog(pool,incidentLog)
	reportIncident(incidentLog)
}