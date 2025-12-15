package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	// PersonsCollection is the name of the persons collection
	PersonsCollection = "persons"
)

// InitializeIndexes creates all required database indexes
func InitializeIndexes(ctx context.Context, db *mongo.Database) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	// Initialize persons collection indexes
	if err := initPersonsIndexes(ctx, db); err != nil {
		return fmt.Errorf("failed to initialize persons indexes: %w", err)
	}

	return nil
}

// initPersonsIndexes creates indexes for the persons collection
func initPersonsIndexes(ctx context.Context, db *mongo.Database) error {
	collection := db.Collection(PersonsCollection)

	indexes := []mongo.IndexModel{
		{
			// Index on created_at for sorting and pagination
			Keys: bson.D{{Key: "created_at", Value: -1}},
			Options: options.Index().
				SetName("idx_person_created_at"),
		},
		{
			// Compound index on created_at and _id for efficient pagination
			Keys: bson.D{
				{Key: "created_at", Value: -1},
				{Key: "_id", Value: 1},
			},
			Options: options.Index().
				SetName("idx_person_created_id"),
		},
	}

	// Create indexes
	_, err := collection.Indexes().CreateMany(ctx, indexes)
	if err != nil {
		return fmt.Errorf("failed to create indexes: %w", err)
	}

	return nil
}

// DropIndexes drops all custom indexes (useful for testing)
func DropIndexes(ctx context.Context, db *mongo.Database) error {
	ctx, cancel := context.WithTimeout(ctx, 10*time.Second)
	defer cancel()

	collection := db.Collection(PersonsCollection)

	// Drop all indexes except _id
	_, err := collection.Indexes().DropAll(ctx)
	if err != nil {
		return fmt.Errorf("failed to drop indexes: %w", err)
	}

	return nil
}
