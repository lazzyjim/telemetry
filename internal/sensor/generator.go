package sensor

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"telemetry/internal/model"
)

func StartSensor(ctx context.Context, name string, rate int, sinkURL string) {
	interval := time.Second / time.Duration(rate)
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	client := &http.Client{
		Timeout: 3 * time.Second,
	}

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Sensor stopped gracefully")
			return
		case <-ticker.C:
			msg := model.TelemetryMessage{
				SensorName: name,
				Value:      rand.Intn(100),
				Timestamp:  time.Now().UTC(),
			}
			err := sendMessage(client, sinkURL, msg)
			if err != nil {
				fmt.Println("Failed to send message:", err)
			} else {
				// 🔹 Логируем успешную отправку
				fmt.Printf("[%s] Sent value %d to %s\n", msg.Timestamp.Format("15:04:05"), msg.Value, sinkURL)
			}
		}
	}
}

func sendMessage(client *http.Client, url string, msg model.TelemetryMessage) error {
	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	for i := 0; i < 3; i++ {
		resp, err := client.Post(url, "application/json", bytes.NewBuffer(data))
		if err == nil && resp.StatusCode < 400 {
			resp.Body.Close()
			return nil
		}

		if resp != nil {
			resp.Body.Close()
		}

		backoff := time.Duration(1<<i) * 100 * time.Millisecond
		time.Sleep(backoff)
	}

	return fmt.Errorf("failed to send message after retries")
}
