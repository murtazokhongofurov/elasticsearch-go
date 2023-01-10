package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic"
	"github.com/teris-io/shortid"
)

const (
	elasticIndexName = "products"
	elasticTypeName  = "product"
)

type Product struct {
	Id    string  `json:"id"`
	Name  string  `json:"name"`
	Model string  `json:"model"`
	Price float64 `json:"price"`
}

type ProductRequest struct {
	Name  string  `json:"name"`
	Model string  `json:"model"`
	Price float64 `json:"price"`
}

var (
	elasticClient *elastic.Client
)

func main() {
	var err error
	// Create Elastic client and wait for Elasticsearch to be ready
	for {
		elasticClient, err = elastic.NewClient(
			elastic.SetURL("http://elasticsearch:9200"),
			elastic.SetSniff(false),
		)
		if err != nil {
			log.Println(err)
			// Retry every 3 seconds
			time.Sleep(3 * time.Second)
		} else {
			break
		}
	}

	router := gin.Default()
	router.POST("/products", CreateProduct)
	// router.GET("/products", Product)

	if err = router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func CreateProduct(c *gin.Context) {
	var data []ProductRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		errorResponse(c, http.StatusBadRequest, "Malformed request body")
		return
	}

	bulk := elasticClient.Bulk().Index(elasticIndexName).Type(elasticTypeName)
	for _, product := range data {
		temp := Product{
			Id:    shortid.MustGenerate(),
			Name:  product.Name,
			Model: product.Model,
			Price: product.Price,
		}
		bulk.Add(elastic.NewBulkIndexRequest().Id(temp.Id).Doc(temp))
	}

	if _, err := bulk.Do(c.Request.Context()); err != nil {
		log.Println("error", err)
		errorResponse(c, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	c.JSON(http.StatusOK, bulk)
}

func errorResponse(c *gin.Context, code int, err string) {
	c.JSON(code, gin.H{
		"error": err,
	})
}
