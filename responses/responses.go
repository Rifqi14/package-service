package responses

import "github.com/gofiber/fiber/v2"

type Responses struct {
	ResponseFactory IFactoryResponse
}

func NewResponse(responseFactory IFactoryResponse) Responses {
	return Responses{
		ResponseFactory: responseFactory,
	}
}

func (response Responses) Send(ctx *fiber.Ctx) error {
	return response.ResponseFactory.Create(ctx)
}
