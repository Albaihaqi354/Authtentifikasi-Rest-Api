package main

import (
	"honnef.co/go/tools/lintcmd/cache"
	"shellrean.id/golang_auth/internal/component"
	"shellrean.id/golang_auth/internal/config"
	"shellrean.id/golang_auth/internal/repository"
	"shellrean.id/golang_auth/internal/service"
)

func main() {
	cnf := config.Get()
	dbConnection := component.GetDatabaseConnection(cnf)
	cacheConnection := cache.GetCacheConnection()

	userRepository := repository.NewUser(dbConnection)

	userService := service.NewUser(userRepository, cacheConnection)

}
