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

func (d *db) Aggregate(ctx context.Context) ([]service.Beverage, error) {

	var beverages []service.Beverage
	cursor, err := d.client.Database(datastore).Collection(exchangeCollection).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err := cursor.All(ctx, &beverages); err != nil {
		return nil, err
	}

	return beverages, nil

}
