-- name: CreateMovie :exec
INSERT INTO Movies (ID, Title, ReleaseYear) VALUES ($1, $2, $3);

-- name: GetMovie :one
SELECT ID, Title, ReleaseYear FROM Movies WHERE ID = $1;

-- name: UpdateMovie :exec
UPDATE Movies SET Title = $2, ReleaseYear = $3 WHERE ID = $1;

-- name: DeleteMovie :exec
DELETE FROM Movies WHERE ID = $1;

-- name: AddMovieToActor :exec
INSERT INTO MovieActors (MovieID, ActorID) VALUES ($1, $2);

-- name: AllMovies :many
SELECT ID, Title, ReleaseYear FROM Movies;

-- name: GetMoviesWithActor :many
SELECT m.ID, m.Title, m.ReleaseYear FROM Movies m INNER JOIN MovieActors ma ON m.ID = ma.MovieID WHERE ma.ActorID = $1;

