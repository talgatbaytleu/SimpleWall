package dal

import (
	"errors"
	"fmt"

	"github.com/jackc/pgx"
	"github.com/jackc/pgx/v5/pgxpool"

	"liker/pkg/apperrors"
)

type LikeDalInterface interface {
	InsertLike(post_id int, user_id int) error
	SelectLikesCount(post_id int) (string, error)
	SelectLikesList(post_id int) (string, error)
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

func (d *likeDal) SelectLikesCount(post_id int) (string, error) {
	var jsonPost *string

	query := `SELECT to_jsonb(
              json_build_object(
                'post_id', l.post_id,
                'like_count', COUNT(l.user_id)
              )
            ) AS result
            FROM likes l
            WHERE l.post_id = $1
            GROUP BY l.post_id;
`

	err := d.DB.QueryRow(ctx, query, post_id).Scan(&jsonPost)
	if err != nil {
		if err == pgx.ErrNoRows || errors.Is(err, pgx.ErrNoRows) ||
			err.Error() == "no rows in result set" {
			// If there are no likes, return a valid empty JSON response
			return `{"post_id":` + fmt.Sprintf("%d", post_id) + `,"like_count":0}`, nil
		}
		return "", err
	}

	if jsonPost == nil {
		return `{"post_id":` + fmt.Sprintf("%d", post_id) + `,"like_count":0}`, nil
	}

	return *jsonPost, nil
}

func (d *likeDal) SelectLikesList(post_id int) (string, error) {
	var jsonPost *string

	query := `SELECT jsonb_agg(likes)
            FROM likes
            WHERE post_id = $1;`

	err := d.DB.QueryRow(ctx, query, post_id).Scan(&jsonPost)
	if err != nil {
		return *jsonPost, err
	}

	if jsonPost == nil {
		return "[]", nil
	}

	return *jsonPost, nil
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
