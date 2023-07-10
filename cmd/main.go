package main

import (
	_ "database/sql"
	"log"
	"os"
	"project/internal/config"
	"project/storage"

	"golang.org/x/exp/slog"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	//TODO init config
	cfg := config.MustReadConfigFile()

	//TODO init logger
	logs := setupLoger(cfg.Env)
	logs.Info("starting")

	//TODO init storage
	storage := storage.New()
	err := storage.StorageConnect()
	if err != nil {
		log.Fatal(err)
	}

	//TODO init router

	//TODO run server

}

func setupLoger(env string) *slog.Logger {
	var log *slog.Logger
	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)

	}
	return log

}
