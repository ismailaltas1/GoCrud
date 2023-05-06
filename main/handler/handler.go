package handler

import (
	"awesomeProject/main/internal/models"
	"awesomeProject/main/internal/repositories"
	"context"
	"github.com/labstack/echo/v4"
	"net/http"
)

type BookHandler struct {
	bookRepository repositories.IBookRepository
}

func NewBookHandler(bookRepository repositories.IBookRepository) *BookHandler {
	return &BookHandler{
		bookRepository: bookRepository,
	}
}

func (h *BookHandler) GetBooks(c echo.Context) error {
	books, _ := h.bookRepository.GetBooks(context.Background())
	return c.JSON(http.StatusOK, books)
}

/*
func (h *BookHandler) GetBook(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}
	var book models.Book
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

*/

func (h *BookHandler) PostBook(c echo.Context) error {
	var book models.Book
	if err := c.Bind(&book); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}

	err := h.bookRepository.CreateBooks(context.Background(), book)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusCreated, book)
}

/*
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


*/
