package db

import (
	"context"

	"github.com/AyushDubey63/go-monitor/internal/models"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetActiveServices(pool *pgxpool.Pool) ([]models.MonitorService, error) {
	rows, err := pool.Query(
		context.Background(),
		`SELECT s.id,
       s.endpoint,
       s.timeout_ms,
       s.retry_count,
       s.interval_seconds
FROM services s
LEFT JOIN LATERAL (
    SELECT i.status
    FROM incidents i
    WHERE i.service_id = s.id
    ORDER BY i.started_at DESC
    LIMIT 1
) latest_incident ON true
WHERE s.status = 'active'
  AND (latest_incident.status IS NULL OR latest_incident.status = 'resolved');
`,
	)
	if err != nil {
		return nil, err
	}
	services := []models.MonitorService{}

	for rows.Next() {
		var s models.MonitorService
		err := rows.Scan(
			&s.ID,
			&s.Endpoint,
			&s.TimeoutMS,
			&s.RetryCount,
			&s.IntervalSeconds,
			&s.CheckRule.ExpectedCode,
			&s.CheckRule.ExpectedBody,
		)
		if err != nil {
			return nil, err
		}
		services = append(services, s)
	}
	return services, nil
}

func InsertHealthLog(pool *pgxpool.Pool, serviceHealthlog models.ServiceHealthLog) error {
	_, err := pool.Exec(context.Background(), `
		INSERT INTO service_health_log (service_id,status,latency_ms,error) VALUES ($1,$2,$3,$4)
	`, serviceHealthlog.ServiceID, serviceHealthlog.Status, serviceHealthlog.LatencyMs, serviceHealthlog.Error)
	if err != nil {
		return err
	}
	return nil
}

func InsertIncidentLog(pool *pgxpool.Pool, incidentLog models.Incident) (uuid.UUID,error) {
	var id uuid.UUID
	 err := pool.QueryRow(context.Background(), `
		INSERT INTO incidents (service_id,started_at,error_message,trigger_status_code,trigger_latency_ms) VALUES ($1,$2,$3,$4,$5)
	`, incidentLog.ServiceID, incidentLog.StartedAt, incidentLog.ErrorMessage, incidentLog.TriggerStatusCode, incidentLog.TriggerLatencyMS).Scan(&id)
	if err != nil {
		return uuid.Nil,err
	}
	return id,err
}
