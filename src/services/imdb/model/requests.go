package model

import (
	"time"

	"github.com/google/uuid"
)

//********************************************************************************
type AddActorToMovieParams struct {
	Actorid uuid.UUID
	Movieid uuid.UUID
}

type CreateActorParams struct {
	Name      string
	Birthyear *int
}

type UpdateActorParams struct {
	ID        uuid.UUID
	Name      string
	Birthyear *int
}
//********************************************************************************
type CreateDirectorParams struct {
	ID        uuid.UUID
	Name      string
	Birthyear *int
}

type UpdateDirectorParams struct {
	ID        uuid.UUID
	Name      string
	Birthyear *int
}

//********************************************************************************
type CreateReviewParams struct {
	ID          uuid.UUID
	Movieid     uuid.UUID
	Comment     string
	Rating      int
	Commenttime time.Time
}

type UpdateReviewParams struct {
	ID      uuid.UUID
	Comment string
	Rating  int
}

//********************************************************************************

type AddMovieToActorParams struct {
	Movieid uuid.UUID
	Actorid uuid.UUID
}

type CreateMovieParams struct {
	ID          uuid.UUID
	Title       string
	Releaseyear int
}

type UpdateMovieParams struct {
	ID          uuid.UUID
	Title       string
	Releaseyear int
}

//**************************************************************************

type CreateAwardParams struct {
	ID         uuid.UUID
	Name       string
	Year       int
	Movieid    uuid.UUID
	Actorid    *uuid.UUID
	Directorid *uuid.UUID
}

type UpdateAwardParams struct {
	Name       string
	Year       int
	Movieid    uuid.UUID
	Actorid    *uuid.UUID
	Directorid *uuid.UUID
}