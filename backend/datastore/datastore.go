package datastore

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const datastore = "Bar-Exchange"

type db struct {
	client *mongo.Client
}

type Config struct {
	URI     string
	Timeout time.Duration
}

func New(config *Config) (*db, error) {

	if config.Timeout == 0 {
		config.Timeout = 5 * time.Second
	}

	if config.URI == "" {
		return nil, fmt.Errorf("URI is empty, missing connection string")
	}

	ctx, cancel := context.WithTimeout(context.Background(), config.Timeout)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.URI))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongo: %v", err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		return nil, fmt.Errorf("failed to ping mongo: %v", err)
	}
	c := &db{client: client}
	return c, nil

}
