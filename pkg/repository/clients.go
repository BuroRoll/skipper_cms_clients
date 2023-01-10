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

func (c ClientsPostgres) GetClients(pagination **models.Pagination) ([]models.User, error) {
	var clients []models.User
	offset := ((*pagination).Page - 1) * (*pagination).Limit
	queryBuider := c.db.Limit((*pagination).Limit).Offset(offset)

	var result *gorm.DB
	if (len((*pagination).Search)) > 0 {
		result = queryBuider.
			Preload(clause.Associations).
			Where("phone IN (?) OR first_name IN (?) OR second_name IN (?) OR patronymic IN (?)", (*pagination).Search, (*pagination).Search, (*pagination).Search, (*pagination).Search).
			Find(&clients)
	} else {
		result = queryBuider.
			Preload(clause.Associations).
			Find(&clients)
	}
	//c.db.Preload(clause.Associations).Find(&clients)
	if result.Error != nil {
		return nil, result.Error
	}

	return clients, nil
}

func (c ClientsPostgres) GetClient(userId uint) (models.User, error) {
	var client models.User
	c.db.Preload(clause.Associations).Find(&client, userId)
	return client, nil
}
