package main

import (
	"context"
	"github.com/forkpoons/weather/internal/service"
)

func main() {
	ctx, _ := context.WithCancel(context.Background())
	s := service.Service{Ctx: ctx}
	s.Start()
}
