package main

import (
	"gitlab/elasticsearch-go/internal/pkg/storage/elasticsearch"
	"log"
	"github.com/julienschmidt/httprouter"
)

func main() {
	elastic, err := elasticsearch.New([]string("http://localhost:9200"))
	if err != nil {
		log.Fatal(err)
	}
	if err := elastic.CreateIndex("post"); err != nil {
		log.Fatal(err)
	}
	storage, err := elasticsearch.NewPostStorage(*elastic)
	if err != nil {
		log.Fatal(err)
	}

	postApi := post.New(storage)

	router := httprouter.New()
}
