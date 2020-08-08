package post

//
// Repository.
//

import (
	"errors"
	"sync"
)

// ErrNotFound is used when post was not found in store.
var ErrNotFound = errors.New("not found")

// Repository provides access a post store.
type Repository interface {
	Find(id ID) (*Post, error)
	Save(post *Post) error
}

type repository struct {
	mtx   sync.RWMutex
	posts map[ID]*Post
}

// NewRepository creates new repository for accessing posts.
func NewRepository() Repository {
	return &repository{
		posts: make(map[ID]*Post),
	}
}

func (r *repository) Find(id ID) (*Post, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	if p, ok := r.posts[id]; ok {
		return p, nil
	}
	return nil, ErrNotFound
}

func (r *repository) Save(post *Post) error {
	r.mtx.Lock()
	defer r.mtx.Unlock()
	r.posts[post.ID] = post
	return nil
}
