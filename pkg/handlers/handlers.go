package handlers

import (
	docs "Skipper_cms_clients/pkg/docs"
	service "Skipper_cms_clients/pkg/servises"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() {
	router := gin.Default()
	router.Use(corsMiddleware())
	api_v1 := router.Group("/api/v1")
	{
		clients := api_v1.Group("/clients")
		{
			clients.GET("/", h.GetAllClients)
			//clients.GET("/:clientId", h.GetClientData)
			//reports := clients.Group("/reports")
			//{
			//	reports.GET("/:clientId", h.GetClientReports)
			//}
		}
	}

	docs.SwaggerInfo.BasePath = "/api/v1"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	router.Run()
}
