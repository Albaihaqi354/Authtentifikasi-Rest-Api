package api

import (
	"github.com/gofiber/fiber"
	"shellrean.id/golang_auth/domain"
	"shellrean.id/golang_auth/dto"
)

type authApi struct {
	userService domain.UserService
}

func NewAuth(app *fiber.App, userService domain.UserService) {
	h := authApi{
		userService: userService,
	}

	app.Post("token/generate", h.GenerateToken)
}

func (a authApi) GenerateToken(ctx *fiber.Ctx) error {
	var req dto.AuthReq
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.SendStatus(400)
	}

	user, err := a.userService.Authenticate(ctx.Context(), req)
	if err != nil {

	}
}
