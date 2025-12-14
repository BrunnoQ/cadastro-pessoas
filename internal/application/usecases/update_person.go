package usecases

import (
	"context"
	"fmt"

	"github.com/BrunnoQ/cadastro-pessoas/internal/application/dto"
	"github.com/BrunnoQ/cadastro-pessoas/internal/domain/entities"
	"github.com/BrunnoQ/cadastro-pessoas/internal/domain/repositories"
	"github.com/BrunnoQ/cadastro-pessoas/internal/infrastructure/logger"
	"github.com/BrunnoQ/cadastro-pessoas/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
)

// UpdatePersonUseCase handles updating a person
type UpdatePersonUseCase struct {
	personRepo repositories.PersonRepository
	logger     *logger.Logger
}

// NewUpdatePersonUseCase creates a new UpdatePersonUseCase
func NewUpdatePersonUseCase(personRepo repositories.PersonRepository, logger *logger.Logger) *UpdatePersonUseCase {
	return &UpdatePersonUseCase{
		personRepo: personRepo,
		logger:     logger,
	}
}

// Execute updates a person
func (uc *UpdatePersonUseCase) Execute(ctx context.Context, id string, req dto.UpdatePersonRequest) (*dto.PersonResponse, error) {
	// Validate and parse ObjectID
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		uc.logger.Warn("Invalid person ID format", zap.String("id", id), zap.Error(err))
		return nil, errors.NewValidationError("invalid person ID format", map[string]interface{}{
			"id": id,
		})
	}

	// Find existing person
	person, err := uc.personRepo.FindByID(ctx, objectID)
	if err != nil {
		if errors.IsNotFoundError(err) {
			uc.logger.Warn("Person not found for update", zap.String("id", id))
			return nil, errors.NewNotFoundError(fmt.Sprintf("person with id %s not found", id))
		}
		uc.logger.Error("Failed to find person for update", zap.String("id", id), zap.Error(err))
		return nil, errors.NewDatabaseError("failed to retrieve person", err)
	}

	// Update fields if provided
	if req.Name != nil {
		person.Name = *req.Name
	}
	if req.Surname != nil {
		person.Surname = *req.Surname
	}
	if req.Birthdate != nil {
		person.Birthdate = req.Birthdate.Time
	}
	if req.Sex != nil {
		person.Sex = *req.Sex
	}

	// Replace addresses if provided
	if req.Addresses != nil {
		person.Addresses = []entities.Address{}
		for _, addrReq := range *req.Addresses {
			addr, err := addrReq.ToAddress()
			if err != nil {
				uc.logger.Warn("Invalid address in update", zap.Error(err))
				return nil, errors.NewValidationError(fmt.Sprintf("invalid address: %v", err), nil)
			}
			if err := person.AddAddress(*addr); err != nil {
				uc.logger.Warn("Failed to add address", zap.Error(err))
				return nil, errors.NewValidationError(err.Error(), nil)
			}
		}
	}

	// Replace contacts if provided
	if req.Contacts != nil {
		person.Contacts = []entities.Contact{}
		for _, contactReq := range *req.Contacts {
			contact, err := contactReq.ToContact()
			if err != nil {
				uc.logger.Warn("Invalid contact in update", zap.Error(err))
				return nil, errors.NewValidationError(fmt.Sprintf("invalid contact: %v", err), nil)
			}
			if err := person.AddContact(*contact); err != nil {
				uc.logger.Warn("Failed to add contact", zap.Error(err))
				return nil, errors.NewValidationError(err.Error(), nil)
			}
		}
	}

	// Validate updated person
	if err := person.Validate(); err != nil {
		uc.logger.Warn("Updated person validation failed", zap.Error(err))
		return nil, errors.NewValidationError(err.Error(), nil)
	}

	// Update in repository
	if err := uc.personRepo.Update(ctx, person); err != nil {
		if errors.IsConcurrencyError(err) {
			uc.logger.Warn("Concurrent modification detected", zap.String("id", id))
			return nil, errors.NewConcurrencyError("person was modified by another process, please retry")
		}
		uc.logger.Error("Failed to update person in database", zap.Error(err), zap.String("id", id))
		return nil, errors.NewDatabaseError("failed to update person", err)
	}

	uc.logger.Info("Person updated successfully",
		zap.String("person_id", person.ID.Hex()),
		zap.String("name", person.Name),
	)

	return dto.FromPerson(person), nil
}
