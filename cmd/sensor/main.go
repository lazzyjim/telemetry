package main

import (
	"context"
	"fmt"
	"os/signal"
	"syscall"
	"telemetry/internal/config"
	"telemetry/internal/sensor"
)

func main() {
	cfg := config.ParseSensorConfig()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	fmt.Printf("Starting sensor '%s' at %d msgs/sec → %s\n", cfg.Name, cfg.Rate, cfg.SinkAddr)
	sensor.StartSensor(ctx, cfg.Name, cfg.Rate, cfg.SinkAddr)
}
