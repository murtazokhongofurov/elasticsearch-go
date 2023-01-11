package post

import (
	"context"
	"gitlab/elasticsearch-go/internal/pkg/storage"
	"time"

	"github.com/google/uuid"
)

type service struct {
	storage storage.PostStorer
}

func (s service) Create(ctx context.Context, req createRequest) (createResponse, error) {
	id := uuid.New().String()
	create := time.Now().UTC()

	doc := storage.Post{
		Id:          id,
		Title:       req.Title,
		Description: req.Description,
		CreatedAt:   &create,
	}
	if err := s.storage.CreatePost(ctx, doc); err != nil {
		return createResponse{}, err
	}
	return createResponse{Id: id}, nil
}

func (s service) Update(ctx context.Context, req updateRequest) error {
	doc := storage.Post{
		Id:          req.Id,
		Title:       req.Title,
		Description: req.Description,
	}

	if err := s.storage.UpdatePost(ctx, doc); err != nil {
		return err
	}
	return nil
}

func (s service) Delete(ctx context.Context, req deleteRequest) error {
	if err := s.storage.DeletePost(ctx, req.Id); err != nil {
		return err
	}
	return nil
}

func (s service) Get(ctx context.Context, req getRequest) (getResponse, error) {
	post, err := s.storage.GetPostById(ctx, req.Id)
	if err != nil {
		return getResponse{}, err
	}
	return getResponse{
		Id:          post.Id,
		Title:       post.Title,
		Description: post.Description,
		CreatedAt:   *post.CreatedAt,
	}, nil
}
