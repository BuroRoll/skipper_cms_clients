package main

import (
	"Skipper_cms_clients/pkg/handlers"
	"Skipper_cms_clients/pkg/models"
	"Skipper_cms_clients/pkg/repository"
	"Skipper_cms_clients/pkg/servises"
)

// @title Clients service
// @version 1.0
// @description Clients methods for skipper cms
func main() {
	db := models.GetDB()
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlerses := handlers.NewHandler(services)
	handlerses.InitRoutes()
}
