package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os/signal"
	"syscall"
	"telemetry/internal/config"
	"telemetry/internal/sink"
)

func main() {
	cfg := config.ParseSinkConfig()

	var key []byte
	if cfg.EncryptKey != "" {
		if len(cfg.EncryptKey) != 32 {
			panic("encrypt-key must be exactly 32 bytes")
		}
		key = []byte(cfg.EncryptKey)
	}

	buf, err := sink.NewBuffer(cfg.LogFile, cfg.BufferSize, cfg.FlushInterval, key)
	if err != nil {
		panic(err)
	}

	server := sink.NewServer(buf, cfg.RateLimit)

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go buf.Run(ctx)

	http.HandleFunc("/ingest", server.Handler)
	srv := &http.Server{Addr: cfg.BindAddr}

	go func() {
		fmt.Printf("Telemetry sink listening on %s\n", cfg.BindAddr)
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			panic(err)
		}
	}()

	<-ctx.Done()
	buf.Close()
	fmt.Println("Shutting down...")
	srv.Shutdown(context.Background())
}
