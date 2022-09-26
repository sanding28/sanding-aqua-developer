package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Product struct {
	ID int `json: "id"`
	Name string `json : "name"`
}

var (
	product = map[int]*Product{}
	nomor = 1 
)



func CreateProduct(c echo.Context)error{
	p := &Product{
		ID: nomor,
	}
	if err := c.Bind(p); err != nil {
		return err
	}
	product[p.ID] = p
	nomor++
	return c.JSON(http.StatusCreated, p)
}

func getProduct(c echo.Context) error{
	return c.JSON(http.StatusOK, product)
}

func main(){
	e := echo.New()
	e.POST("/", CreateProduct)
	e.GET("/products", getProduct)
	e.Logger.Fatal(e.Start(":1323"))

}
