package model

import "time"

type TelemetryMessage struct {
	SensorName string    `json:"sensor_name"`
	Value      int       `json:"value"`
	Timestamp  time.Time `json:"timestamp"`
}

type SensorConfig struct {
	Rate     int
	Name     string
	SinkAddr string
}

type SinkConfig struct {
	BindAddr      string
	LogFile       string
	BufferSize    int
	FlushInterval time.Duration
	RateLimit     int64
	EncryptKey    string
}
