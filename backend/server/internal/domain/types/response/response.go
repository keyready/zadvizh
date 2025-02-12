package response

import "server/internal/domain/types/models"

type Data struct {
	Label     string          `json:"label"`
	DataLabel string          `json:"data-label,omitempty"`
	Employee  models.Employee `json:"employee"`
}

type Node struct {
	ID       string `json:"id"`
	Data     Data   `json:"data"`
	Children []Node `json:"children,omitempty"`
}
