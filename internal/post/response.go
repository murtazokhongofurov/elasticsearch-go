package post

import "time"

type createResponse struct {
	Id string `json:"id"`
}

type getResponse struct {
	Id          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
}
