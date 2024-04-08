-- name: CreateDirector :exec
INSERT INTO Directors (ID, Name, BirthYear) VALUES ($1, $2, $3);

-- name: GetDirector :one
SELECT ID, Name, BirthYear FROM Directors WHERE ID = $1;

-- name: UpdateDirector :exec
UPDATE Directors SET Name = $2, BirthYear = $3 WHERE ID = $1;

-- name: DeleteDirector :exec
DELETE FROM Directors WHERE ID = $1;

-- name: AllDirectors :many
SELECT ID, Name, BirthYear FROM Directors;