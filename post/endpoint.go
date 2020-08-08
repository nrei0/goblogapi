package post

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type loadPostRequest struct {
	ID ID
}

type loadPostResponse struct {
	Post *Post `json:"post,omitempty"`
	Err  error `json:"error,omitempty"`
}

func (r loadPostResponse) error() error {
	return r.Err
}

func makeLoadPostEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(loadPostRequest)
		p, err := s.LoadPost(req.ID)
		return loadPostResponse{Post: &p, Err: err}, nil
	}
}
