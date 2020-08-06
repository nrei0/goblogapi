package post

import (
	"errors"
	"math"
	"time"
)

const (
	// TitleCharsLimit is limitation in number of charaters for post title.
	TitleCharsLimit = 80
	// ContentCharsLimit is limitation in number of characters for post content.
	ContentCharsLimit = 2097152
)

// ID is unique post identifier.
type ID string

// Title is 120 charaters post's title.
type Title string

// Content is UTF-8 post's text with limitation of 2097152 characters.
type Content string

// Post is general model for keeping post's information.
type Post struct {
	ID         ID
	Title      Title
	Content    Content
	CreateDate time.Time
	ModifyDate time.Time
}

// New creates new Post.
// Plese note that title is limited by 120 chars and content by 2097152.
// Supposed that ID should be generated on database side.
func New(id ID, title Title, content Content) *Post {
	titleNumChars := int8(math.Min(TitleCharsLimit, float64(len(title))))
	contentNumChars := int8(math.Min(ContentCharsLimit, float64(len(content))))
	return &Post{
		ID:         id,
		Title:      title[:titleNumChars],
		Content:    content[:contentNumChars],
		CreateDate: time.Now().UTC(),
		ModifyDate: time.Time{}.UTC(),
	}
}

// Repository provides access a post store.
type Repository interface {
	Find(id ID) (*Post, error)
}

// ErrNotFound is used when post was not found in store.
var ErrNotFound = errors.New("not found")
