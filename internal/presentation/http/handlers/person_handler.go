package handlers

import (
	"net/http"

	"github.com/BrunnoQ/cadastro-pessoas/internal/application/dto"
	"github.com/BrunnoQ/cadastro-pessoas/internal/application/usecases"
	"github.com/BrunnoQ/cadastro-pessoas/internal/infrastructure/logger"
	"github.com/BrunnoQ/cadastro-pessoas/pkg/errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// ErrorResponse represents a standardized error response
type ErrorResponse struct {
	Code    string                 `json:"code"`
	Message string                 `json:"message"`
	Details map[string]interface{} `json:"details,omitempty"`
}

// ToErrorResponse converts an AppError to HTTP error response
func ToErrorResponse(err error) (int, ErrorResponse) {
	appErr, ok := err.(*errors.AppError)
	if !ok {
		return http.StatusInternalServerError, ErrorResponse{
			Code:    string(errors.CodeInternal),
			Message: "An unexpected error occurred",
		}
	}

	statusCode := getHTTPStatus(appErr.Code)
	return statusCode, ErrorResponse{
		Code:    string(appErr.Code),
		Message: appErr.Message,
		Details: appErr.Details,
	}
}

// getHTTPStatus maps error codes to HTTP status codes
func getHTTPStatus(code errors.ErrorCode) int {
	statusMap := map[errors.ErrorCode]int{
		errors.CodeValidation:      http.StatusBadRequest,
		errors.CodeInvalidInput:    http.StatusBadRequest,
		errors.CodeInvalidEmail:    http.StatusBadRequest,
		errors.CodeInvalidPhone:    http.StatusBadRequest,
		errors.CodeInvalidDate:     http.StatusBadRequest,
		errors.CodeInvalidName:     http.StatusBadRequest,
		errors.CodeInvalidSex:      http.StatusBadRequest,
		errors.CodeInvalidAge:      http.StatusBadRequest,
		errors.CodeBadRequest:      http.StatusBadRequest,
		errors.CodeNotFound:        http.StatusNotFound,
		errors.CodePersonNotFound:  http.StatusNotFound,
		errors.CodeConflict:        http.StatusConflict,
		errors.CodeDuplicate:       http.StatusConflict,
		errors.CodeConcurrency:     http.StatusConflict,
		errors.CodeUnauthorized:    http.StatusUnauthorized,
		errors.CodeForbidden:       http.StatusForbidden,
		errors.CodeTooManyRequests: http.StatusTooManyRequests,
		errors.CodeBusinessRule:    http.StatusUnprocessableEntity,
		errors.CodeInternal:        http.StatusInternalServerError,
		errors.CodeDatabase:        http.StatusInternalServerError,
		errors.CodeUnknown:         http.StatusInternalServerError,
	}

	if status, ok := statusMap[code]; ok {
		return status
	}
	return http.StatusInternalServerError
}

// PersonHandler handles person-related HTTP requests
type PersonHandler struct {
	createPersonUseCase *usecases.CreatePersonUseCase
	getPersonUseCase    *usecases.GetPersonUseCase
	listPersonsUseCase  *usecases.ListPersonsUseCase
	updatePersonUseCase *usecases.UpdatePersonUseCase
	logger              *logger.Logger
}

// NewPersonHandler creates a new PersonHandler
func NewPersonHandler(
	createPersonUseCase *usecases.CreatePersonUseCase,
	getPersonUseCase *usecases.GetPersonUseCase,
	listPersonsUseCase *usecases.ListPersonsUseCase,
	updatePersonUseCase *usecases.UpdatePersonUseCase,
	log *logger.Logger,
) *PersonHandler {
	return &PersonHandler{
		createPersonUseCase: createPersonUseCase,
		getPersonUseCase:    getPersonUseCase,
		listPersonsUseCase:  listPersonsUseCase,
		updatePersonUseCase: updatePersonUseCase,
		logger:              log,
	}
}

// Create handles POST /api/v1/persons
func (h *PersonHandler) Create(c *gin.Context) {
	h.logger.Info("Received create person request")

	var req dto.CreatePersonRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Warn("Failed to bind JSON", zap.Error(err))
		statusCode, errResponse := ToErrorResponse(err)
		c.JSON(statusCode, errResponse)
		return
	}

	h.logger.Info("JSON bound successfully", zap.Any("request", req))

	person, err := h.createPersonUseCase.Execute(c.Request.Context(), req)
	if err != nil {
		h.logger.Error("Failed to execute use case", zap.Error(err))
		statusCode, errResponse := ToErrorResponse(err)
		c.JSON(statusCode, errResponse)
		return
	}

	h.logger.Info("Person created successfully")
	c.JSON(http.StatusCreated, person)
}

// GetByID handles GET /api/v1/persons/:id
func (h *PersonHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	h.logger.Info("Received get person request", zap.String("id", id))

	person, err := h.getPersonUseCase.Execute(c.Request.Context(), id)
	if err != nil {
		h.logger.Error("Failed to get person", zap.String("id", id), zap.Error(err))
		statusCode, errResponse := ToErrorResponse(err)
		c.JSON(statusCode, errResponse)
		return
	}

	h.logger.Info("Person retrieved successfully", zap.String("id", id))
	c.JSON(http.StatusOK, person)
}

// List handles GET /api/v1/persons
func (h *PersonHandler) List(c *gin.Context) {
	var pagination dto.PaginationRequest

	// Bind query parameters
	if err := c.ShouldBindQuery(&pagination); err != nil {
		h.logger.Warn("Failed to bind pagination parameters", zap.Error(err))
		pagination = dto.DefaultPagination()
	} else {
		pagination.Validate()
	}

	h.logger.Info("Received list persons request",
		zap.Int("limit", pagination.Limit),
		zap.Int("offset", pagination.Offset),
	)

	result, err := h.listPersonsUseCase.Execute(c.Request.Context(), pagination)
	if err != nil {
		h.logger.Error("Failed to list persons", zap.Error(err))
		statusCode, errResponse := ToErrorResponse(err)
		c.JSON(statusCode, errResponse)
		return
	}

	h.logger.Info("Persons listed successfully", zap.Int("count", len(result.Data)))
	c.JSON(http.StatusOK, result)
}

// Update handles PUT /api/v1/persons/:id
func (h *PersonHandler) Update(c *gin.Context) {
	id := c.Param("id")
	h.logger.Info("Received update person request", zap.String("id", id))

	var req dto.UpdatePersonRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.Warn("Failed to bind JSON", zap.Error(err))
		statusCode, errResponse := ToErrorResponse(err)
		c.JSON(statusCode, errResponse)
		return
	}

	person, err := h.updatePersonUseCase.Execute(c.Request.Context(), id, req)
	if err != nil {
		h.logger.Error("Failed to update person", zap.String("id", id), zap.Error(err))
		statusCode, errResponse := ToErrorResponse(err)
		c.JSON(statusCode, errResponse)
		return
	}

	h.logger.Info("Person updated successfully", zap.String("id", id))
	c.JSON(http.StatusOK, person)
}
