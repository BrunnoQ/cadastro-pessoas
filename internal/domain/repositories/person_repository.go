package repositories

import (
	"context"

	"github.com/BrunnoQ/cadastro-pessoas/internal/domain/entities"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// PersonRepository defines the interface for person persistence
type PersonRepository interface {
	// Create creates a new person
	Create(ctx context.Context, person *entities.Person) error

	// FindByID finds a person by ID
	FindByID(ctx context.Context, id primitive.ObjectID) (*entities.Person, error)

	// Update updates an existing person
	Update(ctx context.Context, person *entities.Person) error

	// Delete deletes a person by ID
	Delete(ctx context.Context, id primitive.ObjectID) error

	// List lists persons with pagination
	List(ctx context.Context, limit, offset int) ([]*entities.Person, int64, error)
}
