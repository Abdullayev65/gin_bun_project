package io

type Sign struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type PostInput struct {
	Description string `json:"description"`
}
type CommentInput struct {
	PostID int64  `json:"post_id"`
	Text   string `json:"text"`
}
