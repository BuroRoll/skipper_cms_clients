package outputForms

import "Skipper_cms_clients/pkg/models"

type ClientsResponse struct {
	Clients            []models.User `json:"clients"`
	ClientsCount       int           `json:"clients_count"`
	ClientsCountSearch int           `json:"clients_count_search"`
}
