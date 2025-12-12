package incidentmanager

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/AyushDubey63/go-monitor/internal/db"
	"github.com/AyushDubey63/go-monitor/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func reportIncident(id uuid.UUID){

	var serverUrl = os.Getenv("SERVER_URL")
	serverUrl += `/api/v1/incident/report-incident/`
	serverUrl += id.String()
	req, err := http.NewRequest(http.MethodPost,serverUrl,nil)
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
	id,err := db.InsertIncidentLog(pool,incidentLog)
	if err != nil{
		log.Printf("Error while inserting log in db %s",err)
	}
	reportIncident(id)
}