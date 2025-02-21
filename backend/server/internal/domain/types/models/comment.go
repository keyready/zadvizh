package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"server/internal/domain/types/dto"
)

type Comment struct {
	ID       bson.ObjectID `bson:"_id,omitempty" json:"id"`
	Content  string        `bson:"content" json:"content"`
	Author   bson.ObjectID `bson:"author" json:"author"`
	Likes    dto.Like      `bson:"likes" json:"likes"`
	Dislikes dto.Dislike   `bson:"dislikes" json:"dislikes"`
}
