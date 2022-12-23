package handlers

import (
	"Skipper_cms_clients/pkg/models/forms/outputForms"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetAllClients godoc
// @Description Получение списка менторов и менти
// @Security BearerAuth
// @Tags Clients
// @Accept json
// @Produce json
// @Success 	200 		{object} 	[]models.User
// @Router /clients/ [get]
func (h *Handler) GetAllClients(c *gin.Context) {
	clients, err := h.services.GetClients()
	if err != nil {
		c.JSON(http.StatusNotFound, outputForms.ErrorResponse{Error: "Ошибка получения всех пользователей"})
	}
	c.JSON(http.StatusOK, clients)
}
