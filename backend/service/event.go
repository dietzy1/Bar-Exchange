package service

import (
	"context"
	"fmt"

	"go.uber.org/zap"
)

type Event struct {
	Id       string
	Duration int64
}

type eventService struct {
	store eventStore

	logger *zap.Logger
}

type eventStore interface {
	StartEvent(ctx context.Context, req Event) (Event, error)
	StopEvent(ctx context.Context, req Event) (Event, error)
	GetEvent(ctx context.Context, req Event) (Event, error)
}

func NewEventService(store eventStore, logger *zap.Logger) (*eventService, error) {

	if store == nil {
		return nil, fmt.Errorf("EventStore is nil")
	}

	return &eventService{store: store, logger: logger}, nil
}

func (s *eventService) StartEvent(ctx context.Context, req Event) (Event, error) {
	return s.store.StartEvent(ctx, req)
}

func (s *eventService) StopEvent(ctx context.Context, req Event) (Event, error) {
	return s.store.StopEvent(ctx, req)
}

func (s *eventService) GetEvent(ctx context.Context, req Event) (Event, error) {
	return s.store.GetEvent(ctx, req)
}
