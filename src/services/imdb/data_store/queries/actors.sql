-- name: CreateActor :exec
INSERT INTO Actors (ID, Name, BirthYear) VALUES ($1, $2, $3);

-- name: GetActor :one
SELECT ID, Name, BirthYear FROM Actors WHERE ID = $1;

-- name: UpdateActor :exec
UPDATE Actors SET Name = $2, BirthYear = $3 WHERE ID = $1;

-- name: DeleteActor :exec
DELETE FROM Actors WHERE ID = $1;

-- name: AddActorToMovie :exec
INSERT INTO MovieActors (ActorID, MovieID) VALUES ($1, $2);

-- name: AllActors :many
SELECT ID, Name, BirthYear FROM Actors;

-- name: GetActorsInMovie :many
SELECT a.ID, a.Name, a.BirthYear FROM Actors a INNER JOIN MovieActors ma ON a.ID = ma.ActorID WHERE ma.MovieID = $1;
