package service

import (
	"fmt"
	"strings"

	"go.uber.org/zap"
)

type domain struct {
	logger       *zap.Logger
	eventService *eventService
}

func New(event eventStore, logger *zap.Logger) (*domain, error) {

	errorBuilder := strings.Builder{}
	if event == nil {
		errorBuilder.WriteString("event store is nil")
	}

	if errorBuilder.Len() > 0 {
		return nil, fmt.Errorf("failed to initialize domain: %s", errorBuilder.String())
	}

	return &domain{
		eventService: newEventService(event, logger),
	}, nil
}
