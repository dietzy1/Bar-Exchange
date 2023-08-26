package broker

import (
	"context"

	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

const priceUpdatesChannel = "price-updates"

// PubSub is an interface for a Redis Pub/Sub client.
//This interface should be used by the user of the package
/* type Broker interface {
	Subscribe(ctx context.Context) (*redis.PubSub, error)
	unsubscribe(ctx context.Context, pubsub *redis.PubSub) error
	Publish(ctx context.Context, message []byte) error
} */

// RedisPubSub is an implementation of the PubSub interface using the go-redis library.
type broker struct {
	client *redis.Client
	logger *zap.Logger
}

type Options struct {
	Logger *zap.Logger
	Uri    string
}

// NewRedisPubSub creates a new RedisPubSub instance.
func New(o *Options) (*broker, error) {
	otps, err := redis.ParseURL(o.Uri)
	if err != nil {

		o.Logger.Info("Failed to parse redis url", zap.Error(err))
		return nil, err
	}
	client := redis.NewClient(otps)
	if _, err := client.Ping(context.Background()).Result(); err != nil {
		o.Logger.Info("Failed to ping redis", zap.Error(err))
		return nil, err
	}

	return &broker{
		client: client,
		logger: o.Logger,
	}, nil
}

// Subscribe subscribes to a channel and returns a PubSub instance for receiving messages.
func (b *broker) Subscribe(ctx context.Context) (*redis.PubSub, error) {
	pubsub := b.client.Subscribe(ctx, priceUpdatesChannel)
	_, err := pubsub.Receive(ctx)
	if err != nil {
		b.logger.Error("Failed to subscribe to channel", zap.String("channel", priceUpdatesChannel))
		return nil, err
	}

	return pubsub, nil
}

func (b *broker) Unsubscribe(ctx context.Context, pubsub *redis.PubSub) error {
	err := pubsub.Unsubscribe(ctx)
	if err != nil {
		b.logger.Error("Failed to unsubscribe from channel")
		return err
	}

	b.logger.Info("Unsubscribed from channel", zap.String("channel", pubsub.String()))
	return nil
}

// Send message to channel
func (b *broker) Publish(ctx context.Context, message []byte) error {
	err := b.client.Publish(ctx, priceUpdatesChannel, message).Err()
	if err != nil {
		b.logger.Error("Failed to publish message to channel", zap.String("channel", priceUpdatesChannel))
		return err
	}
	return nil
}
