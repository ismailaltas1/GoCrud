// create package main
package main

import (
	"awesomeProject/main/handler"
	"context"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func main() {

	e := echo.New()

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(ctx)

	bookHandler := handler.NewBookHandler(client.Database("bookstore").Collection("books"))
	e.GET("/books", bookHandler.GetBooks)
	e.GET("/books/:id", bookHandler.GetBook)
	e.POST("/books", bookHandler.PostBook)
	e.PUT("/books/:id", bookHandler.PutBook)
	e.DELETE("books/:id", bookHandler.DeleteBook)
	e.Logger.Fatal(e.Start(":8080"))

}
