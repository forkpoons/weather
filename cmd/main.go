package main

import (
	"context"
	"github.com/forkpoons/weather/internal/service"
	"time"
)

func main() {
	ctx, _ := context.WithCancel(context.Background())
	start(ctx)
}

func start(ctx context.Context) {
	for {
		timer := time.NewTimer(time.Minute)
		select {
		case <-ctx.Done():
			println("exit")
		case <-timer.C:
			service.GetFiveForecasts()
		}
	}
}
