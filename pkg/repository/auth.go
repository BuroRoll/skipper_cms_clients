package repository

import (
	"Skipper_cms_clients/pkg/models"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ClientsPostgres struct {
	db *gorm.DB
}

func NewClientsPostgres(db *gorm.DB) *ClientsPostgres {
	return &ClientsPostgres{db: db}
}

func (c ClientsPostgres) GetClients() ([]models.User, error) {
	var clients []models.User
	c.db.Preload(clause.Associations).Find(&clients)
	return clients, nil
}
