package post

import (
	"encoding/json"
	"fmt"
	"gitlab/elasticsearch-go/internal/pkg/domain"
	"gitlab/elasticsearch-go/internal/pkg/storage"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type Handler struct {
	service service
}

func New(storage storage.PostStorer) Handler {
	return Handler{
		service: service{storage: storage},
	}
}

// POST/api/v1/post
func (h Handler) Create(w http.ResponseWriter, r *http.Request) {
	var req createRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		return
	}
	res, err := h.service.Create(r.Context(), req)
	if err != nil {
		switch err {
		case domain.ErrConflict:
			w.WriteHeader(http.StatusConflict)
		default:
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	body, _ := json.Marshal(res)
	_, _ = w.Write(body)
}

// PACHT/api/v1/post/:id
func (h Handler) Update(w http.ResponseWriter, r *http.Request) {
	var req updateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	req.Id = httprouter.ParamsFromContext(r.Context()).ByName("id")

	if err := h.service.Update(r.Context(), req); err != nil {
		switch err {
		case domain.ErrNotFound:
			w.WriteHeader(http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

// DELETE/api/v1/post/:id
func (h Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id := httprouter.ParamsFromContext(r.Context()).ByName("id")

	if err := h.service.Delete(r.Context(), deleteRequest{Id: id}); err != nil {
		switch err {
		case domain.ErrNotFound:
			w.WriteHeader(http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// GET/api/v1/post/:id
func (h Handler) Get(w http.ResponseWriter, r *http.Request) {
	id := httprouter.ParamsFromContext(r.Context()).ByName("id")
	res, err := h.service.Get(r.Context(), getRequest{Id: id})
	if err != nil {
		switch err {
		case domain.ErrNotFound:
			w.WriteHeader(http.StatusNotFound)
		default:
			w.WriteHeader(http.StatusInternalServerError)
			log.Println(err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	body, _ := json.Marshal(res)
	_, _ = w.Write(body)
}
