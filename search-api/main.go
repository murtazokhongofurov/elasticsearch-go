package main

import (
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic"
	"github.com/teris-io/shortid"

)
type Product struct {
	Id 		int 	`json:"id"`
	Name 	string 	`json:"name"`
	Model 	string 	`json:"model"`
	Price 	float64 `json:"price"`
}



type ProductRequest	struct {
    Name 	string 	`json:"name"`
    Model 	string 	`json:"model"`
    Price 	float64 `json:"price"`
}

func main() {
	var err error

	for { elasticClient, err := elastic.NewClient(
		elastic.SetURL("http://localhost:9200"),
        elastic.SetSniff(false)
	)
	if err != nil {
		fmt.Println("error",err)
		}else {
			break
		}
	}

	router := gin.Default()
	router.POST("/products", Product)
	router.GET("/products", Product)

	if err = router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}


func CreateProduct(c *gin.Context) {
	var data []ProductRequest
	if err := c.ShouldBindJSON(&data); err!= nil {
		errorResponse(c, http.StatusBadRequest, "Malformed request body")
		return
	}

	bulk := elasticClient.Bulk().Index(elasticIndexName).Type(elasticTypeName)
	for _, product := range data {
		temp := Product{
			Id:   shortid.MustGenerate(),
            Name: product.Name,
            Model: product.Model,
            Price: product.Price,
        }
		bulk.Add(elastic.NewBulkIndexRequest().Id(temp.Id).Doc(temp))
		}

		if _, err := bulk.Do(c.Request.Context); err != nil {
			log.Println("error", err)
			errorResponse(c, http.StatusInternalServerError, "Internal Server Error")
			return
		}
		c.JSON(http.StatusOK)
}