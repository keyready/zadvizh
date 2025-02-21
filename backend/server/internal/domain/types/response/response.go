package response

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"server/internal/domain/types/models"
)

type Data struct {
	Label     string `json:"label"`
	DataLabel string `json:"data-label,omitempty"`
	models.Employee
}

type Node struct {
	ID       string `json:"id"`
	Data     Data   `json:"data"`
	Children []Node `json:"children,omitempty"`
}

type Token struct {
	ID          bson.ObjectID `json:"id"`
	AccessToken string        `json:"accessToken"`
}
