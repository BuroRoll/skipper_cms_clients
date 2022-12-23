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

func (c ClientsService) GetClients() ([]models.User, error) {
	return c.repo.GetClients()
}
