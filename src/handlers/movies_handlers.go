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
	getMoviesHandler   = common.NewHttpHandler(GetMovies)
	getMovieHandler    = common.NewHttpHandler(GetMovie)
	newMovieHandler    = common.NewHttpHandler(NewMovie, common.WithTx(true))
	deleteMovieHandler = common.NewHttpHandler(DeleteMovie)
)

func AddMoviesHandlers(mux *http.ServeMux) {

	mux.HandleFunc("GET /movies", getMoviesHandler.HandlerFunc)
	mux.HandleFunc("GET /movie/{id}", getMovieHandler.HandlerFunc)
	mux.HandleFunc("POST /movie", newMovieHandler.HandlerFunc)
	mux.HandleFunc("DELETE /movie/{id}", deleteMovieHandler.HandlerFunc)

}

func GetMovies(api *imdb.ImdbAPI, w http.ResponseWriter, r *http.Request) error {
	movies, err := api.MovieAPI.AllMovies(r.Context())
	if err != nil {
		w.Write([]byte(err.Error()))
		return err
	}
	moviesJSON, err := json.Marshal(movies)
	if err != nil {
		w.Write([]byte(err.Error()))
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(moviesJSON)
	return nil
}

func GetMovie(api *imdb.ImdbAPI, w http.ResponseWriter, r *http.Request) error {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		w.Write([]byte("id not in correct format"))
		return err
	}
	movie, err := api.MovieAPI.GetMovie(r.Context(), id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return err
	}
	movieJSON, err := json.Marshal(movie)
	if err != nil {
		w.Write([]byte(err.Error()))
		return err
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(movieJSON)
	return nil
}

func NewMovie(api *imdb.ImdbAPI, w http.ResponseWriter, r *http.Request) error {
	decoder := json.NewDecoder(r.Body)
	var movieInput model.CreateMovieParams
	err := decoder.Decode(&movieInput)
	if err != nil {
		w.Write([]byte(err.Error()))
		return err
	}
	err = api.MovieAPI.CreateMovie(r.Context(), movieInput)
	if err != nil {
		w.Write([]byte(err.Error()))
		return err
	}
	w.Write([]byte("movie created"))
	return nil
}

func DeleteMovie(api *imdb.ImdbAPI, w http.ResponseWriter, r *http.Request) error {
	id, err := uuid.Parse(r.PathValue("id"))
	if err != nil {
		w.Write([]byte("id not in correct format"))
		return err
	}
	err = api.MovieAPI.DeleteMovie(r.Context(), id)
	if err != nil {
		w.Write([]byte(err.Error()))
		return err
	}
	w.Write([]byte("movie deleted"))
	return nil
}
