package elasticsearch

import (
	"context"
	"encoding/json"
	"fmt"
	"gitlab/elasticsearch-go/internal/pkg/storage"
	"time"
)

var _ storage.PostStorer = PostStorage{}

type PostStorage struct {
	elastic ElasticSearch
	timeout time.Duration
}

func NewPostStorage(elastic ElasticSearch) (PostStorage, error) {
	return PostStorage{
		elastic: elastic,
		timeout: time.Second * 10,
	}, nil
}

func (p PostStorage) CreatePost(ctx context.Context, post storage.Post) error {
	body, err := json.Marshal(post)
	if err != nil {
		return fmt.Errorf("insert: marshal: %w", err)
	}

	
}
