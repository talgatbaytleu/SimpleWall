package dal

import (
	"github.com/jackc/pgx/v5/pgxpool"

	"commenter/pkg/apperrors"
)

type CommentDalInterface interface {
	InsertComment(post_id int, user_id int, content string) error
	UpdateComment(comment_id int, content string) error
	SelectComment(comment_id int) (string, error)
	SelectPostComments(post_id int) (string, error)
	DeleteComment(comment_id int, user_id int) error
	SelectUserOfComment(comment_id int) (int, error)
}

type commentDal struct {
	DB *pgxpool.Pool
}

func NewCommentDal(db *pgxpool.Pool) *commentDal {
	return &commentDal{DB: db}
}

func (d *commentDal) InsertComment(post_id int, user_id int, content string) error {
	query := `INSERT INTO comments(post_id,user_id,content)
            VALUES ($1,$2,$3);`

	cmdTag, err := d.DB.Exec(ctx, query, post_id, user_id, content)
	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return apperrors.ErrCommentNotCreated
	}

	return nil
}

func (d *commentDal) UpdateComment(comment_id int, content string) error {
	query := `UPDATE comments
            SET content = $2
            WHERE comment_id = $1;`

	cmdTag, err := d.DB.Exec(ctx, query, comment_id, content)
	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return apperrors.ErrCommentNotUpdated
	}

	return nil
}

func (d *commentDal) SelectComment(comment_id int) (string, error) {
	var jsonPost string

	query := `SELECT to_json(comments)
            FROM comments
            WHERE comment_id = $1`

	err := d.DB.QueryRow(ctx, query, comment_id).Scan(&jsonPost)
	if err != nil {
		return "", err
	}

	return jsonPost, nil
}

func (d commentDal) SelectPostComments(post_id int) (string, error) {
	var jsonPost string
	query := `SELECT COALESCE(jsonb_agg(comments), '[]'::jsonb)
            FROM comments
            WHERE post_id = $1;

  `

	err := d.DB.QueryRow(ctx, query, post_id).Scan(&jsonPost)
	if err != nil {
		return "", err
	}

	return jsonPost, nil
}

func (d *commentDal) DeleteComment(comment_id int, user_id int) error {
	query := `DELETE FROM comments
            WHERE comment_id = $1 AND user_id = $2`

	cmdTag, err := d.DB.Exec(ctx, query, comment_id, user_id)
	if err != nil {
		return err
	}

	if cmdTag.RowsAffected() == 0 {
		return apperrors.ErrCommentNotDeleted
	}

	return nil
}

func (d *commentDal) SelectUserOfComment(comment_id int) (int, error) {
	var user_id int
	query := `SELECT user_id FROM comments
            WHERE comment_id = $1;`

	err := d.DB.QueryRow(ctx, query, comment_id).Scan(&user_id)

	return user_id, err
}
