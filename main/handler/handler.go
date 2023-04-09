package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

type Book struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book = []Book{
	{ID: 1, Title: "title1", Author: "ismail1"},
	{ID: 2, Title: "title2", Author: "ismail2"},
	{ID: 3, Title: "title3", Author: "ismail3"},
}

func GetBooks(c echo.Context) error {
	return c.JSON(http.StatusOK, books)
}

func GetBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "convertion erorr")
	}
	for _, b := range books {
		if id == b.ID {
			return c.JSON(http.StatusOK, b)
		}
	}
	return c.JSON(http.StatusNotFound, "book not found")
}

func PostBook(c echo.Context) error {
	b := new(Book)
	if err := c.Bind(b); err != nil {
		return err
	}
	b.ID = len(books) + 1
	books = append(books, *b)
	return c.JSON(http.StatusCreated, books)
}

func PutBook(c echo.Context) error {
	paramId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "invalid book id")
	}

	for i, book := range books {

		if book.ID == paramId {
			b := &Book{ID: paramId}
			if err := c.Bind(b); err != nil {
				return err
			}
			books[i] = *b
			return c.JSON(http.StatusOK, books)
		}
	}
	return c.JSON(http.StatusNotFound, "book not found")

}

func DeleteBook(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "convertion error")
	}
	for i, book := range books {
		if book.ID == id {
			books = append(books[:i], books[i+1:]...) // ilk indexden i nin indexe kadar yeni bir slice oluşturulur
			//i+1 den array deki son elemana kadar yeni bir slice oluşturulur. ... ikinci arrayin birinci arrayden farklı bir arguman olduğunu beliritir.
			// append methoduyla iki slice birleştirilir.
			return c.JSON(http.StatusOK, books)
		}
	}
	return c.JSON(http.StatusNotFound, "element not found")

}
