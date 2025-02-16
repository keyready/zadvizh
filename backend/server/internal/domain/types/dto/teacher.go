package dto

import (
	"go.mongodb.org/mongo-driver/v2/bson"
	"time"
)

type Comment struct {
	Author    bson.ObjectID `bson:"author" json:"author"`
	Content   string        `bson:"content" json:"content"`
	CreatedAt time.Time     `bson:"createdAt" json:"createdAt"`
}

type Like struct {
	Value   int64           `bson:"value" json:"value"`
	Authors []bson.ObjectID `bson:"authors" json:"authors"`
}

type Dislike struct {
	Value   int64           `bson:"value" json:"value"`
	Authors []bson.ObjectID `bson:"authors" json:"authors"`
}
