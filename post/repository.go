package post

//
// Repository.
//

import (
	"fmt"
	"errors"
	"sync"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

const tableName = "Posts"

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
	db *dynamodb.DynamoDB
}

// NewRepository creates new repository for accessing posts.
func NewRepository(db *dynamodb.DynamoDB) Repository {
	return &repository{
		posts: make(map[ID]*Post),
		db: db,
	}
}

func (r *repository) Find(id ID) (*Post, error) {
	r.mtx.Lock()
	defer r.mtx.Unlock()

	// tbd @nrei instead of local posts use dynamodb database
	result, err := r.db.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"ID": {
				S: aws.String(string(id)),
			},
		},
	})

	fmt.Printf("%s, %s", result, err)

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
