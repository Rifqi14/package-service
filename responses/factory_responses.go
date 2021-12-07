package responses

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gitlab.com/s2.1-backend/shm-package-svc/str"
)

type IFactoryResponse interface {
	Create(ctx *fiber.Ctx) error
}

type FactoryBaseResponse struct {
	Data   interface{}           `json:"data"`
	Meta   interface{}           `json:"meta"`
	Status FactoryStatusResponse `json:"status"`
}

type FactoryStatusResponse struct {
	Success      bool        `json:"success"`
	Code         int         `json:"code"`
	Message      string      `json:"message"`
	ErrorMessage interface{} `json:"error_message"`
}

func ResponseSuccess(data, meta interface{}, message string) IFactoryResponse {
	return &FactoryBaseResponse{
		Data: data,
		Meta: meta,
		Status: FactoryStatusResponse{
			Success:      true,
			Code:         http.StatusOK,
			Message:      message,
			ErrorMessage: nil,
		},
	}
}

func ResponseError(data, meta interface{}, code int, message string, errorMessage error) IFactoryResponse {
	return &FactoryBaseResponse{
		Data: data,
		Meta: meta,
		Status: FactoryStatusResponse{
			Success:      false,
			Code:         code,
			Message:      message,
			ErrorMessage: errorMessage.Error(),
		},
	}
}

func ResponseErrorValidation(data, meta interface{}, code int, message string, errorMessage validator.ValidationErrors) IFactoryResponse {
	return &FactoryBaseResponse{
		Data: data,
		Meta: meta,
		Status: FactoryStatusResponse{
			Success:      false,
			Code:         code,
			Message:      message,
			ErrorMessage: buildErrorValidation(errorMessage),
		},
	}
}

func (resp FactoryBaseResponse) Create(ctx *fiber.Ctx) error {
	return ctx.Status(http.StatusOK).JSON(resp)
}

func buildErrorValidation(errorResponse validator.ValidationErrors) interface{} {
	errorMessage := map[string][]string{}

	for _, err := range errorResponse {
		errKey := str.Underscore(err.StructField())
		errorMessage[errKey] = append(
			errorMessage[errKey],
			buildErrorValidationMessage(err),
		)
	}
	return errorMessage
}

func buildErrorValidationMessage(errorRes validator.FieldError) string {
	var sb strings.Builder

	sb.WriteString("Validation failed on field '" + str.Underscore(errorRes.StructField()) + "'")
	sb.WriteString(", condition: " + errorRes.ActualTag())

	// Print conidition parameters, e.g. min=4 -> { 4 }
	if errorRes.Param() != "" {
		sb.WriteString(" { " + errorRes.Param() + " }")
	}

	if errorRes.Value() != nil && errorRes.Value() != "" {
		sb.WriteString(fmt.Sprintf(", actual: %v", errorRes.Value()))
	}

	return sb.String()
}
