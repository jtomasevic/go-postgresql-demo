package imdb

import (
	"context"

	"github.com/google/uuid"
	datasource "github.com/jtomasevic/go-postgresql-demo/src/services/imdb/data_source"
	"github.com/jtomasevic/go-postgresql-demo/src/services/imdb/data_store/repos"
	"github.com/jtomasevic/go-postgresql-demo/src/services/imdb/model"
	"github.com/pkg/errors"
)

type ActorAPI interface {
	CreateActor(ctx context.Context, arg model.CreateActorParams) (uuid.UUID, error)
	GetActor(ctx context.Context, id uuid.UUID) (model.Actor, error)
	UpdateActor(ctx context.Context, arg model.UpdateActorParams) error
	DeleteActor(ctx context.Context, id uuid.UUID) error

	AllActors(ctx context.Context) ([]model.Actor, error)
	AddActorToMovie(ctx context.Context, arg model.AddActorToMovieParams) error
	GetActorsInMovie(ctx context.Context, movieid uuid.UUID) ([]model.Actor, error)
}

type MovieAPI interface {
	CreateMovie(ctx context.Context, arg model.CreateMovieParams) error
	GetMovie(ctx context.Context, id uuid.UUID) (model.Movie, error)
	UpdateMovie(ctx context.Context, arg model.UpdateMovieParams) error
	DeleteMovie(ctx context.Context, id uuid.UUID) error
	AllMovies(ctx context.Context) ([]model.Movie, error)
}

type DirectorAPI interface {
	AllDirectors(ctx context.Context) ([]model.Director, error)
	CreateDirector(ctx context.Context, arg model.CreateDirectorParams) error
	DeleteDirector(ctx context.Context, id uuid.UUID) error
	GetDirector(ctx context.Context, id uuid.UUID) (model.Director, error)
	UpdateDirector(ctx context.Context, arg model.UpdateDirectorParams) error
}

type ReviewAPI interface {
	AllReviews(ctx context.Context) ([]model.Review, error)
	CreateReview(ctx context.Context, arg model.CreateReviewParams) error
	DeleteReview(ctx context.Context, id uuid.UUID) error
	GetReview(ctx context.Context, id uuid.UUID) (model.Review, error)
	UpdateReview(ctx context.Context, arg model.UpdateReviewParams) error
}

type AwardAPI interface {
	ActorAwards(ctx context.Context, actorid *uuid.UUID) ([]model.Award, error)
	AwardsInYear(ctx context.Context, year int) ([]model.Award, error)
	CreateAward(ctx context.Context, arg model.CreateAwardParams) error
	DeleteAward(ctx context.Context, id uuid.UUID) error
	DirectorAwards(ctx context.Context, directorid *uuid.UUID) ([]model.Award, error)
	GetAward(ctx context.Context, id uuid.UUID) (model.Award, error)
	MovieAwards(ctx context.Context, movieid uuid.UUID) ([]model.Award, error)
	UpdateAward(ctx context.Context, arg model.UpdateAwardParams) error
}

type ImdbAPI struct {
	DatSource   datasource.DataSource
	ActorAPI    ActorAPI
	MovieAPI    MovieAPI
	DirectorAPI DirectorAPI
	ReviewAPI   ReviewAPI
	AwardAPI    AwardAPI
}

// This is usuall way to create optional parameters in go
// 1. First we define a struct with the optional parameters
type ApiOptions struct {
	WithTx bool
}

// 2. Then we define a function type that takes a pointer to the struct and returns nothing
type ApiOption func(o *ApiOptions)

// 3. Then we define functions that take the struct which represents the optional parameters and set the value.
func WithTx(tx bool) ApiOption {
	return func(o *ApiOptions) {
		o.WithTx = tx
	}
}

// How to use the optional parameters
// in example above it is used like this:
// api, err := initApi(WithTx(true))

func NewImdbAPI(ctx context.Context, options ...ApiOption) (*ImdbAPI, error) {
	apiOptions := ApiOptions{
		WithTx: false, // just set explicitly default value, just to be clear what's going on
	}
	for _, option := range options {
		option(&apiOptions)
	}
	dataSource := datasource.NewDataSource()
	var dataStore *repos.ImdbStores
	// **** with Transaction
	if apiOptions.WithTx {
		_, err := dataSource.OpenConnection(ctx)
		tx, err := dataSource.StartTransaction(ctx)
		if err != nil {
			return nil, errors.Wrap(err, ErrServerErrorStartTransaction)
		}
		txStarted = true
		dataStore = repos.NewImdbStore(tx)
	} else {
		// **** without Transaction
		_, err := dataSource.OpenConnection(ctx)
		if err != nil {
			return nil, errors.Wrap(err, ErrServerErrorOpenConnection)
		}
		txStarted = false
		dataStore = repos.NewImdbStore(dataSource.GetConnection())
	}

	return &ImdbAPI{
		DatSource: dataSource,
		ActorAPI:  NewActorAPI(dataStore.ActorStore),
		MovieAPI:  NewMovieAPI(dataStore.MovieStore),
	}, nil
}

func (api *ImdbAPI) TearDown(ctx context.Context, options ...TearDownOption) error {
	tearDownOptions := TearDownOptions{}
	for _, option := range options {
		option(&tearDownOptions)
	}
	if txStarted {
		if tearDownOptions.ServiceError != nil {
			err := api.DatSource.RollbackTransaction(ctx)
			if err != nil {
				return errors.Wrap(err, ErrServerCommitTransaction)
			}
		} else {
			err := api.DatSource.CommitTransaction(ctx)
			if err != nil {
				return errors.Wrap(err, ErrServerCommitTransaction)
			}
		}
	} else {
		err := api.DatSource.CloseConnection(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

// This is usuall way to create optional parameters in go
// 1. First we define a struct with the optional parameters
type TearDownOptions struct {
	ServiceError *error
}

// 2. Then we define a function type that takes a pointer to the struct and returns nothing
type TearDownOption func(o *TearDownOptions)

// 3. Then we define functions that take the struct which represents the optional parameters and set the value.
func WithError(err error) TearDownOption {
	return func(options *TearDownOptions) {
		options.ServiceError = &err
	}
}

func NewActorAPI(actorStore repos.ActorStore) *ActorService {
	return &ActorService{
		ActorStore: actorStore,
	}
}

func NewMovieAPI(movieStore repos.MovieStore) *MovieService {
	return &MovieService{
		MovieStore: movieStore,
	}
}

// func NewDirectorAPI(directorStore repos.DirectorStore) *DirectorService {
// 	return &DirectorService{
// 		DirectorStore: directorStore,
// 	}
// }

// func NewReviewAPI(reviewStore repos.ReviewStore) *ReviewService {
// 	return &ReviewService{
// 		ReviewStore: reviewStore,
// 	}
// }

// func NewAwardAPI(awardStore repos.AwardStore) *AwardService {
// 	return &AwardService{
// 		AwardStore: awardStore,
// 	}
// }
