package handler

import (
	"awesomeProject/main/internal/models"
	"awesomeProject/main/internal/repositories/mocks"
	"bytes"
	"context"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	mock "github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestBookHandler_GetBooks(t *testing.T) {
	e := echo.New()
	mockBookRepository := new(mocks.IBookRepository)
	mockBooks := []models.Book{
		{ID: "1", Title: "Book1", Author: "Author1"},
		{ID: "2", Title: "Book2", Author: "Author2"},
	}

	mockBookRepository.On("GetBooks", mock.Anything).Return(mockBooks, nil)
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	h := NewBookHandler(mockBookRepository)
	err := h.GetBooks(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `[{"_id":"1","title":"Book1","author":"Author1"},{"_id":"2","title":"Book2","author":"Author2"}]`, rec.Body.String())
	mockBookRepository.AssertExpectations(t)
}

func TestBookHandler_GetBookById(t *testing.T) {
	e := echo.New()
	mockBookRepository := new(mocks.IBookRepository)
	mockBooks := models.Book{ID: "1", Title: "Book1", Author: "Author1"}

	mockBookRepository.On("GetBookById", context.Background(), "1").Return(mockBooks, nil)
	req := httptest.NewRequest(http.MethodGet, "/books/1", nil)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	h := NewBookHandler(mockBookRepository)
	err := h.GetBookById(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	assert.JSONEq(t, `{"_id":"1","title":"Book1","author":"Author1"}`, rec.Body.String())
	mockBookRepository.AssertExpectations(t)

}

func TestBookHandler_PostBook(t *testing.T) {
	e := echo.New()
	mockBookRepository := new(mocks.IBookRepository)
	mockBook := models.Book{ID: "1", Title: "Book1", Author: "Author1"}

	mockBookRepository.On("CreateBooks", mock.Anything, mockBook).Return(nil)
	bookJson, _ := json.Marshal(mockBook)
	req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewReader(bookJson))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	h := NewBookHandler(mockBookRepository)
	err := h.PostBook(c)

	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
	assert.JSONEq(t, `{"_id":"1","title":"Book1","author":"Author1"}`, rec.Body.String())
	mockBookRepository.AssertExpectations(t)

}

func TestBookHandler_PutBook(t *testing.T) {
	e := echo.New()
	mockBookRepository := new(mocks.IBookRepository)
	mockBook := models.Book{ID: "1", Title: "Book1", Author: "Author1"}

	mockBookRepository.On("UpdateBook", mock.Anything, "1", mockBook).Return(nil)
	bookBytes, _ := json.Marshal(mockBook)
	req := httptest.NewRequest(http.MethodPut, "/books/1", bytes.NewReader(bookBytes))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/books/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	h := NewBookHandler(mockBookRepository)
	err := h.PutBook(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)

}

func TestBookHandler_DeleteBook(t *testing.T) {
	e := echo.New()
	mockBookRepository := new(mocks.IBookRepository)
	mockBookRepository.On("DeleteBook", mock.Anything, "1").Return(nil)

	req := httptest.NewRequest(http.MethodDelete, "/books/1", nil)
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetPath("/books/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")

	h := NewBookHandler(mockBookRepository)
	err := h.DeleteBook(c)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
}
