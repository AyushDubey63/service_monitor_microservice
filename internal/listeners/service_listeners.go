package listeners

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/AyushDubey63/go-monitor/internal/models"
	"github.com/AyushDubey63/go-monitor/internal/scheduler"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

func ListenToChannel(ctx context.Context, pool *pgxpool.Pool, channel string) {
    cfg := pool.Config().ConnConfig
    conn, err := pgx.ConnectConfig(ctx, cfg)

    if err != nil {
        log.Printf("Error acquiring connection for channel %v: %v\n", channel, err)
        return
    }
    defer conn.Close(ctx)
    _, err = conn.Exec(ctx, "LISTEN "+channel)
    if err != nil {
        log.Printf("Error while listening on channel %v: %v\n", channel, err)
        return
    }

    log.Printf("Listening on channel %v...\n", channel)

    for {
        select {
        case <-ctx.Done():
            log.Printf("Listener for channel %v shutting down\n", channel)
            return
        default:
            notification, err := conn.WaitForNotification(ctx)
            if err != nil {
                log.Printf("Error while waiting for notification: %v\n", err)
                time.Sleep(time.Second)
                continue
            }
            log.Printf("Received notification: %v\n", notification)
            handleNotification(notification.Payload)
        }
    }
}

func handleNotification(payload string,){
    var data struct{
        Action string `json:"action"`
        Service models.MonitorService `json:"service"`
        ID string `json:"id"`
    }

    json.Unmarshal([]byte(payload),&data)

    switch data.Action{
    case "add","update" :
        scheduler.S.AddOrUpdateService(data.Service)
    case "delete":
        scheduler.S.RemoveService(data.ID)
    }
}
