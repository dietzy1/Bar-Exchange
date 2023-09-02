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

	stopSimulationChan chan struct{} // Channel to signal stop simulation

	conf configuration
}

type dataStore interface {
	Store(ctx context.Context, req service.Beverage) error
	Aggregate(ctx context.Context, timestamp string) ([]service.Beverage, error)
}

type Broker interface {
	Subscribe(ctx context.Context) (*redis.PubSub, error)
	Unsubscribe(ctx context.Context, pubsub *redis.PubSub) error
	Publish(ctx context.Context, message []byte) error
}

func New(logger *zap.Logger, broker Broker, store dataStore) (*exchange, error) {

	if broker == nil {
		return nil, fmt.Errorf("Broker is nil")
	}

	return &exchange{
		logger: logger,
		broker: broker,
		store:  store,

		stopSimulationChan: make(chan struct{}),
	}, nil
}

// Main loop which handles incoming exchange bids queue messages
func (e *exchange) recieve(ctx context.Context, ch <-chan *redis.Message) {
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

func (e *exchange) aggregate(ctx context.Context) {

	timestamp := time.Now().Format(time.RFC3339)
	//Aggregate data //TODO: I need to change the aggregate function so it only aggregates the last 5 minutes of data
	data, err := e.store.Aggregate(ctx, timestamp)
	if err != nil {
		e.logger.Error("Failed to aggregate data", zap.Error(err))
		return
	}

	newPriceData := e.calculatePrice(data)

	//json marshal the data
	jsonData, err := json.Marshal(newPriceData)
	if err != nil {
		e.logger.Error("Failed to marshal aggregated data", zap.Error(err))
		return
	}

	if err := e.broker.Publish(ctx, jsonData); err != nil {
		e.logger.Error("Failed to publish aggregated data", zap.Error(err))
		return
	}

}

// Function which runs at the start of an event - the goroutine will run until the event is over
func (e *exchange) Simulate(ctx context.Context, timestamp string) error {

	//Here we would want to load in some configuration from the database
	//But for now we are going to skip and use constants
	e.conf = *newBaseConfiguration()

	//Convert timestamp to time.time
	target, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		e.logger.Error("Failed to parse timestamp", zap.Error(err))
		return err
	}

	//convert target to time.Duration
	ctxDuration := time.Until(target)

	ctx, cancel := context.WithTimeout(context.Background(), ctxDuration)
	defer cancel()

	go func() {
		//Subscribe to the exchange bids queuex
		pubsub, err := e.broker.Subscribe(ctx)
		if err != nil {
			e.logger.Error("Failed to subscribe to exchange bids queue", zap.Error(err))
			return
		}
		for {
			select {

			case <-ctx.Done():
				e.logger.Info("Context done. Stopping recieve.")
				e.broker.Unsubscribe(ctx, pubsub)
				return

			case <-e.stopSimulationChan:
				e.logger.Info("Stop simulation signal recieved. Stopping recieve.")
				e.broker.Unsubscribe(ctx, pubsub)
				return

			default:
				//Recieve messages from the exchange bids queue
				e.recieve(ctx, pubsub.Channel())
			}
		}
	}()

	go func() {
		for {
			select {
			case <-ctx.Done():
				e.logger.Info("Context done. Stopping aggregator.")
				return

			case <-e.stopSimulationChan:
				e.logger.Info("Stop simulation signal recieved. Stopping recieve.")
				return

			default:
				e.aggregate(ctx)
				time.Sleep(e.conf.intervalDuration)
			}
		}
	}()
	return nil
}

func (e *exchange) StopSimulating() {
	//Stop goroutine to simulate price changes
	//Stop goroutine to aggregate data
	close(e.stopSimulationChan)
}
