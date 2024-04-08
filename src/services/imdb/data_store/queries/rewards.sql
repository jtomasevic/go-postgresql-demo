-- name: CreateAward :exec
INSERT INTO Awards (ID, Name, Year, MovieID, ActorID, DirectorID)
VALUES ($1, $2, $3, $4, $5, $6);

-- name: GetAward :one
SELECT ID, Name, Year, MovieID, ActorID, DirectorID FROM Awards WHERE ID = $1;

-- name: UpdateAward :exec
UPDATE Awards
SET Name = $1, Year = $2, MovieID = $3, ActorID = $4, DirectorID = $5
WHERE ID = $1;

-- name: DeleteAward :exec
DELETE FROM Awards WHERE ID = $1;


-- name: MovieAwards :many
SELECT ID, Name, Year, MovieID, ActorID, DirectorID FROM Awards WHERE MovieID = $1;

-- name: ActorAwards :many
SELECT ID, Name, Year, MovieID, ActorID, DirectorID FROM Awards WHERE ActorID = $1;

-- name: DirectorAwards :many
SELECT ID, Name, Year, MovieID, ActorID, DirectorID FROM Awards WHERE DirectorID = $1;

-- name: AwardsInYear :many
SELECT ID, Name, Year, MovieID, ActorID, DirectorID FROM Awards WHERE Year = $1;

-- name: AwardsByName :one
SELECT ID, Name, Year, MovieID, ActorID, DirectorID FROM Awards WHERE Name = $1;

