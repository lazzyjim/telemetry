package sink

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"telemetry/internal/model"
)

type Server struct {
	buffer      *Buffer
	rateLimiter *RateLimiter
}

func NewServer(buffer *Buffer, rateLimit int64) *Server {
	return &Server{
		buffer:      buffer,
		rateLimiter: NewRateLimiter(rateLimit),
	}
}

func (s *Server) Handler(w http.ResponseWriter, r *http.Request) {
	data, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "cannot read body", http.StatusBadRequest)
		return
	}

	if !s.rateLimiter.TryConsume(int64(len(data))) {
		http.Error(w, "rate limit exceeded", http.StatusTooManyRequests)
		return
	}

	var msg model.TelemetryMessage
	if err := json.Unmarshal(data, &msg); err != nil {
		fmt.Println("Incorrect message:", string(data))
		http.Error(w, "bad request", http.StatusBadRequest)
		return
	}

	s.buffer.Add(msg)

	fmt.Printf("[%s] Received from %s: %d\n", msg.Timestamp.Format("15:04:05"), msg.SensorName, msg.Value)
	w.WriteHeader(http.StatusAccepted)
}
