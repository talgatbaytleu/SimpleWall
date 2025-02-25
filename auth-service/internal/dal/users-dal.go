package dal

type UsersDalInterface interface {
	InsertUser(username, pasHash string) error
	SelectUser(username string) (string, error)
	CheckUser(id string) error
}

type userDal struct{}

func NewUserDal() *userDal {
	return &userDal{}
}

func (d *userDal) InsertUser(username, hashedPassword string) error {
	query := `INSERT INTO users (username, hashed_password)
            VALUES ($1, $2);`
	_, err := DB.Exec(ctx, query, username, hashedPassword)
	if err != nil {
		return err
	}

	return nil
}

func (d *userDal) SelectUser(username string) (string, error) {
	var jsonString string

	query := `SELECT to_jsonb(users) 
            FROM users
            WHERE username = $1;`

	err := DB.QueryRow(ctx, query, username).Scan(&jsonString)
	if err != nil {
		return jsonString, err
	}

	return jsonString, nil
}

func (d *userDal) CheckUser(id string) error {
	var a int
	query := `SELECT user_id 
            FROM users    
            WHERE user_id = $1`

	err := DB.QueryRow(ctx, query, id).Scan(&a)
	if err != nil {
		return err
	}

	return nil
}
