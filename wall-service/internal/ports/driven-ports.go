package ports

type PsqlPort interface {
	SelectWall(user_id int) (string, error)
}

type RedisPort interface {
	GetWall(user_id string) (string, error)
	SetWall(user_id string, jsonWall string) error
}
