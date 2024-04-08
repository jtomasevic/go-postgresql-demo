// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: reviews.sql

package store

import (
	"context"
	"time"

	"github.com/google/uuid"
)

const allReviews = `-- name: AllReviews :many
SELECT ID, MovieID, Comment, Rating, CommentTime FROM Reviews
`

func (q *Queries) AllReviews(ctx context.Context) ([]Review, error) {
	rows, err := q.db.Query(ctx, allReviews)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Review
	for rows.Next() {
		var i Review
		if err := rows.Scan(
			&i.ID,
			&i.Movieid,
			&i.Comment,
			&i.Rating,
			&i.Commenttime,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const createReview = `-- name: CreateReview :exec
INSERT INTO Reviews (ID, MovieID, Comment, Rating, CommentTime) VALUES ($1, $2, $3, $4, $5)
`

type CreateReviewParams struct {
	ID          uuid.UUID
	Movieid     uuid.UUID
	Comment     string
	Rating      int
	Commenttime time.Time
}

func (q *Queries) CreateReview(ctx context.Context, arg CreateReviewParams) error {
	_, err := q.db.Exec(ctx, createReview,
		arg.ID,
		arg.Movieid,
		arg.Comment,
		arg.Rating,
		arg.Commenttime,
	)
	return err
}

const deleteReview = `-- name: DeleteReview :exec
DELETE FROM Reviews WHERE ID = $1
`

func (q *Queries) DeleteReview(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteReview, id)
	return err
}

const getReview = `-- name: GetReview :one
SELECT ID, MovieID, Comment, Rating, CommentTime FROM Reviews WHERE ID = $1
`

func (q *Queries) GetReview(ctx context.Context, id uuid.UUID) (Review, error) {
	row := q.db.QueryRow(ctx, getReview, id)
	var i Review
	err := row.Scan(
		&i.ID,
		&i.Movieid,
		&i.Comment,
		&i.Rating,
		&i.Commenttime,
	)
	return i, err
}

const updateReview = `-- name: UpdateReview :exec
UPDATE Reviews SET Comment = $2, Rating = $3 WHERE ID = $1
`

type UpdateReviewParams struct {
	ID      uuid.UUID
	Comment string
	Rating  int
}

func (q *Queries) UpdateReview(ctx context.Context, arg UpdateReviewParams) error {
	_, err := q.db.Exec(ctx, updateReview, arg.ID, arg.Comment, arg.Rating)
	return err
}
