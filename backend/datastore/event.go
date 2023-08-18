package datastore

import (
	"context"

	"github.com/dietzy1/Bar-Exchange/service"
)

func (d *db) StartEvent(ctx context.Context, req service.Event) (service.Event, error) {
	return service.Event{}, nil
}

func (d *db) StopEvent(ctx context.Context, req service.Event) (service.Event, error) {
	return service.Event{}, nil
}

func (d *db) GetEvent(ctx context.Context, req service.Event) (service.Event, error) {
	return service.Event{}, nil
}
