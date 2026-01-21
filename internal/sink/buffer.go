package sink

import (
	"bufio"
	"context"
	"encoding/json"
	"os"
	"sync"
	"time"

	"telemetry/internal/model"
)

type Buffer struct {
	mu            sync.Mutex
	messages      []model.TelemetryMessage
	currentSize   int
	maxSize       int
	flushInterval time.Duration
	writer        *bufio.Writer
	file          *os.File
	encryptKey    []byte // optional
	ticker        *time.Ticker
	closed        bool
}

func NewBuffer(filePath string, maxSizeBytes int, flushInterval time.Duration, encryptKey []byte) (*Buffer, error) {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	buf := &Buffer{
		messages:      make([]model.TelemetryMessage, 0),
		maxSize:       maxSizeBytes,
		writer:        bufio.NewWriter(file),
		flushInterval: flushInterval,
		encryptKey:    encryptKey,
	}

	return buf, nil
}

func (b *Buffer) Add(msg model.TelemetryMessage) error {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.closed {
		return nil
	}

	data, _ := json.Marshal(msg)
	msgSize := len(data) + 1

	b.messages = append(b.messages, msg)
	b.currentSize += msgSize

	if b.currentSize >= b.maxSize {
		b.flushLocked()
		b.currentSize = 0
	}

	return nil
}

func (b *Buffer) Run(ctx context.Context) {
	ticker := time.NewTicker(b.flushInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			b.mu.Lock()
			b.flushLocked()
			b.currentSize = 0
			b.mu.Unlock()
		case <-ctx.Done():
			b.mu.Lock()
			b.flushLocked()
			b.file.Close()
			b.mu.Unlock()
			return
		}
	}
}

func (b *Buffer) flushLocked() {
	for _, msg := range b.messages {
		data, _ := json.Marshal(msg)

		if len(b.encryptKey) == 32 {
			encrypted, err := EncryptAESGCM(b.encryptKey, data)
			if err == nil {
				data = []byte(encrypted)
			}
		}

		b.writer.Write(data)
		b.writer.WriteByte('\n')
	}
	b.writer.Flush()
	b.messages = b.messages[:0]
}

func (b *Buffer) Close() {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.closed {
		return
	}

	b.closed = true
	b.ticker.Stop()
	b.flushLocked()
	b.writer.Flush()
}
