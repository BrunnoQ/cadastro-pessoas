package usecases

import (
	"context"
	"fmt"

	"github.com/BrunnoQ/cadastro-pessoas/internal/application/dto"
	"github.com/BrunnoQ/cadastro-pessoas/internal/domain/repositories"
	"github.com/BrunnoQ/cadastro-pessoas/internal/infrastructure/logger"
	"github.com/BrunnoQ/cadastro-pessoas/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

// GetPersonUseCase handles retrieving a person by ID
type GetPersonUseCase struct {
	personRepo repositories.PersonRepository
	logger     *logger.Logger
}

// NewGetPersonUseCase creates a new GetPersonUseCase
func NewGetPersonUseCase(personRepo repositories.PersonRepository, logger *logger.Logger) *GetPersonUseCase {
	return &GetPersonUseCase{
		personRepo: personRepo,
		logger:     logger,
	}
}

// Execute retrieves a person by ID
func (uc *GetPersonUseCase) Execute(ctx context.Context, id string) (*dto.PersonResponse, error) {
	// Validate and parse ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		uc.logger.Warn("Invalid person ID format", zap.String("id", id), zap.Error(err))
		return nil, errors.NewValidationError("invalid person ID format", map[string]interface{}{
			"id": id,
		})
	}

	// Find person
	person, err := uc.personRepo.FindByID(ctx, objectID)
	if err != nil {
		if errors.IsNotFoundError(err) {
			uc.logger.Warn("Person not found", zap.String("id", id))
			return nil, errors.NewNotFoundError(fmt.Sprintf("person with id %s not found", id))
		}
		uc.logger.Error("Failed to find person", zap.String("id", id), zap.Error(err))
		return nil, errors.NewDatabaseError("failed to retrieve person", err)
	}

	uc.logger.Info("Person retrieved successfully",
		zap.String("person_id", person.ID.Hex()),
		zap.String("name", person.Name),
	)

	return dto.FromPerson(person), nil
}
