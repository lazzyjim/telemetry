package config

import (
	"flag"
	"telemetry/internal/model"
	"time"
)

func ParseSensorConfig() model.SensorConfig {
	var cfg model.SensorConfig
	flag.IntVar(&cfg.Rate, "rate", 1, "Messages per second")
	flag.StringVar(&cfg.Name, "name", "sensor1", "Sensor name")
	flag.StringVar(&cfg.SinkAddr, "sink", "http://localhost:8080/ingest", "Telemetry sink address")
	flag.Parse()
	return cfg
}

func ParseSinkConfig() model.SinkConfig {
	var cfg model.SinkConfig
	flag.StringVar(&cfg.BindAddr, "bind", ":8080", "Bind address for HTTP server")
	flag.StringVar(&cfg.LogFile, "file", "telemetry.log", "Path to log file")
	flag.IntVar(&cfg.BufferSize, "buffer", 1024*16, "Buffer size in bytes")
	flag.DurationVar(&cfg.FlushInterval, "flush", 100*time.Millisecond, "Buffer flush interval")
	flag.Int64Var(&cfg.RateLimit, "rate-limit", 1024*1024, "Max input rate in bytes/sec")
	flag.StringVar(&cfg.EncryptKey, "encrypt-key", "", "32-byte encryption key (optional)")
	flag.Parse()
	return cfg
}
