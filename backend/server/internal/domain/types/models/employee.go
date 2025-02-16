package models

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Employee struct {
	ID         bson.ObjectID `bson:"_id" json:"id"`
	Firstname  string        `bson:"firstname" json:"firstname"`
	Lastname   string        `bson:"lastname" json:"lastname"`
	Department string        `bson:"department" json:"department"`
	Field      string        `bson:"field" json:"field"`
	Position   string        `bson:"position" json:"position"`
	TeamName   string        `bson:"teamname" json:"teamName"`
	TeamRole   string        `bson:"teamrole" json:"teamRole"`
	Scidir     string        `bson:"scidir" json:"scidir"`

	TgId         string `bson:"tgid" json:"tgId"`
	Ref          string `bson:"ref" json:"ref"`
	TgInviteLink string `bson:"tgInviteLink" json:"tgInviteLink"`
}
