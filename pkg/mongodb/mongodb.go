package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	DefaultTimeout = 5 * time.Second
)

// Конструктор устанавливает связь с mongodb
func Connect(ctx context.Context, options *options.ClientOptions) (*mongo.Client, error) {
	client, err := mongo.Connect(ctx, options)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to mongodb: %w", err)
	}

	ctxTimeout, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := client.Ping(ctxTimeout, nil); err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	return client, nil
}
