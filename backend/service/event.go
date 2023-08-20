package service

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

type Event struct {
	Id              string
	FutureTimeStamp string
}

type countdown struct {
	t int
	d int
	h int
	m int
	s int
}

type currentEvent struct {
	id        string
	timestamp time.Time
	timer     *time.Timer
	stopCh    chan struct{} // Channel to signal event stop
}

type eventService struct {
	store eventStore

	logger *zap.Logger

	mu        sync.RWMutex
	countdown currentEvent
}

type eventStore interface {
	StartEvent(ctx context.Context, req Event) (Event, error)
	StopEvent(ctx context.Context, req Event) (Event, error)
	GetEvent(ctx context.Context, req Event) (Event, error)
	GetEvents(ctx context.Context, req Event) ([]Event, error)
}

func NewEventService(store eventStore, logger *zap.Logger) (*eventService, error) {

	if store == nil {
		return nil, fmt.Errorf("EventStore is nil")
	}

	//use store object to check if there is an ongoing event -- To ensure fault tolerance
	events, err := store.GetEvents(context.TODO(), Event{})
	if err != nil {
		logger.Info("No current event")
	}
	//If a current event exists call into startEvent
	if len(events) > 0 {
		logger.Info("Current event exists")
		_, err := store.StartEvent(context.TODO(), Event{events[0].Id, events[0].FutureTimeStamp})
		if err != nil {
			logger.Info("Error starting event", zap.Error(err))
		}
	}

	return &eventService{store: store, logger: logger, mu: sync.RWMutex{}}, nil
}

func (s *eventService) StartEvent(ctx context.Context, req Event) (Event, error) {
	if req.FutureTimeStamp == "" {
		return Event{}, fmt.Errorf("empty future timestamp")
	}

	timestamp, err := time.Parse(time.RFC3339, req.FutureTimeStamp)
	if err != nil {
		return Event{}, fmt.Errorf("invalid future timestamp")
	}
	s.logger.Info("timestamp", zap.String("timestamp", timestamp.String()))

	//We need a function which checks if the timestamp is in the past
	if timestamp.Before(time.Now()) {
		return Event{}, fmt.Errorf("timestamp is in the past")
	}

	// Check if there is an ongoing event
	if s.countdown.id != "" {
		//Stop the ongoing event
		s.logger.Info("Stopping ongoing event")
		_, err := s.store.StopEvent(ctx, Event{s.countdown.id, s.countdown.timestamp.Format(time.RFC3339)})
		if err != nil {
			s.logger.Info("Error stopping event", zap.Error(err))
		}
		//Call StopEvent
		_, err = s.StopEvent(ctx, Event{s.countdown.id, s.countdown.timestamp.Format(time.RFC3339)})
		if err != nil {
			s.logger.Info("Error stopping event", zap.Error(err))
		}
	}

	//Now we shouldn't be leaking any goroutines

	ce := currentEvent{
		id:        uuid.Must(uuid.NewRandom()).String(),
		timestamp: timestamp,
		stopCh:    make(chan struct{}),
	}

	s.mu.Lock()
	defer s.mu.Unlock()

	s.countdown = ce

	go s.startCountdown()

	//return s.store.StartEvent(ctx, req)

	return Event{
		Id:              ce.id,
		FutureTimeStamp: req.FutureTimeStamp,
	}, nil
}

func (s *eventService) startCountdown() {

	// Get the time remaining until the event
	remaining, _ := getTimeRemaining(s.countdown.timestamp)
	s.logger.Info("Timer", zap.Int("Days", remaining.d), zap.Int("Hours", remaining.h), zap.Int("Minutes", remaining.m), zap.Int("Seconds", remaining.s))

	// Create a timer for the countdown
	s.countdown.timer = time.NewTimer(time.Duration(remaining.t) * time.Second)

	// Wait for the timer to expire
	select {
	case <-s.countdown.timer.C:
		s.logger.Info("timer expired")
	case <-s.countdown.stopCh:
		s.logger.Info("timer stopped")
	}

	// Clear the current event
	s.countdown = currentEvent{}
}

func (s *eventService) StopEvent(ctx context.Context, req Event) (Event, error) {

	s.mu.Lock()
	defer s.mu.Unlock()

	if s.countdown.id != "" && s.countdown.id == req.Id {
		// Signal the stop channel to stop the countdown
		close(s.countdown.stopCh)
		s.countdown = currentEvent{} // Clear the current event
	}

	return s.store.StopEvent(ctx, req)
}

func (s *eventService) GetEvent(ctx context.Context, req Event) (Event, error) {

	s.mu.Lock()
	defer s.mu.Unlock()
	// Check if there is an ongoing event
	if s.countdown.id == "" {
		return Event{}, fmt.Errorf("no ongoing event")
	}

	// Retrieve the ongoing event details
	if s.countdown.id == req.Id {

		remaining, rfc3339Timestamp := getTimeRemaining(s.countdown.timestamp)
		s.logger.Info("Timer", zap.Int("Days", remaining.d), zap.Int("Hours", remaining.h), zap.Int("Minutes", remaining.m), zap.Int("Seconds", remaining.s))

		return Event{
			Id:              s.countdown.id,
			FutureTimeStamp: rfc3339Timestamp,
		}, nil
	}

	return s.store.GetEvent(ctx, req)
}

//

// Helper function to get the time remaining until the event
func getTimeRemaining(t time.Time) (countdown, string) {
	currentTime := time.Now()
	difference := t.Sub(currentTime)

	total := int(difference.Seconds())
	days := int(total / (60 * 60 * 24))
	hours := int(total / (60 * 60) % 24)
	minutes := int(total/60) % 60
	seconds := int(total % 60)

	rfc3339Timestamp := t.Format(time.RFC3339)

	return countdown{
		t: total,
		d: days,
		h: hours,
		m: minutes,
		s: seconds,
	}, rfc3339Timestamp
}
