package model

import (
	"time"

	"github.com/google/uuid"
)

type Actor struct {
	ID        uuid.UUID
	Name      string
	Birthyear *int
}

type Award struct {
	ID         uuid.UUID
	Name       string
	Year       int
	Movieid    uuid.UUID
	Actorid    *uuid.UUID
	Directorid *uuid.UUID
}

type Director struct {
	ID        uuid.UUID
	Name      string
	Birthyear *int
}

type Movie struct {
	ID          uuid.UUID
	Title       string
	Releaseyear int
}

type Movieactor struct {
	Movieid uuid.UUID
	Actorid uuid.UUID
}

type Moviedirector struct {
	Movieid    uuid.UUID
	Directorid uuid.UUID
}

type Review struct {
	ID          uuid.UUID
	Movieid     uuid.UUID
	Comment     string
	Rating      int
	Commenttime time.Time
}