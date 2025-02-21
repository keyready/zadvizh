package dto

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Like struct {
	Value   int64           `bson:"value" json:"value"`
	Authors []bson.ObjectID `bson:"authors" json:"authors"`
}

type Dislike struct {
	Value   int64           `bson:"value" json:"value"`
	Authors []bson.ObjectID `bson:"authors" json:"authors"`
}
