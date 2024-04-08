package repos

import (
	"context"

	"github.com/google/uuid"
	datasource "github.com/jtomasevic/go-postgresql-demo/src/services/imdb/data_source"
	"github.com/jtomasevic/go-postgresql-demo/src/services/imdb/data_store/store"
)

type ActorStore interface {
	AddActorToMovie(ctx context.Context, arg store.AddActorToMovieParams) error
	AllActors(ctx context.Context) ([]store.Actor, error)
	CreateActor(ctx context.Context, arg store.CreateActorParams) error
	DeleteActor(ctx context.Context, id uuid.UUID) error
	GetActor(ctx context.Context, id uuid.UUID) (store.Actor, error)
	GetActorsInMovie(ctx context.Context, movieid uuid.UUID) ([]store.Actor, error)
	UpdateActor(ctx context.Context, arg store.UpdateActorParams) error
}

type DirectorStore interface {
	AllDirectors(ctx context.Context) ([]store.Director, error)
	CreateDirector(ctx context.Context, arg store.CreateDirectorParams) error
	DeleteDirector(ctx context.Context, id uuid.UUID) error
	GetDirector(ctx context.Context, id uuid.UUID) (store.Director, error)
	UpdateDirector(ctx context.Context, arg store.UpdateDirectorParams) error
}

type MovieStore interface {
	AllMovies(ctx context.Context) ([]store.Movie, error)
	CreateMovie(ctx context.Context, arg store.CreateMovieParams) error
	DeleteMovie(ctx context.Context, id uuid.UUID) error
	GetMovie(ctx context.Context, id uuid.UUID) (store.Movie, error)
	UpdateMovie(ctx context.Context, arg store.UpdateMovieParams) error
}

type ReviewStore interface {
	AllReviews(ctx context.Context) ([]store.Review, error)
	CreateReview(ctx context.Context, arg store.CreateReviewParams) error
	DeleteReview(ctx context.Context, id uuid.UUID) error
	GetReview(ctx context.Context, id uuid.UUID) (store.Review, error)
	UpdateReview(ctx context.Context, arg store.UpdateReviewParams) error
}

type AwardStore interface {
	ActorAwards(ctx context.Context, actorid *uuid.UUID) ([]store.Award, error)
	AwardsInYear(ctx context.Context, year int) ([]store.Award, error)
	CreateAward(ctx context.Context, arg store.CreateAwardParams) error
	DeleteAward(ctx context.Context, id uuid.UUID) error
	DirectorAwards(ctx context.Context, directorid *uuid.UUID) ([]store.Award, error)
	GetAward(ctx context.Context, id uuid.UUID) (store.Award, error)
	MovieAwards(ctx context.Context, movieid uuid.UUID) ([]store.Award, error)
	UpdateAward(ctx context.Context, arg store.UpdateAwardParams) error
}

type ImdbStores struct {
	Datastore     datasource.DataSource
	ActorStore    ActorStore
	DirectorStore DirectorStore
	MovieStore    MovieStore
	ReviewStore   ReviewStore
	AwardStore    AwardStore
}

func NewImdbStore(db store.DBTX) *ImdbStores {
	queries := store.New(db)
	return &ImdbStores{
		ActorStore:    queries,
		DirectorStore: queries,
		MovieStore:    queries,
		ReviewStore:   queries,
		AwardStore:    queries,
	}
}
