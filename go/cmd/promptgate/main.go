package main

import (
	"context"
	"log/slog"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/PedroMosquera/promptgate/internal/cache"
	"github.com/PedroMosquera/promptgate/internal/gateway"
	"github.com/PedroMosquera/promptgate/internal/llm"
	"github.com/PedroMosquera/promptgate/internal/ratelimit"
)

func main() {
	log := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	addr := ":" + envOr("PORT", "8080")
	rate := envFloat("RATE_PER_SEC", 5)
	burst := envFloat("BURST", 10)

	backend := llm.WithRetry(llm.NewMock(), llm.RetryOptions{MaxAttempts: 3, BaseDelay: 50 * time.Millisecond})
	gw := gateway.New(backend, cache.New(), ratelimit.New(rate, burst), log)

	srv := gateway.NewServer(addr, gw.Routes(), log)

	errCh := make(chan error, 1)
	go func() { errCh <- srv.Start() }()

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	select {
	case err := <-errCh:
		if err != nil {
			log.Error("server error", "err", err)
			os.Exit(1)
		}
	case <-ctx.Done():
		log.Info("shutdown signal received, draining")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := srv.Shutdown(shutdownCtx); err != nil {
			log.Error("graceful shutdown failed", "err", err)
			os.Exit(1)
		}
	}
	log.Info("stopped")
}

func envOr(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func envFloat(key string, def float64) float64 {
	if v := os.Getenv(key); v != "" {
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			return f
		}
	}
	return def
}
