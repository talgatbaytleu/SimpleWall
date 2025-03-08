package ports

type WallService interface {
	GetUserWall(user_idStr string) (string, error)
	// methods
}
