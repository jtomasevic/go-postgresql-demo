package server

import (
	"net/http"

	"github.com/jtomasevic/go-postgresql-demo/src/handlers"
)

func StartServer() {
	mux := http.NewServeMux()
	handlers.AddActorHandlers(mux)
	handlers.AddMoviesHandlers(mux)

	http.ListenAndServe("localhost:8090", mux)
}
