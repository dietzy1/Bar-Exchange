package service

import (
	"context"

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
}

func newEventService(store eventStore, logger *zap.Logger) *eventService {
	return &eventService{store: store}
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
