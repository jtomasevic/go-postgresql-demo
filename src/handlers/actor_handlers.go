package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/jtomasevic/go-postgresql-demo/src/handlers/common"
	"github.com/jtomasevic/go-postgresql-demo/src/services/imdb"
	"github.com/jtomasevic/go-postgresql-demo/src/services/imdb/model"
)

var (
	getActorsHandler   = common.NewHttpHandler(GetActors)
	getActorHandler    = common.NewHttpHandler(GetActor)
	saveActorHandler   = common.NewHttpHandler(NewActor, common.WithTx(true))
	deleteActorHandler = common.NewHttpHandler(DeleteActor)
)

func AddActorHandlers(mux *http.ServeMux) {
	mux.HandleFunc("GET /actors", getActorsHandler.HandlerFunc)
	mux.HandleFunc("GET /actor/{id}", getActorHandler.HandlerFunc)
	mux.HandleFunc("POST /actor", saveActorHandler.HandlerFunc)
	mux.HandleFunc("DELETE /actor/{id}", saveActorHandler.HandlerFunc)
}

func GetActors(api *imdb.ImdbAPI, w http.ResponseWriter, r *http.Request) error {
	actors, err := api.ActorAPI.AllActors(r.Context())
	if err != nil {
		w.Write([]byte(err.Error()))
		return err
	}
	actorsJSON, err := json.Marshal(actors)
	if err != nil {
		w.Write([]byte(err.Error()))
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(actorsJSON)
	return nil
}

func GetActor(api *imdb.ImdbAPI, w http.ResponseWriter, r *http.Request) error {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		w.Write([]byte("id not in correct format"))
		return err
	}
	actor, err := api.ActorAPI.GetActor(r.Context(), id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return err
	}
	actorJSON, err := json.Marshal(actor)
	if err != nil {
		w.Write([]byte(err.Error()))
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(actorJSON)
	return nil
}

func NewActor(api *imdb.ImdbAPI, w http.ResponseWriter, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	var actorInput model.CreateActorParams
	err := decoder.Decode(&actorInput)
	if err != nil {
		w.Write([]byte(err.Error()))
		return nil
	}
	id, err := api.ActorAPI.CreateActor(r.Context(), actorInput)
	if err != nil {
		w.Write([]byte(err.Error()))
		return nil
	}
	w.Write([]byte(id.String()))
	return nil
}

func DeleteActor(api *imdb.ImdbAPI, w http.ResponseWriter, r *http.Request) error {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		w.Write([]byte("id not in correct format"))
		return err
	}
	err = api.ActorAPI.DeleteActor(r.Context(), id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return err
	}
	w.Write([]byte(id.String()))
	return nil
}
