package main

import (
	"context"
	"effective_mobile_test_task/internal/app/metrics"
	"effective_mobile_test_task/internal/app/server"
	"effective_mobile_test_task/internal/app/services"
	"effective_mobile_test_task/internal/config"
	"effective_mobile_test_task/internal/infra/storage/postgres"
	"effective_mobile_test_task/internal/infra/storage/redis"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"log/slog"
	"os"
)

func main() {
	cfg, err := config.Load()

	if err != nil {
		slog.Error("Load config err", "error", err)
	}

	slog.Info("Load config success",
		slog.Any("cfg", cfg),
	)

	ctx := context.Background()
	pool, err := pgxpool.New(ctx,
		fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
			cfg.DBUser,
			cfg.DBPassword,
			cfg.DBHost,
			cfg.DBPort,
			cfg.DBName),
	)

	if err != nil {
		slog.Error("Create pool err", "error", err)
		os.Exit(1)
	}

	err = pool.Ping(context.Background())
	if err != nil {
		log.Fatal("Could not ping database")
	}

	personRepository := postgres.NewPersonRepository(pool)
	personService := services.NewPersonService(personRepository)
	metricsService := metrics.NewMetricsService()
	cache := redis_cache.NewRedisClient()
	srv := server.NewServer(personService, metricsService, cache)

	if err := srv.Start(cfg.ServerPort); err != nil {
		log.Fatalf("Server failed: %v", err)
	}
}
