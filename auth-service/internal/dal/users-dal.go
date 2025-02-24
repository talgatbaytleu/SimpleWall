package dal

type UsersDalInterface interface {
	InsertUser(username, pasHash string) error
	SelectUser(username string) (string, error)
}

type userDal struct{}

func NewUserDal() *userDal {
	return &userDal{}
}

func (d *userDal) InsertUser(username, hashedPassword string) error {
	query := `INSERT INTO users (username, hashed_pswd)
            VALUES ($1, $2);`
	_, err := DB.Exec(ctx, query, username, hashedPassword)
	if err != nil {
		return err
	}

	return nil
}

func (d *userDal) SelectUser(username string) (string, error) {
	var jsonString string

	query := `SELECT jsonb_agg(to_jsonb(users)) 
            FROM users
            WHERE username = $1;`

	err := DB.QueryRow(ctx, query, username).Scan(jsonString)
	if err != nil {
		return jsonString, err
	}

	return jsonString, nil
}
