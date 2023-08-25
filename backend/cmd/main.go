package main

import (
	"context"
	"fmt"
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

	fmt.Println("config: ", config)

	//what is 5 seconds in nano seconds? 5 * 10^9 = 5000000000

	dbConfig := datastore.Config{
		URI:     config.DBURI,
		Timeout: time.Duration(config.DBTimeout) * time.Second,
	}

	fmt.Println("dbConfig: ", dbConfig)

	store, err := datastore.New(&dbConfig)
	if err != nil {
		// Handle datastore initialization error
		logger.Fatal("failed to initialize datastore", zap.Error(err))
	}

	eventService, err := service.NewEventService(store, logger)
	if err != nil {
		logger.Fatal("failed to initialize event service", zap.Error(err))
	}

	beverageService, err := service.NewBeverageService(store, logger)
	if err != nil {
		logger.Fatal("failed to initialize beverage service", zap.Error(err))
	}

	serverConfig := &server.Config{
		// Set your server configuration here
		Addr:        config.ServerPort,
		GatewayAddr: config.GatewayPort,
		Logger:      logger,
	}

	s := server.New(serverConfig, eventService, beverageService)

	// Start the gRPC server in a separate goroutine
	go func() {
		if err := s.ListenAndServe(); err != nil {
			// Handle gRPC server start error
			logger.Fatal("failed to start server", zap.Error(err))
		}
	}()

	go func() {
		//Here I want to call the gateway server
		if err := s.RunGateway(); err != nil {
			logger.Fatal("failed to start gateway", zap.Error(err))

		}
	}()

	// Wait for the termination signal
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, os.Interrupt, syscall.SIGTERM)
	<-stopChan

	//Create new context with timeout for graceful shutdown

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Start the graceful shutdown
	s.Stop(ctx)
	logger.Info("Application gracefully stopped")
	os.Exit(0)

}
