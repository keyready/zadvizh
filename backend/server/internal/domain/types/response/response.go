package response

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"server/internal/domain/types/dto"
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

type Teacher struct {
	ID         bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Firstname  string        `bson:"firstname" json:"firstname"`
	Middlename string        `bson:"middlename" json:"middlename"`
	Lastname   string        `bson:"lastname" json:"lastname"`

	Likes    dto.Like         `bson:"likes" json:"likes"`
	Dislikes dto.Dislike      `bson:"dislikes" json:"dislikes"`
	Comments []models.Comment `json:"comments"`
}
