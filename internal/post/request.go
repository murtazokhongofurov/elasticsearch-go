package post

type createRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type updateRequest struct {
	Id          string
	Title       string `json:"title"`
	Description string `json:"description`
}

type deleteRequest struct {
	Id string
}

type getRequest struct {
	Id string
}
