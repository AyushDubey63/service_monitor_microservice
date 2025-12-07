package incidentmanager

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/AyushDubey63/go-monitor/internal/db"
	"github.com/AyushDubey63/go-monitor/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

func reportIncident(incident models.Incident){
	
	body  := map[string]any{
		"service_id" :incident.ServiceID,
		"status" : incident.Status,
		"error_message" : incident.ErrorMessage,
		"trigger_status_code" : incident.TriggerStatusCode,
		"trigger_latency_ms" : incident.TriggerLatencyMS,
		"started_at" : incident.StartedAt,
	} 
	jsonData, err := json.Marshal(body)
	if err != nil {
		fmt.Printf("Error generating json body: %v",err)
	}
	var serverUrl = os.Getenv("SERVER_URL")
	serverUrl += `/api/v1/incident/report-incident`
	req, err := http.NewRequest("POST",serverUrl,bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Error while creating http request: %v",err)
	}

	req.Header.Set("Content-Type","application/json")
	client := &http.Client{}
	res, err := client.Do(req)
	if err!= nil {
		fmt.Print("Error while reporting ", err)
	}
	defer res.Body.Close()
}

func HandleIncident(pool *pgxpool.Pool,service models.MonitorService,statusCode int,latency time.Duration,err error){
	incidentLog := models.Incident{
			ServiceID:  service.ID,
			Status: "open",
			StartedAt: time.Now(),
			ErrorMessage: err.Error(),
			TriggerStatusCode: statusCode,
			TriggerLatencyMS: int(latency),
		}
	db.InsertIncidentLog(pool,incidentLog)
	reportIncident(incidentLog)
}