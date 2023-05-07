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

func (h *BookHandler) GetBookById(c echo.Context) error {
	id := c.Param("id")
	book, err := h.bookRepository.GetBookById(context.Background(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid id")
	}
	return c.JSON(http.StatusOK, book)

}

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

func (h *BookHandler) PutBook(c echo.Context) error {
	id := c.Param("id")

	req := new(models.Book)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "bad request")
	}
	err := h.bookRepository.UpdateBook(context.Background(), id, *req)

	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "")
}

func (h *BookHandler) DeleteBook(c echo.Context) error {
	id := c.Param("id")
	err := h.bookRepository.DeleteBook(context.Background(), id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "not deleted")
	}
	return c.JSON(http.StatusOK, "")
}
