package handler

import (
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetBooks(t *testing.T) {
	e := echo.New() // yeni bir instance oluşturulur.
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	err := GetBooks(c)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "title1")
	}
}

func TestPostBooks(t *testing.T) {
	e := echo.New()
	book := Book{Title: "newtitle", Author: "new author"}

	bookJSON, _ := json.Marshal(book)
	req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewBuffer(bookJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	err := PostBook(c)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Contains(t, rec.Body.String(), "newtitle")
	}

}

func TestGetBook(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/books/1", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("books/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := GetBook(c)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "title1")
	}

}

func TestPutBook(t *testing.T) {
	e := echo.New()

	updatedBook := Book{
		Title:  "newTitletest",
		Author: "newauthortest",
	}
	updatedBookJSON, _ := json.Marshal(updatedBook)

	req := httptest.NewRequest(http.MethodPut, "/books/1", bytes.NewBuffer(updatedBookJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("books/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	err := PutBook(c)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Contains(t, rec.Body.String(), "newTitletest")
	}

}

func TestDeleteBook(t *testing.T) {
	e := echo.New()
	book := Book{
		ID:     100,
		Title:  "newtitle",
		Author: "ismailaltas",
	}
	books = append(books, book)

	req := httptest.NewRequest(http.MethodDelete, "/books/100", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetPath("books/:id")
	c.SetParamNames("id")
	c.SetParamValues("100")

	err := DeleteBook(c)

	if assert.NoError(t, err) {
		assert.Equal(t, http.StatusNoContent, rec.Code)
		reqGet := httptest.NewRequest(http.MethodGet, "/books/100", nil)
		recGet := httptest.NewRecorder()

		cGet := e.NewContext(reqGet, recGet)
		cGet.SetPath("books/:id")
		cGet.SetParamNames("id")
		cGet.SetParamValues("100")
		errGet := GetBook(cGet)
		if assert.NoError(t, errGet) {
			assert.Contains(t, recGet.Body.String(), "book not found")
		}
	}
}
