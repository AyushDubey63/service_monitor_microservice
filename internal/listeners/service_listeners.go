package listeners

import (
	"context"
	"log"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ListenToChannel(ctx context.Context, pool *pgxpool.Pool, channel string) {
    conn, err := pool.Acquire(ctx)
    if err != nil {
        log.Printf("Error acquiring connection for channel %v: %v\n", channel, err)
        return
    }
    defer conn.Release()

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
            notification, err := conn.Conn().WaitForNotification(ctx)
            if err != nil {
                log.Printf("Error while waiting for notification: %v\n", err)
                time.Sleep(time.Second)
                continue
            }
            log.Printf("Received notification: %v\n", notification)
        }
    }
}
