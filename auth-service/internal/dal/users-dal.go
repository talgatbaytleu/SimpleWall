package dal

type UsersDalInterface interface {
	InsertUser(username, pasHash string) error
	SelectUser(username string) (string, error)
}

type userDal struct{}

func NewUserDal() *userDal {
	return &userDal{}
}

func (d *userDal) InsertUser(username, pasHash string) error {
	query := `INSERT INTO users (username, pas_hash)
            VALUES ($1, $2);`
	_, err := DB.Exec(ctx, query, username, pasHash)
	if err != nil {
		return err
	}

	return nil
}

func (d *userDal) SelectUser(username string) (string, error) {
	var pas_hash string

	query := `SELECT pas_hash FROM users
            WHERE username = $1;`

	err := DB.QueryRow(ctx, query, username).Scan(pas_hash)
	if err != nil {
		return pas_hash, err
	}

	return pas_hash, nil
}
