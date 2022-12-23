package service

import (
	"Skipper_cms_clients/pkg/models"
	"Skipper_cms_clients/pkg/repository"
)

type Clients interface {
	GetClients() ([]models.User, error)
}

type Service struct {
	Clients
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Clients: NewClientsService(repos.Clients),
	}
}
