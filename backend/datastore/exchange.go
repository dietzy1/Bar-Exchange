package datastore

import (
	"context"

	"github.com/dietzy1/Bar-Exchange/service"
	"go.mongodb.org/mongo-driver/bson"
)

const exchangeCollection = "exchange"

func (d *db) Store(ctx context.Context, req service.Beverage) error {

	_, err := d.client.Database(datastore).Collection(exchangeCollection).InsertOne(ctx, req)
	if err != nil {
		return err
	}
	return nil

}

// Fucntion which aggregates all of the purchases within a given time frame of 5 minutes of the timestamp
func (d *db) Aggregate(ctx context.Context, timestamp string) ([]service.Beverage, error) {

	//Create a slice of service.Beverage
	var beverages []service.Beverage
	//Create a filter to find all of the purchases within a given time frame
	filter := bson.M{"timestamp": bson.M{"$gte": timestamp, "$lt": timestamp}}

	//Find all of the purchases within a given time frame
	cursor, err := d.client.Database(datastore).Collection(exchangeCollection).Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	//Iterate through the cursor and append each beverage to the slice
	for cursor.Next(ctx) {
		var beverage service.Beverage
		if err := cursor.Decode(&beverage); err != nil {
			return nil, err
		}
		beverages = append(beverages, beverage)
	}
	return beverages, nil

}
