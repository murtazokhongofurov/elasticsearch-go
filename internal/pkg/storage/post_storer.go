package storage

import (
	"context"
	"time"
)

type PostStorer interface {
	CreatePost(ctx context.Context, post Post) error
	UpdatePost(ctx context.Context, post Post) error
	DeletePost(ctx context.Context, id string) error
	GetPostById(ctx context.Context, id string) (Post, error)
}

type Post struct {
	Id          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
}
