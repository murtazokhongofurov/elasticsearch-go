package main


type Product struct {
	Id int `json:"id"`
	Name string `json:"name"`
	Model string `json:"model"`
	Price float64 `json:"price"`
}