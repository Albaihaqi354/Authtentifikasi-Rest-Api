package main

import (
	"github.com/gofiber/fiber/v2"
	"shellrean.id/golang_auth/internal/api"
	"shellrean.id/golang_auth/internal/component"
	"shellrean.id/golang_auth/internal/config"
	"shellrean.id/golang_auth/internal/middleware"
	"shellrean.id/golang_auth/internal/repository"
	"shellrean.id/golang_auth/internal/service"
)

func main() {
	cnf := config.Get()
	dbConnection := component.GetDatabaseConnection(cnf)
	cacheConnection := component.GetCacheConnection()

	userRepository := repository.NewUser(dbConnection)

	userService := service.NewUser(userRepository, cacheConnection)

	authMid := middleware.Authenticate(userService)

	app := fiber.New()
	api.NewAuth(app, userService, authMid)

	_ = app.Listen(cnf.Server.Host + ":" + cnf.Server.Port)

}
