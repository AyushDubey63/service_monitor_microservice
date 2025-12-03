package db

import (
	"context"

	"github.com/AyushDubey63/go-monitor/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

func GetActiveServices(pool *pgxpool.Pool)([]models.MonitorService,error){
	rows, err := pool.Query(
		context.Background(),
		`SELECT id, endpoint, timeout_ms, retry_count, interval_seconds FROM services WHERE status = 'active'`,
	)
	if err!= nil{
		return  nil,err
	}
	services := []models.MonitorService{}

	for rows.Next(){
		var s models.MonitorService
		err := rows.Scan(
			&s.ID,
			&s.Endpoint,
			&s.TimeoutMS,
			&s.RetryCount,
			&s.IntervalSeconds,
		)
		if err!= nil{
			return  nil,err
		}
		services  = append(services, s)
	}
	return services,nil
}