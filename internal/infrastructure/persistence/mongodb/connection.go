package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// ConnectionConfig holds MongoDB connection configuration
type ConnectionConfig struct {
	URI             string
	Database        string
	Timeout         time.Duration
	MinPoolSize     uint64
	MaxPoolSize     uint64
	MaxConnIdleTime time.Duration
}

// Connection wraps MongoDB client and database
type Connection struct {
	Client   *mongo.Client
	Database *mongo.Database
}

// NewConnection creates a new MongoDB connection
func NewConnection(ctx context.Context, cfg ConnectionConfig) (*Connection, error) {
	// Set client options
	clientOptions := options.Client().
		ApplyURI(cfg.URI).
		SetMinPoolSize(cfg.MinPoolSize).
		SetMaxPoolSize(cfg.MaxPoolSize).
		SetMaxConnIdleTime(cfg.MaxConnIdleTime).
		SetServerSelectionTimeout(cfg.Timeout)

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// Ping the database to verify connection
	ctx, cancel := context.WithTimeout(ctx, cfg.Timeout)
	defer cancel()

	if err := client.Ping(ctx, readpref.Primary()); err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	// Get database
	database := client.Database(cfg.Database)

	return &Connection{
		Client:   client,
		Database: database,
	}, nil
}

// Close closes the MongoDB connection
func (c *Connection) Close(ctx context.Context) error {
	if c.Client != nil {
		return c.Client.Disconnect(ctx)
	}
	return nil
}

// HealthCheck performs a health check on the MongoDB connection
func (c *Connection) HealthCheck(ctx context.Context) error {
	if c.Client == nil {
		return fmt.Errorf("MongoDB client is nil")
	}

	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	if err := c.Client.Ping(ctx, readpref.Primary()); err != nil {
		return fmt.Errorf("MongoDB health check failed: %w", err)
	}

	return nil
}

// Collection returns a collection from the database
func (c *Connection) Collection(name string) *mongo.Collection {
	return c.Database.Collection(name)
}
