package dal

import "github.com/jackc/pgx/v5/pgxpool"

type UsersDalInterface interface {
	InsertUser(username, pasHash string) error
	SelectUser(username string) (string, error)
	CheckUser(id string) error
	TruncateAllUsers() error
}

type userDal struct {
	DB *pgxpool.Pool
}

func NewUserDal(DB *pgxpool.Pool) *userDal {
	return &userDal{DB: DB}
}

func (d *userDal) InsertUser(username, hashedPassword string) error {
	query := `INSERT INTO users (username, hashed_password)
            VALUES ($1, $2);`
	_, err := d.DB.Exec(ctx, query, username, hashedPassword)
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

	err := d.DB.QueryRow(ctx, query, username).Scan(&jsonString)
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

	err := d.DB.QueryRow(ctx, query, id).Scan(&a)
	if err != nil {
		return err
	}

	return nil
}

func (d *userDal) TruncateAllUsers() error {
	query := `TRUNCATE TABLE users RESTART IDENTITY;`

	_, err := d.DB.Exec(ctx, query)
	if err != nil {
		return err
	}

	return nil
}
