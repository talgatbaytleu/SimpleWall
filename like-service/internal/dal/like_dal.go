package dal

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"liker/pkg/apperrors"
)

type LikeDalInterface interface {
	InsertLike(post_id int, user_id int) error
	SelectLikeCount(post_id int) (*string, error)
	DeleteLike(post_id int, user_id int) error
}

type likeDal struct {
	DB *pgxpool.Pool
}

func NewLikeDal(db *pgxpool.Pool) *likeDal {
	return &likeDal{DB: db}
}

func (d *likeDal) InsertLike(post_id int, user_id int) error {
	query := `INSERT INTO likes(post_id,user_id)
            VALUES ($1,$2);`

	_, err := d.DB.Exec(ctx, query, post_id, user_id)
	if err != nil {
		return err
	}

	return nil
}

func (d *likeDal) SelectLikeCount(post_id int) (*string, error) {
	var jsonPost string

	query := `SELECT to_jsonb(
              json_build_object(
                'post_id', post_id,
                'like_count', COUNT(user_id)
              )
            ) AS result
            FROM likes
            WHERE post_id = $1
            GROUP BY post_id;
`

	err := d.DB.QueryRow(ctx, query, post_id).Scan(&jsonPost)
	if err != nil {
		return &jsonPost, err
	}

	return &jsonPost, nil
}

func (d *likeDal) DeleteLike(post_id int, user_id int) error {
	query := `DELETE FROM likes
            WHERE post_id = $1 AND user_id = $2`

	cmdTag, err := d.DB.Exec(ctx, query, post_id, user_id)
	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return apperrors.ErrLikeNotDeleted
	}

	return nil
}
