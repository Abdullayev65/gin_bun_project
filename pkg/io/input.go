package io

type Sign struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type PostInput struct {
	Description   string `json:"description"`
	AttachmentIDs []int  `json:"attachmentIDs"`
}
type CommentInput struct {
	PostID int    `json:"post_id"`
	Text   string `json:"text"`
}
