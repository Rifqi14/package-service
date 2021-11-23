package response

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

type IFactoryResponses interface {
	Create(ctx *fiber.Ctx) error
}

// FactoryResponseWithMeta response with meta
type FactoryResponseWithMeta struct {
	data         interface{}
	meta         interface{}
	errorMessage error
}

func NewResponseWithMeta(bodyData interface{}, meta interface{}, errorMessage error) IFactoryResponses {
	return &FactoryResponseWithMeta{
		data:         bodyData,
		meta:         meta,
		errorMessage: errorMessage,
	}
}

func (resp FactoryResponseWithMeta) Create(ctx *fiber.Ctx) error {
	if resp.errorMessage != nil {
		return NewResponseUnprocessableEntity(resp.errorMessage).Create(ctx)
	}
	data := newBodySuccessWithMeta(resp.data, resp.meta)

	return ctx.Status(http.StatusOK).JSON(data.Create())
}

// FactoryResponseWithOutMeta response without meta
type FactoryResponseWithOutMeta struct {
	data         interface{}
	errorMessage error
}

func NewResponseWithOutMeta(bodyData interface{}, errorMessage error) IFactoryResponses {
	return &FactoryResponseWithOutMeta{
		data:         bodyData,
		errorMessage: errorMessage,
	}
}

func (resp FactoryResponseWithOutMeta) Create(ctx *fiber.Ctx) error {
	if resp.errorMessage != nil {
		return NewResponseUnprocessableEntity(resp.errorMessage).Create(ctx)
	}

	data := newBodyWithOutMeta(resp.data).Create()

	return ctx.Status(http.StatusOK).JSON(data)
}

// FactoryResponseBadRequest response bad request
type FactoryResponseBadRequest struct {
	errorMessage error
}

func NewResponseBadRequest(errorMessage error) IFactoryResponses {
	return FactoryResponseBadRequest{errorMessage: errorMessage}
}

func (resp FactoryResponseBadRequest) Create(ctx *fiber.Ctx) error {
	errorMessageVm := newBasicError(resp.errorMessage)

	return ctx.Status(http.StatusBadRequest).JSON(errorMessageVm)
}

// FactoryResponseUnprocessableEntity response unprocessable entity
type FactoryResponseUnprocessableEntity struct {
	errorMessage error
}

func NewResponseUnprocessableEntity(errorMessage error) IFactoryResponses {
	return &FactoryResponseUnprocessableEntity{errorMessage: errorMessage}
}

func (resp FactoryResponseUnprocessableEntity) Create(ctx *fiber.Ctx) error {
	errorMessageVm := newBasicError(resp.errorMessage).CreateError()

	return ctx.Status(http.StatusUnprocessableEntity).JSON(errorMessageVm)
}

//response error validator

type FactoryResponseErrorValidator struct {
	errorMessage validator.ValidationErrors
	translator   ut.Translator
}

func NewResponseErrorValidator(errorMessage validator.ValidationErrors, translator ut.Translator) IFactoryResponses {
	return &FactoryResponseErrorValidator{errorMessage: errorMessage, translator: translator}
}

func (resp FactoryResponseErrorValidator) Create(ctx *fiber.Ctx) error {
	errorMessageVm := newErrorFromValidator(resp.errorMessage, resp.translator)

	return ctx.Status(http.StatusBadRequest).JSON(errorMessageVm)
}

//response unauthorized

type FactoryResponseUnauthorized struct {
	errorMessage error
}

func NewResponseUnauthorized(errorMessage error) IFactoryResponses {
	return FactoryResponseUnauthorized{errorMessage: errorMessage}
}

func (resp FactoryResponseUnauthorized) Create(ctx *fiber.Ctx) error {
	errorMessageVm := newBasicError(resp.errorMessage).CreateError()

	return ctx.Status(http.StatusUnauthorized).JSON(errorMessageVm)
}
