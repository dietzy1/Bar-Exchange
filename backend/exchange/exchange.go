package exchange

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/dietzy1/Bar-Exchange/service"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

type exchange struct {
	logger *zap.Logger

	broker Broker

	store dataStore
}

type Broker interface {
	Subscribe(ctx context.Context) (*redis.PubSub, error)
	Unsubscribe(ctx context.Context, pubsub *redis.PubSub) error
	Publish(ctx context.Context, message []byte) error
}

type dataStore interface {
	Store(ctx context.Context, req service.Beverage) error
	aggregate(ctx context.Context) ([]service.Beverage, error)
}

func New(logger *zap.Logger, broker Broker, store dataStore) *exchange {
	return &exchange{
		logger: logger,
		broker: broker,
	}
}

// Main loop which handles incoming exchange bids queue messages
func (e *exchange) recieve(ch <-chan *redis.Message) {
	for range ch {
		//Messages recieved from the recieve channel is from the user itself
		fmt.Println("Message recieved from user")
		msg, ok := <-ch
		if !ok {
			return
		}
		//convert from string to slice of bytes

		//convert msg.Payload to service.Beverage
		var payload service.Beverage
		//Unmarshal payload
		if err := json.Unmarshal([]byte(msg.Payload), &payload); err != nil {
			e.logger.Error("Failed to unmarshal payload", zap.Error(err))
			return
		}

		//Verify validity of message
		if err := service.ValidateCreateBeverage(payload); err != nil {
			e.logger.Error("Failed to validate payload", zap.Error(err))
			return
		}

		//store the message in the database
		if err := e.store.Store(context.Background(), payload); err != nil {
			e.logger.Error("Failed to store payload", zap.Error(err))
			return
		}

	}
}

func (e *exchange) aggregate() {
	//Aggregate data //TODO: I need to change the aggregate function so it only aggregates the last 5 minutes of data
	data, err := e.store.aggregate(context.Background())
	if err != nil {
		e.logger.Error("Failed to aggregate data", zap.Error(err))
		return
	}

	//Do some transforma

	//Publish aggregated data
	if err := e.broker.Publish(context.Background(), data); err != nil {
		e.logger.Error("Failed to publish aggregated data", zap.Error(err))
		return
	}

}

const (
	intervalDuration = 1 * time.Minute
)

// Function which runs at the start of an event - the goroutine will run until the event is over
func (e *exchange) Simulate(timestamp string) {

	//Here we would want to load in some configuration from the database
	//But for now we are going to skip and use constants

	//Convert timestamp to time.time
	target, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		e.logger.Error("Failed to parse timestamp", zap.Error(err))
		return
	}

	//convert target to time.Duration
	ctxDuration := target.Sub(time.Now())

	ctx, cancel := context.WithTimeout(context.Background(), ctxDuration)
	defer cancel()

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Context done. Stopping aggregator.")
				return
			default:
				e.aggregate() // Call the aggregate method on the EventAggregator instance
				time.Sleep(time.Second)
			}
		}
	}()

	//Start goroutine to aggregate data
	go e.aggregate()
}

func (e *exchange) StopSimulating() {
	//Stop goroutine to simulate price changes
	//Stop goroutine to aggregate data
}
