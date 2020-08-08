package post

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/transport"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

var errBadPostID = errors.New("Bad post ID")

// MakeHandler returns a handler for the posts service.
func MakeHandler(ps Service, logger log.Logger) http.Handler {
	opts := []kithttp.ServerOption{
		kithttp.ServerErrorHandler(transport.NewLogErrorHandler(logger)),
		kithttp.ServerErrorEncoder(encodeError),
	}

	loadPostHandler := kithttp.NewServer(
		makeLoadPostEndpoint(ps),
		decodeLoadPostRequest,
		encodeResponse,
		opts...,
	)

	r := mux.NewRouter()

	r.Handle("/posts/v1/{id}", loadPostHandler).Methods("GET")

	return r
}

func decodeLoadPostRequest(_ctx context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, errBadPostID
	}
	return loadPostRequest{ID: ID(id)}, nil
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(interface{ error() error }); ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	case ErrNotFound:
		w.WriteHeader(http.StatusNotFound)
	case ErrInvalidArgument:
		w.WriteHeader(http.StatusBadRequest)
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}
