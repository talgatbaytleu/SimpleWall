package models

type PostType struct {
	PostID      int    `json:"post_id"`
	UserID      int    `json:"user_id"`
	Description string `json:"description"`
	ImageLink   string `json:"image_link"`
}
