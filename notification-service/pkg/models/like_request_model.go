package models

type LikeType struct {
	PostID int `json:"post_id"`
	UserID int `json:"user_id"`
}
