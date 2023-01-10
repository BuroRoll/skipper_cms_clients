package repository

import (
	"Skipper_cms_clients/pkg/models"
	"gorm.io/gorm"
)

type Clients interface {
	GetClients(**models.Pagination) ([]models.User, error)
	GetClient(userId uint) (models.User, error)
}

type Repository struct {
	Clients
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Clients: NewClientsPostgres(db),
	}
}
