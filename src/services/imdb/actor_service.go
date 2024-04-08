package imdb

import (
	"context"

	"github.com/pkg/errors"

	"github.com/google/uuid"
	datasource "github.com/jtomasevic/go-postgresql-demo/src/services/imdb/data_source"
	"github.com/jtomasevic/go-postgresql-demo/src/services/imdb/data_store/repos"
	"github.com/jtomasevic/go-postgresql-demo/src/services/imdb/data_store/store"

	"github.com/jtomasevic/go-postgresql-demo/src/services/imdb/model"
)

type ActorService struct {
	DataSource datasource.DataSource
	ActorStore repos.ActorStore
}

var (
	ErrCreateActor      = "error creating actor in service layer\n"
	ErrAllActors        = "error fetching all actors in service layer\n"
	ErrUpdateActor      = "error updating actor in service layer\n"
	ErrGetActor         = "error fetching actor in service layer\n"
	ErrDeleteActor      = "error deleting actor in service layer\n"
	ErrAddActorToMovie  = "error adding actor to movie in service layer\n"
	ErrGetActorsInMovie = "error fetching actors in movie in service layer\n"
)

func (service *ActorService) CreateActor(ctx context.Context, arg model.CreateActorParams) (uuid.UUID, error) {
	id := uuid.New()
	err := service.ActorStore.CreateActor(ctx, store.CreateActorParams{
		ID:        id,
		Name:      arg.Name,
		Birthyear: arg.Birthyear,
	})
	if err != nil {
		return uuid.UUID{}, errors.Wrap(err, ErrCreateActor)
	}
	return id, nil
}

func (service *ActorService) GetActor(ctx context.Context, id uuid.UUID) (model.Actor, error) {
	actor, err := service.ActorStore.GetActor(ctx, id)
	if err != nil {
		return model.Actor{}, errors.Wrap(err, ErrGetActor)
	}
	return model.Actor(actor), nil
}

func (service *ActorService) UpdateActor(ctx context.Context, arg model.UpdateActorParams) error {
	err := service.ActorStore.UpdateActor(ctx, store.UpdateActorParams(arg))
	if err != nil {
		return errors.Wrap(err, ErrUpdateActor)
	}
	return nil
}

func (service *ActorService) DeleteActor(ctx context.Context, id uuid.UUID) error {
	err := service.ActorStore.DeleteActor(ctx, id)
	if err != nil {
		return errors.Wrap(err, ErrDeleteActor)
	}
	return nil
}

func (service *ActorService) AllActors(ctx context.Context) ([]model.Actor, error) {
	actors, err := service.ActorStore.AllActors(ctx)
	if err != nil {
		return nil, errors.Wrap(err, ErrAllActors)
	}
	result := make([]model.Actor, 0, len(actors))
	for _, actor := range actors {
		result = append(result, model.Actor(actor))
	}
	return result, nil
}

func (service *ActorService) GetActorsInMovie(ctx context.Context, movieid uuid.UUID) ([]model.Actor, error) {
	actors, err := service.ActorStore.GetActorsInMovie(ctx, movieid)
	if err != nil {
		return nil, errors.Wrap(err, ErrGetActorsInMovie)
	}
	result := make([]model.Actor, 0, len(actors))
	for _, actor := range actors {
		result = append(result, model.Actor(actor))
	}
	return result, nil
}

func (service *ActorService) AddActorToMovie(ctx context.Context, arg model.AddActorToMovieParams) error {
	err := service.ActorStore.AddActorToMovie(ctx, store.AddActorToMovieParams(arg))
	if err != nil {
		return errors.Wrap(err, ErrAddActorToMovie)
	}
	return nil
}
