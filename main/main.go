// create package main
package main

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
	{ID: 1, Title: "title1", Author: "ismail"},
	{ID: 1, Title: "title1", Author: "ismail"},
	{ID: 1, Title: "title1", Author: "ismail"},
}

func main() {

	e := echo.New()
	e.GET("/books", func(c echo.Context) error {
		return c.JSON(http.StatusOK, books)
	})

	e.POST("/books", func(c echo.Context) error {
		b := new(Book)
		if err := c.Bind(b); err != nil {
			return err
		}
		b.ID = len(books) + 1
		books = append(books, *b)
		return c.JSON(http.StatusCreated, books)
	})
	e.GET("books/:id", func(c echo.Context) error {
		idInteger, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err)
		}
		for _, item := range books {
			if item.ID == idInteger {
				return c.JSON(http.StatusOK, item)
			}
		}
		return c.JSON(http.StatusNotFound, "not found")
	})

	e.PUT("books/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "invalid book id")
		}

		for i := range books {
			if books[i].ID == id {
				b := new(Book)
				if err := c.Bind(b); err != nil {
					return err
				}
				b.ID = id
				books[i] = *b
				return c.JSON(http.StatusOK, b)
			}

		}
		return c.JSON(http.StatusNotFound, "book not found")
	})

	e.DELETE("books/:id", func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "conversion error")
		}
		for i, b := range books {
			if b.ID == id {
				//books := append(books[:i], books[i+1:]...)
				books = append(books[:i], books[i+1:]...)

				return c.NoContent(http.StatusNoContent)
			}
		}
		return echo.NewHTTPError(http.StatusNotFound, "Book not found")
	})

	e.Logger.Fatal(e.Start(":8080"))

}
