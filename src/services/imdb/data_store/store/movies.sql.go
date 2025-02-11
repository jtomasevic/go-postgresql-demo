// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: movies.sql

package store

import (
	"context"

	"github.com/google/uuid"
)

const addMovieToActor = `-- name: AddMovieToActor :exec
INSERT INTO MovieActors (MovieID, ActorID) VALUES ($1, $2)
`

type AddMovieToActorParams struct {
	Movieid uuid.UUID
	Actorid uuid.UUID
}

func (q *Queries) AddMovieToActor(ctx context.Context, arg AddMovieToActorParams) error {
	_, err := q.db.Exec(ctx, addMovieToActor, arg.Movieid, arg.Actorid)
	return err
}

const allMovies = `-- name: AllMovies :many
SELECT ID, Title, ReleaseYear FROM Movies
`

func (q *Queries) AllMovies(ctx context.Context) ([]Movie, error) {
	rows, err := q.db.Query(ctx, allMovies)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Movie
	for rows.Next() {
		var i Movie
		if err := rows.Scan(&i.ID, &i.Title, &i.Releaseyear); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const createMovie = `-- name: CreateMovie :exec
INSERT INTO Movies (ID, Title, ReleaseYear) VALUES ($1, $2, $3)
`

type CreateMovieParams struct {
	ID          uuid.UUID
	Title       string
	Releaseyear int
}

func (q *Queries) CreateMovie(ctx context.Context, arg CreateMovieParams) error {
	_, err := q.db.Exec(ctx, createMovie, arg.ID, arg.Title, arg.Releaseyear)
	return err
}

const deleteMovie = `-- name: DeleteMovie :exec
DELETE FROM Movies WHERE ID = $1
`

func (q *Queries) DeleteMovie(ctx context.Context, id uuid.UUID) error {
	_, err := q.db.Exec(ctx, deleteMovie, id)
	return err
}

const getMovie = `-- name: GetMovie :one
SELECT ID, Title, ReleaseYear FROM Movies WHERE ID = $1
`

func (q *Queries) GetMovie(ctx context.Context, id uuid.UUID) (Movie, error) {
	row := q.db.QueryRow(ctx, getMovie, id)
	var i Movie
	err := row.Scan(&i.ID, &i.Title, &i.Releaseyear)
	return i, err
}

const getMoviesWithActor = `-- name: GetMoviesWithActor :many
SELECT m.ID, m.Title, m.ReleaseYear FROM Movies m INNER JOIN MovieActors ma ON m.ID = ma.MovieID WHERE ma.ActorID = $1
`

func (q *Queries) GetMoviesWithActor(ctx context.Context, actorid uuid.UUID) ([]Movie, error) {
	rows, err := q.db.Query(ctx, getMoviesWithActor, actorid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Movie
	for rows.Next() {
		var i Movie
		if err := rows.Scan(&i.ID, &i.Title, &i.Releaseyear); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateMovie = `-- name: UpdateMovie :exec
UPDATE Movies SET Title = $2, ReleaseYear = $3 WHERE ID = $1
`

type UpdateMovieParams struct {
	ID          uuid.UUID
	Title       string
	Releaseyear int
}

func (q *Queries) UpdateMovie(ctx context.Context, arg UpdateMovieParams) error {
	_, err := q.db.Exec(ctx, updateMovie, arg.ID, arg.Title, arg.Releaseyear)
	return err
}
