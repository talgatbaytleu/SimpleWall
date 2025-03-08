package psqladapter

import (
	"github.com/jackc/pgx/v5/pgxpool"
)

type psqlAdapter struct {
	DB *pgxpool.Pool
}

func NewPsqlAdapter(db *pgxpool.Pool) *psqlAdapter {
	return &psqlAdapter{DB: db}
}

func (a *psqlAdapter) SelectWall(user_id int) (string, error) {
	var jsonWall string

	query := `
            SELECT jsonb_agg(jsonb_build_object(
              'post_id', p.post_id,
              'user_id', p.user_id,
              'description', p.description,
              'image_link', p.image_link,
              'likes_count', COALESCE(l.like_count, 0),
              'comments_count', COALESCE(c.comment_count, 0),
              'liked_by_me', CASE WHEN lme.user_id IS NOT NULL THEN true ELSE false END
              ))
            FROM posts p
            LEFT JOIN (
              SELECT post_id, COUNT(*) AS like_count
              FROM likes
              GROUP BY post_id
            ) l ON p.post_id = l.post_id
            LEFT JOIN (
              SELECT post_id, COUNT(*) AS comment_count
              FROM comments
              GROUP BY post_id
            ) c ON p.post_id = c.post_id
            LEFT JOIN likes lme ON p.post_id = lme.post_id AND lme.user_id = $1  -- Check if the user has liked the post
            WHERE p.user_id = $1;
    `

	err := a.DB.QueryRow(ctx, query, user_id).Scan(&jsonWall)
	if err != nil {
		return "", err
	}

	return jsonWall, nil
}
