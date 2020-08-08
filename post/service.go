package post

import (
	"errors"
)

// ErrInvalidArgument is used when invalid argument passes through one of the services.
var ErrInvalidArgument = errors.New("invalid argument(s)")

// Service is the service that expose posts methods.
type Service interface {
	LoadPost(id ID) (Post, error)
}

type service struct {
	posts Repository
}

func (s *service) LoadPost(id ID) (Post, error) {
	if id == "" {
		return Post{}, ErrInvalidArgument
	}

	p, err := s.posts.Find(id)
	if err != nil {
		return Post{}, err
	}

	return Post{
		ID:         p.ID,
		Title:      p.Title,
		Content:    p.Content,
		CreateDate: p.CreateDate,
		ModifyDate: p.ModifyDate,
	}, nil
}

// NewService provides interace for accessing different repositories.
func NewService(postRepository Repository) Service {
	return &service{
		posts: postRepository,
	}
}
