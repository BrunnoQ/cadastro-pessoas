package usecases

import (
	"context"

	"github.com/BrunnoQ/cadastro-pessoas/internal/application/dto"
	"github.com/BrunnoQ/cadastro-pessoas/internal/domain/repositories"
	"github.com/BrunnoQ/cadastro-pessoas/internal/infrastructure/logger"
	"github.com/BrunnoQ/cadastro-pessoas/pkg/errors"
	"go.uber.org/zap"
)

// ListPersonsUseCase handles retrieving a paginated list of persons
type ListPersonsUseCase struct {
	personRepo repositories.PersonRepository
	logger     *logger.Logger
}

// NewListPersonsUseCase creates a new ListPersonsUseCase
func NewListPersonsUseCase(personRepo repositories.PersonRepository, logger *logger.Logger) *ListPersonsUseCase {
	return &ListPersonsUseCase{
		personRepo: personRepo,
		logger:     logger,
	}
}

// Execute retrieves a paginated list of persons
func (uc *ListPersonsUseCase) Execute(ctx context.Context, pagination dto.PaginationRequest) (*dto.PersonListResponse, error) {
	// Validate pagination
	pagination.Validate()

	uc.logger.Info("Listing persons",
		zap.Int("limit", pagination.Limit),
		zap.Int("offset", pagination.Offset),
	)

	// Get paginated list from repository
	persons, total, err := uc.personRepo.List(ctx, pagination.Limit, pagination.Offset)
	if err != nil {
		uc.logger.Error("Failed to list persons", zap.Error(err))
		return nil, errors.NewDatabaseError("failed to retrieve persons list", err)
	}

	uc.logger.Info("Persons listed successfully",
		zap.Int("count", len(persons)),
		zap.Int64("total", total),
	)

	return dto.FromPersons(persons, total, pagination.Limit, pagination.Offset), nil
}
