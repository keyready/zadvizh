package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"server/internal/domain/types/dto"
)

type Teacher struct {
	ID         bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Firstname  string        `bson:"firstname" json:"firstname"`
	Middlename string        `bson:"middlename" json:"middlename"`
	Lastname   string        `bson:"lastname" json:"lastname"`

	Likes    dto.Like      `bson:"likes" json:"likes"`
	Dislikes dto.Dislike   `bson:"dislikes" json:"dislikes"`
	Comments []dto.Comment `bson:"comments,omitempty" json:"comments"`
}
