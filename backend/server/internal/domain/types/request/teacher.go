package request

type WriteNewComment struct {
	AuthorID  string `json:"authorId"`
	TeacherID string `json:"teacherId"`
	Content   string `json:"content"`
}

type LikeDislike struct {
	Action    string `json:"action"`
	AuthorID  string `json:"authorId"`
	TeacherID string `json:"teacherId"`
}
