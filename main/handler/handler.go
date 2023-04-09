package handler

import (
	"context"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

type Book struct {
	ID     primitive.ObjectID `json:"_id"`
	Title  string             `json:"title"`
	Author string             `json:"author"`
}
type BookHandler struct {
	collection *mongo.Collection
}

func NewBookHandler(collection *mongo.Collection) *BookHandler {
	return &BookHandler{collection: collection}
}

func (h *BookHandler) GetBooks(c echo.Context) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := h.collection.Find(ctx, bson.M{})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "could not retrieve books.")
	}
	defer cursor.Close(ctx)

	var books []Book
	if err = cursor.All(ctx, &books); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "could not decode books")
	}

	return c.JSON(http.StatusOK, books)
}

func (h *BookHandler) GetBook(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}
	var book Book
	err = h.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&book)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return echo.NewHTTPError(http.StatusNotFound, "book not found")
		} else {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
	}
	return c.JSON(http.StatusOK, book)
}

func (h *BookHandler) PostBook(c echo.Context) error {
	var book Book
	if err := c.Bind(&book); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	result, err := h.collection.InsertOne(context.Background(), book)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	book.ID = result.InsertedID.(primitive.ObjectID)

	return c.JSON(http.StatusCreated, book)
}

func (h *BookHandler) PutBook(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}
	var reqBook Book
	if err := c.Bind(&reqBook); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	update := bson.M{
		"$set": bson.M{
			"title":  reqBook.Title,
			"author": reqBook.Author,
		},
	}
	result, err := h.collection.UpdateOne(
		context.Background(),
		bson.M{"_id": id},
		update,
	)

	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if result.MatchedCount == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "book not found")
	}
	reqBook.ID = id
	return c.JSON(http.StatusOK, reqBook)

}

func (h *BookHandler) DeleteBook(c echo.Context) error {

	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	result, err := h.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "not deleted")
	}
	if result.DeletedCount == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "book not found")
	}
	return c.NoContent(http.StatusNoContent)

}
