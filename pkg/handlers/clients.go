package handlers

import (
	"Skipper_cms_clients/pkg/models/forms/inputForms"
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
// @Param        limit   	query      int  	false  "limit"
// @Param        page   	query      int  	false  "page"
// @Param        search   	query      string  	false  "search"
// @Success 	200 		{object} 	[]outputForms.ClientsResponse
// @Router /clients/ [get]
func (h *Handler) GetAllClients(c *gin.Context) {
	pagination := GeneratePaginationFromRequest(c)
	clients, err := h.services.GetClients(&pagination)
	clientsCount := h.services.GetClientsCount()
	if err != nil {
		c.JSON(http.StatusNotFound, outputForms.ErrorResponse{Error: "Ошибка получения всех пользователей"})
	}
	c.JSON(http.StatusOK, outputForms.ClientsResponse{
		Clients:            clients,
		ClientsCount:       clientsCount,
		ClientsCountSearch: len(clients)})
}

// GetClientData godoc
// @Description Получение конкретного пользователя
// @Security BearerAuth
// @Tags Clients
// @Accept json
// @Produce json
// @Param        ClientId   path      int  true  "Client ID"
// @Success 	200 		{object} 	models.User
// @Router /clients/{ClientId} [get]
func (h *Handler) GetClientData(c *gin.Context) {
	var clientId inputForms.ClientIdInput
	if err := c.ShouldBindUri(&clientId); err != nil {
		c.JSON(400, gin.H{"msg": err})
		return
	}
	client, err := h.services.GetClient(clientId.ClientId)
	if err != nil {
		c.JSON(http.StatusNotFound, outputForms.ErrorResponse{Error: "Ошибка получения пользователя"})
	}
	c.JSON(http.StatusOK, client)
}
