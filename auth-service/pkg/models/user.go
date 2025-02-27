package models

type User struct {
	ID           int    `json:"user_id"`
	Username     string `json:"username"`
	PasswordHash string `json:"hashed_password"`
	Password     string `json:"password"`
	CreatedAt    string `json:"created_at"`
}
