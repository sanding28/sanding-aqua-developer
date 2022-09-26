package main

import (
	"net/http"
	"strconv"

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

// func updateProduct( c echo.Context) error{
// 	n := new(Product)
// 	if err := c.Bind(n); err != nil {
// 		return err
// 	}
// 	ID, _ := strconv.Atoi(c.Param("id"))
// 	product[ID].Name = n.Name
// 	return c.JSON(http.StatusOK, product[ID])
// }

func deleteProduct(c echo.Context) error {
	ID, _  := strconv.Atoi(c.Param("id"))
	delete(product, ID)
	return c.NoContent(http.StatusNoContent)  
}

func main(){
	e := echo.New()
	e.POST("/", CreateProduct)
	e.GET("/products", getProduct)
	// e.PUT("/products/:id" , updateProduct)
	e.DELETE("/products/:id", deleteProduct)
	e.Logger.Fatal(e.Start(":1323"))

}
