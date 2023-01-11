package main

import (
	"gitlab/elasticsearch-go/internal/pkg/storage/elasticsearch"
	"gitlab/elasticsearch-go/internal/post"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	elastic, err := elasticsearch.New([]string{"http://localhost:9200"})
	if err != nil {
		log.Fatalln(err)
	}
	if err := elastic.CreateIndex("post"); err != nil {
		log.Fatalln(err)
	}
	storage, err := elasticsearch.NewPostStorage(*elastic)
	if err != nil {
		log.Fatal(err)
	}

	postApi := post.New(storage)

	router := httprouter.New()
	router.HandlerFunc("POST", "/api/v1/post", postApi.Create)
	router.HandlerFunc("PATCH", "/api/v1/post/:id", postApi.Update)
	router.HandlerFunc("DELETE", "/api/v1/post/:id", postApi.Delete)
	router.HandlerFunc("GET", "/api/v1/post/:id", postApi.Get)

	log.Fatalln(http.ListenAndServe(":8000", router))
}
