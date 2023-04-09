// create package main
package main

import (
	"awesomeProject/main/handler"
	"github.com/labstack/echo/v4"
)

func main() {

	e := echo.New()
	e.GET("/books", handler.GetBooks)
	e.GET("/books/:id", handler.GetBook)
	e.POST("/books", handler.PostBook)
	e.PUT("/books/:id", handler.PutBook)
	e.DELETE("books/:id", handler.DeleteBook)
	e.Logger.Fatal(e.Start(":8080"))

}
