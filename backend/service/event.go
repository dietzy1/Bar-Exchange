package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Event struct {
	Id              string
	FutureTimeStamp string
}

type eventService struct {
	store eventStore

	logger *zap.Logger
}

type eventStore interface {
	StartEvent(ctx context.Context, req Event) error
	StopEvent(ctx context.Context, req Event) error
	GetEvent(ctx context.Context) (Event, error)
}

func NewEventService(store eventStore, logger *zap.Logger) (*eventService, error) {

	if store == nil {
		return nil, fmt.Errorf("EventStore is nil")
	}

	//Create event service
	eventService := &eventService{store: store, logger: logger}

	//use store object to check if there is an ongoing event -- To ensure fault tolerance
	event, err := store.GetEvent(context.TODO())
	if err != nil {
		logger.Info("No current event found, using default configuration")
	}
	//If a current event exists use event information to start the event again
	if event != (Event{}) {

		if err := eventService.restartEvent(context.Background(), event); err != nil {
			logger.Error("failed to restart event", zap.Error(err))
		}

	}

	return eventService, nil
}

func (s *eventService) StartEvent(ctx context.Context, req Event) (Event, error) {
	// Perform check to see if event is in the past and valid
	if err := validateTimeStamp(req.FutureTimeStamp); err != nil {
		return Event{}, err
	}

	req.Id = newID()

	//Start new event also deletes the old event
	if err := s.store.StartEvent(ctx, req); err != nil {
		return Event{}, fmt.Errorf("failed to start event: %w", err)
	}

	//Call into exchange service to start the event
	if err := s.exchange.Simulate(ctx, req); err != nil {
		return Event{}, fmt.Errorf("failed to start event: %w", err)
	}

	return req, nil
}

func (s *eventService) StopEvent(ctx context.Context, req Event) error {

	//Stop the event
	if err := s.store.StopEvent(ctx, req); err != nil {
		return fmt.Errorf("failed to stop event: %w", err)
	}

	if err := s.exchange.StopSimulating(ctx); err != nil {
		return fmt.Errorf("failed to stop event: %w", err)
	}

	return nil

}

func (s *eventService) GetEvent(ctx context.Context) (Event, error) {

	event, err := s.store.GetEvent(ctx)
	if err != nil {
		return Event{}, fmt.Errorf("failed to get event: %w", err)
	}
	return event, nil
}

func (s *eventService) restartEvent(ctx context.Context, req Event) error {
	//Perform check to see if event is in the past and valid
	if err := validateTimeStamp(req.FutureTimeStamp); err != nil {

		//If the event is invalid then delete the event from the database
		if err := s.store.StopEvent(ctx, req); err != nil {
			return fmt.Errorf("failed to stop event: %w", err)
		}

		return err
	}
	return nil
}

func validateTimeStamp(timeStamp string) error {
	if timeStamp == "" {
		return fmt.Errorf("empty future timestamp")
	}

	timestamp, err := time.Parse(time.RFC3339, timeStamp)
	if err != nil {
		return fmt.Errorf("invalid future timestamp")
	}

	//We need a function which checks if the timestamp is in the past
	if timestamp.Before(time.Now()) {
		return fmt.Errorf("timestamp is in the past")
	}
	return nil
}

func newID() string {
	return uuid.Must(uuid.NewRandom()).String()
}
