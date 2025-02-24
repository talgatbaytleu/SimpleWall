package models

type User struct {
	ID          int    `json:"user_id"`
	Username    string `json:"username"`
	PaswordHash string `json:"hashed_pswd"`
	Password    string `json:"password"`
}
