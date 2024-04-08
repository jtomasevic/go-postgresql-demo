-- name: CreateReview :exec
INSERT INTO Reviews (ID, MovieID, Comment, Rating, CommentTime) VALUES ($1, $2, $3, $4, $5);

-- name: GetReview :one
SELECT ID, MovieID, Comment, Rating, CommentTime FROM Reviews WHERE ID = $1;

-- name: UpdateReview :exec
UPDATE Reviews SET Comment = $2, Rating = $3 WHERE ID = $1;

-- name: DeleteReview :exec
DELETE FROM Reviews WHERE ID = $1;

-- name: AllReviews :many
SELECT ID, MovieID, Comment, Rating, CommentTime FROM Reviews;