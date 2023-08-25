package datastore

import (
	"context"

	"github.com/dietzy1/Bar-Exchange/service"
	"go.mongodb.org/mongo-driver/bson"
)

const beverageCollection = "beverages"

func (d *db) GetBeverages(ctx context.Context) ([]service.Beverage, error) {

	var beverages []service.Beverage
	cursor, err := d.client.Database(datastore).Collection(beverageCollection).Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	if err := cursor.All(ctx, &beverages); err != nil {
		return nil, err
	}

	return beverages, nil

}

func (d *db) CreateBeverage(ctx context.Context, req service.Beverage) error {

	_, err := d.client.Database(datastore).Collection(beverageCollection).InsertOne(ctx, req)
	if err != nil {
		return err
	}
	return nil

}

func (d *db) UpdateBeverage(ctx context.Context, req service.Beverage) error {

	_, err := d.client.Database(datastore).Collection(beverageCollection).UpdateOne(ctx, bson.M{"id": req.Id}, bson.M{"$set": req})
	if err != nil {
		return err
	}
	return nil

}

func (d *db) DeleteBeverage(ctx context.Context, req service.Beverage) error {

	_, err := d.client.Database(datastore).Collection(beverageCollection).DeleteOne(ctx, bson.M{"id": req.Id})
	if err != nil {
		return err
	}
	return nil

}
