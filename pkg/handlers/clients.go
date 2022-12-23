package handlers

import (
	"Skipper_cms_clients/pkg/models/forms/outputForms"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Description 	Получение всех пользователей
// @Tags 			Clients
// @Accept 			json
// @Produce 		json
// @Success 		200 		{object} 	outputForms.AuthResponse
// @Failure     	400         {object}  	outputForms.ErrorResponse
// @Failure     	500         {object}  	outputForms.ErrorResponse
// @Router /clients/ [get]
func (h *Handler) GetAllClients(c *gin.Context) {
	clients, err := h.services.GetClients()
	if err != nil {
		c.JSON(http.StatusNotFound, outputForms.ErrorResponse{Error: "Ошибка получения всех пользователей"})
	}
	c.JSON(http.StatusOK, clients)
}
