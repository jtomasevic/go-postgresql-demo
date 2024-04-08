package imdb

import (
	"context"

	"github.com/google/uuid"
	datasource "github.com/jtomasevic/go-postgresql-demo/src/services/imdb/data_source"
	"github.com/jtomasevic/go-postgresql-demo/src/services/imdb/data_store/repos"
	"github.com/jtomasevic/go-postgresql-demo/src/services/imdb/data_store/store"
	"github.com/jtomasevic/go-postgresql-demo/src/services/imdb/model"
	"github.com/pkg/errors"
)

type MovieService struct {
	DataSource datasource.DataSource
	MovieStore repos.MovieStore
}

var (
	ErrCreateMovie = "error creating movie in service layer\n"
	ErrAllMovies   = "error fetching all movies in service layer\n"
	ErrUpdateMovie = "error updating movie in service layer\n"
	ErrGetMovie    = "error fetching movie in service layer\n"
	ErrDeleteMovie = "error deleting movie in service layer\n"
)

func (service *MovieService) CreateMovie(ctx context.Context, arg model.CreateMovieParams) error {
	err := service.MovieStore.CreateMovie(ctx, store.CreateMovieParams(arg))
	if err != nil {
		return errors.Wrap(err, ErrCreateMovie)
	}
	return nil
}

func (service *MovieService) GetMovie(ctx context.Context, id uuid.UUID) (model.Movie, error) {
	movie, err := service.MovieStore.GetMovie(ctx, id)
	if err != nil {
		return model.Movie{}, errors.Wrap(err, ErrGetMovie)
	}
	return model.Movie(movie), nil
}

func (service *MovieService) UpdateMovie(ctx context.Context, arg model.UpdateMovieParams) error {
	err := service.MovieStore.UpdateMovie(ctx, store.UpdateMovieParams(arg))
	if err != nil {
		return errors.Wrap(err, ErrUpdateMovie)
	}
	return nil
}

func (service *MovieService) DeleteMovie(ctx context.Context, id uuid.UUID) error {
	err := service.MovieStore.DeleteMovie(ctx, id)
	if err != nil {
		return errors.Wrap(err, ErrDeleteMovie)
	}
	return nil
}

func (service *MovieService) AllMovies(ctx context.Context) ([]model.Movie, error) {
	movies, err := service.MovieStore.AllMovies(ctx)
	if err != nil {
		return []model.Movie{}, errors.Wrap(err, ErrAllMovies)
	}
	result := make([]model.Movie, 0, len(movies))
	for _, movie := range movies {
		result = append(result, model.Movie(movie))
	}
	return result, nil
}
