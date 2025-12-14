package mongodb

import (
	"context"
	"fmt"

	"github.com/BrunnoQ/cadastro-pessoas/internal/domain/entities"
	"github.com/BrunnoQ/cadastro-pessoas/internal/domain/repositories"
	"github.com/BrunnoQ/cadastro-pessoas/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoPersonRepository implements PersonRepository using MongoDB
type MongoPersonRepository struct {
	collection *mongo.Collection
}

// NewMongoPersonRepository creates a new MongoPersonRepository
func NewMongoPersonRepository(db *mongo.Database) repositories.PersonRepository {
	return &MongoPersonRepository{
		collection: db.Collection(PersonsCollection),
	}
}

// Create creates a new person in the database
func (r *MongoPersonRepository) Create(ctx context.Context, person *entities.Person) error {
	_, err := r.collection.InsertOne(ctx, person)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return errors.NewConflictError("person already exists")
		}
		return fmt.Errorf("failed to insert person: %w", err)
	}
	return nil
}

// FindByID finds a person by ID
func (r *MongoPersonRepository) FindByID(ctx context.Context, id primitive.ObjectID) (*entities.Person, error) {
	var person entities.Person
	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&person)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.NewNotFoundError("person not found")
		}
		return nil, fmt.Errorf("failed to find person: %w", err)
	}
	return &person, nil
}

// Update updates an existing person with optimistic locking
func (r *MongoPersonRepository) Update(ctx context.Context, person *entities.Person) error {
	// Save current version for optimistic locking check
	currentVersion := person.Version

	// Increment version for next update
	person.IncrementVersion()

	// Use optimistic locking with version field
	filter := bson.M{
		"_id":     person.ID,
		"version": currentVersion, // Check current version
	}

	result, err := r.collection.ReplaceOne(ctx, filter, person)
	if err != nil {
		return fmt.Errorf("failed to update person: %w", err)
	}

	if result.MatchedCount == 0 {
		return errors.NewConcurrencyError("person was modified by another process")
	}

	return nil
}

// Delete deletes a person by ID
func (r *MongoPersonRepository) Delete(ctx context.Context, id primitive.ObjectID) error {
	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": id})
	if err != nil {
		return fmt.Errorf("failed to delete person: %w", err)
	}

	if result.DeletedCount == 0 {
		return errors.NewNotFoundError("person not found")
	}

	return nil
}

// List lists persons with pagination
func (r *MongoPersonRepository) List(ctx context.Context, limit, offset int) ([]*entities.Person, int64, error) {
	// Count total documents
	total, err := r.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		return nil, 0, fmt.Errorf("failed to count persons: %w", err)
	}

	// Find with pagination
	opts := options.Find().
		SetSort(bson.D{{Key: "created_at", Value: -1}}).
		SetLimit(int64(limit)).
		SetSkip(int64(offset))

	cursor, err := r.collection.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to find persons: %w", err)
	}
	defer cursor.Close(ctx)

	var persons []*entities.Person
	if err := cursor.All(ctx, &persons); err != nil {
		return nil, 0, fmt.Errorf("failed to decode persons: %w", err)
	}

	return persons, total, nil
}
