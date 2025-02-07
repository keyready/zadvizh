package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Employee struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Firstname  string             `bson:"firstname" json:"firstname"`
	Lastname   string             `bson:"lastname" json:"lastname"`
	Department string             `bson:"department" json:"department"`
	Field      string             `bson:"field" json:"field"`
	TeamName   string             `bson:"teamName" json:"teamName"`
	TeamRole   string             `bson:"teamRole" json:"teamRole"`
	Scidir     string             `bson:"scidir" json:"scidir"`

	Ref string `bson:"ref" json:"ref"`
}
