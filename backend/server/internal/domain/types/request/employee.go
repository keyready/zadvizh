package request

type AuthEmployee struct {
	Firstname  string `json:"firstname" bson:"firstname"`
	Lastname   string `json:"lastname" bson:"lastname"`
	Department string `json:"department" bson:"department"`
	Field      string `json:"field" bson:"field"`
	Position   string `json:"position" bson:"position"`
	TeamName   string `json:"teamName" bson:"teamname"`
	TeamRole   string `json:"teamRole" bson:"teamrole"`
	Scidir     string `json:"scidir" bson:"scidir"`
	TgId       string `bson:"tgid" json:"tgId"`
	Ref        string `bson:"ref" json:"ref"`
}
