package request

type WriteNewComment struct {
	AuthorID  string `json:"authorId"`
	TeacherID string `json:"teacherId"`
	Content   string `json:"content"`
}

type LikeDislikeComment struct {
	Action    string `json:"action"`
	AuthorID  string `json:"authorId"`
	CommentID string `json:"commentId"`
}

type LikeDislike struct {
	Action    string `json:"action"`
	AuthorID  string `json:"authorId"`
	TeacherID string `json:"teacherId"`
}
