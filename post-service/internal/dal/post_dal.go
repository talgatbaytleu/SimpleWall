package dal

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostDalInterface interface {
	InsertPost(user_id int, description string, imageLink string) error
	SelectPost(post_id int) (*string, error)
	UpdateTable(user_id int, post_id int, description string, imageLink string) error
	DeletePost(user_id int, post_id int) error
}

type postDal struct {
	DB *pgxpool.Pool
}

func NewPostDal(db *pgxpool.Pool) *postDal {
	return &postDal{DB: db}
}

func (d *postDal) InsertPost(user_id int, description string, imageLink string) error {
	query := `INSERT INTO posts(user_id,description,image_link)
            VALUES $1,$2,$3;`

	_, err := d.DB.Exec(ctx, query, user_id, description, imageLink)
	if err != nil {
		return err
	}

	return nil
}

func (d *postDal) SelectPost(post_id int) (*string, error) {
	var jsonPost string

	query := `SELECT to_jsonb(posts) 
            FROM posts
            WHERE post_id = $1;`

	err := d.DB.QueryRow(ctx, query, post_id).Scan(&jsonPost)
	if err != nil {
		return &jsonPost, err
	}

	return &jsonPost, nil
}

func (d *postDal) UpdateTable(
	user_id int,
	post_id int,
	description string,
	imageLink string,
) error {
	query := `UPDATE posts
            SET description = $1, imageLink = $2
            WHERE post_id = $3 AND user_id = $4`

	_, err := d.DB.Exec(ctx, query, description, imageLink, post_id, user_id)
	if err != nil {
		return err
	}

	return nil
}

func (d *postDal) DeletePost(user_id int, post_id int) error {
	query := `DELETE FROM posts
            WHERE user_id = $1 AND post_id = $2`

	cmdTag, err := d.DB.Exec(ctx, query, user_id, post_id)
	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return err // ERROr must be created!!!
	}

	return nil
}
