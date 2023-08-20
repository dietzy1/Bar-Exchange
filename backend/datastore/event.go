package datastore

import (
	"context"

	"github.com/dietzy1/Bar-Exchange/service"
	"go.mongodb.org/mongo-driver/bson"
)

const eventCollection = "events"

// I need this function to take a note of the timestamp and then start a timer
func (d *db) StartEvent(ctx context.Context, req service.Event) (service.Event, error) {
	_, err := d.client.Database(datastore).Collection(eventCollection).InsertOne(ctx, req)
	if err != nil {
		return service.Event{}, err
	}
	return service.Event{}, nil

}

func (d *db) StopEvent(ctx context.Context, req service.Event) (service.Event, error) {
	_, err := d.client.Database(datastore).Collection(eventCollection).DeleteOne(ctx, bson.M{"id": req.Id})
	if err != nil {
		return service.Event{}, err
	}
	return service.Event{}, nil

}

func (d *db) GetEvent(ctx context.Context, req service.Event) (service.Event, error) {

	var event service.Event
	if err := d.client.Database(datastore).Collection(eventCollection).FindOne(ctx, bson.M{"id": req.Id}).Decode(&event); err != nil {
		return service.Event{}, err
	}

	return event, nil
}

func (d *db) GetEvents(ctx context.Context, req service.Event) ([]service.Event, error) {

	var events []service.Event
	cursor, err := d.client.Database(datastore).Collection(eventCollection).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err := cursor.All(ctx, &events); err != nil {
		return nil, err
	}

	return events, nil
}
