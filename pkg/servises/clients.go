package service

import (
	"Skipper_cms_clients/pkg/models"
	"Skipper_cms_clients/pkg/repository"
)

type ClientsService struct {
	repo repository.Clients
}

func NewClientsService(repo repository.Clients) *ClientsService {
	return &ClientsService{repo: repo}
}

func (c ClientsService) GetClients(pagination *models.Pagination) ([]models.User, error) {
	return c.repo.GetClients(&pagination)
}

func (c ClientsService) GetClient(userId uint) (models.User, error) {
	return c.repo.GetClient(userId)
}
