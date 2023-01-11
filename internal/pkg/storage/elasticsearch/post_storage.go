package elasticsearch

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"gitlab/elasticsearch-go/internal/pkg/storage"
	"time"

	"github.com/elastic/go-elasticsearch/v7/esapi"
)

// var _ storage.PostStorer = PostStorage{}

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
	req := esapi.CreateRequest{
		Index:      p.elastic.alias,
		DocumentID: post.Id,
		Body:       bytes.NewReader(body),
	}

	ctx, cancel := context.WithTimeout(ctx, p.timeout)
	defer cancel()
	res, err := req.Do(ctx, p.elastic.client)

	if err != nil {
		return fmt.Errorf("insert: request: %w", err)
	}
	if res.StatusCode == 409 {
		fmt.Println("conflict error")
		return fmt.Errorf("conflict error")
	}
	if res.IsError() {
		return fmt.Errorf("insert: response: %s", res.String())
	}
	fmt.Println("response", res)
	return nil
}

func (p PostStorage) UpdatePost(ctx context.Context, post storage.Post) error {
	body, err := json.Marshal(post)
	if err != nil {
		return fmt.Errorf("update: marshal: %w", err)
	}
	req := esapi.UpdateRequest{
		Index:      p.elastic.alias,
		DocumentID: post.Id,
		Body:       bytes.NewReader([]byte(fmt.Sprintf(`{doc:%s}`, body))),
	}

	ctx, cancel := context.WithTimeout(ctx, p.timeout)
	defer cancel()

	res, err := req.Do(ctx, p.elastic.client)
	if err != nil {
		return fmt.Errorf("update: request: %w", err)
	}
	if res.StatusCode == 409 {
		return fmt.Errorf("error not fount")
	}

	if res.IsError() {
		return fmt.Errorf("update: response: %s", res.String())
	}

	return nil
}

func (p PostStorage) DeletePost(ctx context.Context, id string) error {
	req := esapi.DeleteRequest{
		Index:      p.elastic.alias,
		DocumentID: id,
	}

	ctx, cancel := context.WithTimeout(ctx, p.timeout)
	defer cancel()
	res, err := req.Do(ctx, p.elastic.client)
	if err != nil {
		return fmt.Errorf("delete: request: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode == 404 {
		return fmt.Errorf("error not fount")
	}

	if res.IsError() {
		return fmt.Errorf("delete: response: %s", res.String())
	}

	return nil
}

func (p PostStorage) GetPostById(ctx context.Context, id string) (storage.Post, error) {
	req := esapi.GetRequest{
		Index:      p.elastic.alias,
		DocumentID: id,
	}

	ctx, cancel := context.WithTimeout(ctx, p.timeout)
	defer cancel()
	res, err := req.Do(ctx, p.elastic.client)
	if err != nil {
		return storage.Post{}, fmt.Errorf("get: request: %w", err)
	}
	defer res.Body.Close()
	if res.StatusCode == 404 {
		return storage.Post{}, fmt.Errorf("error not fount")
	}
	if res.IsError() {
		return storage.Post{}, fmt.Errorf("get: response: %s", res.String())
	}
	var (
		post storage.Post
		body document
	)
	body.Source = &post
	if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
		return storage.Post{}, fmt.Errorf("get: decode: %w", err)
	}
	return post, nil
}

