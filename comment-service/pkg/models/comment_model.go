package models

type CommentType struct {
	CommentID int    `json:"comment_id"`
	UserID    int    `json:"user_id"`
	PostID    int    `json:"post_id"`
	Content   string `json:"content"`
}
