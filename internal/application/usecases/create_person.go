package usecases

import (
	"context"
	"fmt"

	"github.com/BrunnoQ/cadastro-pessoas/internal/application/dto"
	"github.com/BrunnoQ/cadastro-pessoas/internal/domain/entities"
	"github.com/BrunnoQ/cadastro-pessoas/internal/domain/repositories"
	"github.com/BrunnoQ/cadastro-pessoas/internal/infrastructure/logger"
	"github.com/BrunnoQ/cadastro-pessoas/pkg/errors"
	"go.uber.org/zap"
)

// CreatePersonUseCase handles person creation
type CreatePersonUseCase struct {
	personRepo repositories.PersonRepository
	logger     *logger.Logger
}

// NewCreatePersonUseCase creates a new CreatePersonUseCase
func NewCreatePersonUseCase(personRepo repositories.PersonRepository, logger *logger.Logger) *CreatePersonUseCase {
	return &CreatePersonUseCase{
		personRepo: personRepo,
		logger:     logger,
	}
}

// Execute creates a new person
func (uc *CreatePersonUseCase) Execute(ctx context.Context, req dto.CreatePersonRequest) (*dto.PersonResponse, error) {
	// Create person entity
	person, err := entities.NewPerson(req.Name, req.Surname, req.Birthdate.Time, req.Sex)
	if err != nil {
		uc.logger.Warn("Failed to create person entity", zap.Error(err))
		return nil, errors.NewValidationError(err.Error(), nil)
	}

	// Add addresses
	for _, addrReq := range req.Addresses {
		addr, err := addrReq.ToAddress()
		if err != nil {
			uc.logger.Warn("Invalid address", zap.Error(err))
			return nil, errors.NewValidationError(fmt.Sprintf("invalid address: %v", err), nil)
		}
		if err := person.AddAddress(*addr); err != nil {
			uc.logger.Warn("Failed to add address", zap.Error(err))
			return nil, errors.NewValidationError(err.Error(), nil)
		}
	}

	// Add contacts
	for _, contactReq := range req.Contacts {
		contact, err := contactReq.ToContact()
		if err != nil {
			uc.logger.Warn("Invalid contact", zap.Error(err))
			return nil, errors.NewValidationError(fmt.Sprintf("invalid contact: %v", err), nil)
		}
		if err := person.AddContact(*contact); err != nil {
			uc.logger.Warn("Failed to add contact", zap.Error(err))
			return nil, errors.NewValidationError(err.Error(), nil)
		}
	}

	// Persist person
	if err := uc.personRepo.Create(ctx, person); err != nil {
		uc.logger.Error("Failed to create person in database", zap.Error(err), zap.String("name", person.Name))
		return nil, errors.NewDatabaseError("failed to create person", err)
	}

	uc.logger.Info("Person created successfully",
		zap.String("person_id", person.ID.Hex()),
		zap.String("name", person.Name),
		zap.String("surname", person.Surname),
	)

	// Return response
	return dto.FromPerson(person), nil
}
