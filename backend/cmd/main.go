package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/dietzy1/Bar-Exchange/config"
	"github.com/dietzy1/Bar-Exchange/datastore"
	"github.com/dietzy1/Bar-Exchange/server"
	"github.com/dietzy1/Bar-Exchange/service"
	"go.uber.org/zap"
)

func main() {

	logger, err := zap.NewProduction()
	if err != nil {
		logger.Fatal("failed to initialize logger", zap.Error(err))
	}

	config, err := config.New()
	if err != nil {
		logger.Warn("failed to load config", zap.Error(err))
		logger.Warn("Application will attempt to start with default config")
	}

	dbConfig := datastore.Config{
		URI:     config.DBURI,
		Timeout: time.Duration(config.DBTimeout),
	}

	store, err := datastore.New(&dbConfig)
	if err != nil {
		// Handle datastore initialization error
		logger.Fatal("failed to initialize datastore", zap.Error(err))
	}

	services, err := service.New(store, logger)
	if err != nil {
		logger.Fatal("failed to initialize services", zap.Error(err))
	}

	serverConfig := &server.Config{
		// Set your server configuration here
		Addr:   config.ServerPort,
		Logger: logger,
	}

	s := server.New(serverConfig, services)

	// Start the gRPC server in a separate goroutine
	go func() {
		if err := s.ListenAndServe(); err != nil {
			// Handle gRPC server start error
			logger.Fatal("failed to start server", zap.Error(err))
		}
	}()

	// Wait for the termination signal
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)
	<-stopChan

	// Start the graceful shutdown
	s.Stop()
	logger.Info("Application gracefully stopped")
	os.Exit(0)

}
