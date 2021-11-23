package response

import (
	"github.com/gofiber/fiber/v2"
)

//vm

type SuccessResponseWithOutMetaVm struct {
	Data interface{} `json:"data"`
}

func newSuccessResponseWithOutMeta(data interface{}) SuccessResponseWithOutMetaVm {
	return SuccessResponseWithOutMetaVm{Data: data}
}

type SuccessResponseWithMetaVm struct {
	Data interface{} `json:"data"`
	Meta interface{} `json:"meta"`
}

func newSuccessResponseWithMeta(data, meta interface{}) SuccessResponseWithMetaVm {
	return SuccessResponseWithMetaVm{
		Data: data,
		Meta: meta,
	}
}

type ErrorResponseVm struct {
	Message interface{}
}

func newErrorResponse(message interface{}) ErrorResponseVm {
	return ErrorResponseVm{Message: message}
}

type Response struct {
	ResponseFactory IFactoryResponses
}

func NewResponse(responseFactory IFactoryResponses) Response {
	return Response{
		ResponseFactory: responseFactory,
	}
}

func (response Response) Send(ctx *fiber.Ctx) error {
	return response.ResponseFactory.Create(ctx)
}
