package request

type AuthEmployee struct {
	Firstname  string `json:"firstname"`
	Lastname   string `json:"lastname"`
	Department string `json:"department"`
	Field      string `json:"field"`
	TeamName   string `json:"teamName"`
	TeamRole   string `json:"teamRole"`
	Scidir     string `json:"scidir"`
}
