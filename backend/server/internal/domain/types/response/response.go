package response

import "server/internal/domain/types/models"

type Three struct {
	Org     []models.Employee            `json:"org"`
	Dev     map[string][]models.Employee `json:"dev"`
	Sec     map[string][]models.Employee `json:"sec"`
	Devops  []models.Employee            `json:"devops"`
	Science map[string][]models.Employee `json:"science"`
}
