package common

import (
	"net/http"

	"github.com/jtomasevic/go-postgresql-demo/src/services/imdb"
	"github.com/pkg/errors"
)

var (
	ErrServerErrorInitApi    = "error in server while initializing api"
	ErrServerErrorInTearDown = "error in server while tearing down"
)

type HttpHandler struct {
	HandlerMethod HandlerMethod
	WithTx        bool
}

func NewHttpHandler(handlerMethod HandlerMethod, options ...HandlerOption) *HttpHandler {
	handlerOptions := HandlerOptions{
		WithTx: false, // just set explicitly default value, just to be clear what's going on
	}
	for _, option := range options {
		option(&handlerOptions)
	}
	return &HttpHandler{
		WithTx:        handlerOptions.WithTx,
		HandlerMethod: handlerMethod,
	}
}

func (httpHadler *HttpHandler) HandlerFunc(w http.ResponseWriter, r *http.Request) {
	api, err := imdb.NewImdbAPI(r.Context(), imdb.WithTx(httpHadler.WithTx))
	if err != nil {
		err = errors.Wrap(err, ErrServerErrorInitApi)
		w.Write([]byte(err.Error()))
		return
	}
	err = httpHadler.HandlerMethod(api, w, r)
	if err != nil {
		err = api.TearDown(r.Context(), imdb.WithError(err))
		if err != nil {
			err = errors.Wrap(err, ErrServerErrorInTearDown)
		}
		return
	}
	err = api.TearDown(r.Context())
	if err != nil {
		err = errors.Wrap(err, ErrServerErrorInTearDown)
		w.Write([]byte(err.Error()))
	}
}

type HandlerMethod func(api *imdb.ImdbAPI, w http.ResponseWriter, r *http.Request) error

// This is usuall way to create optional parameters in go
// 1. First we define a struct with the optional parameters
type HandlerOptions struct {
	WithTx bool
}

// 2. Then we define a function type that takes a pointer to the struct and returns nothing
type HandlerOption func(o *HandlerOptions)

// 3. Then we define functions that take the struct which represents the optional parameters and set the value.
func WithTx(tx bool) HandlerOption {
	return func(options *HandlerOptions) {
		options.WithTx = tx
	}
}
