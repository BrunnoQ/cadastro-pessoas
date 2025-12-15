package usecases

import (
	"context"
	"fmt"

	"github.com/BrunnoQ/cadastro-pessoas/internal/domain/repositories"
	"github.com/BrunnoQ/cadastro-pessoas/internal/infrastructure/logger"
	"github.com/BrunnoQ/cadastro-pessoas/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

// DeletePersonUseCase handles deleting a person
type DeletePersonUseCase struct {
	personRepo repositories.PersonRepository
	logger     *logger.Logger
}

// NewDeletePersonUseCase creates a new DeletePersonUseCase
func NewDeletePersonUseCase(personRepo repositories.PersonRepository, logger *logger.Logger) *DeletePersonUseCase {
	return &DeletePersonUseCase{
		personRepo: personRepo,
		logger:     logger,
	}
}

// Execute deletes a person by ID
func (uc *DeletePersonUseCase) Execute(ctx context.Context, id string) error {
	// Validate and parse ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		uc.logger.Warn("Invalid person ID format", zap.String("id", id), zap.Error(err))
		return errors.NewValidationError("invalid person ID format", map[string]interface{}{
			"id": id,
		})
	}

	// Check if person exists before deletion
	_, err = uc.personRepo.FindByID(ctx, objectID)
	if err != nil {
		if errors.IsNotFoundError(err) {
			uc.logger.Warn("Person not found for deletion", zap.String("id", id))
			return errors.NewNotFoundError(fmt.Sprintf("person with id %s not found", id))
		}
		uc.logger.Error("Failed to find person for deletion", zap.String("id", id), zap.Error(err))
		return errors.NewDatabaseError("failed to retrieve person", err)
	}

	// Delete person (cascade delete of addresses and contacts happens at entity level)
	if err := uc.personRepo.Delete(ctx, objectID); err != nil {
		if errors.IsNotFoundError(err) {
			uc.logger.Warn("Person not found for deletion", zap.String("id", id))
			return errors.NewNotFoundError(fmt.Sprintf("person with id %s not found", id))
		}
		uc.logger.Error("Failed to delete person from database", zap.Error(err), zap.String("id", id))
		return errors.NewDatabaseError("failed to delete person", err)
	}

	uc.logger.Info("Person deleted successfully", zap.String("person_id", id))

	return nil
}
