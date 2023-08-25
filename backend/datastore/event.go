package datastore

import (
	"context"

	"github.com/dietzy1/Bar-Exchange/service"
	"go.mongodb.org/mongo-driver/bson"
)

const eventCollection = "events"

// I need this function to take a note of the timestamp and then start a timer
func (d *db) StartEvent(ctx context.Context, req service.Event) error {

	//check if event already exists
	existingEvent, _ := d.GetEvent(ctx)

	if existingEvent.Id != "" {
		//Delete existing events
		if err := d.client.Database(datastore).Collection(eventCollection).Drop(ctx); err != nil {
			return err
		}
	}

	_, err := d.client.Database(datastore).Collection(eventCollection).InsertOne(ctx, req)
	if err != nil {
		return err
	}
	return nil

}

func (d *db) StopEvent(ctx context.Context, req service.Event) error {
	_, err := d.client.Database(datastore).Collection(eventCollection).DeleteOne(ctx, bson.M{"id": req.Id})
	if err != nil {
		return err
	}
	return nil

}

func (d *db) GetEvent(ctx context.Context) (service.Event, error) {

	var event service.Event
	if err := d.client.Database(datastore).Collection(eventCollection).FindOne(ctx, bson.M{}).Decode(&event); err != nil {
		return service.Event{}, err
	}

	return event, nil
}
