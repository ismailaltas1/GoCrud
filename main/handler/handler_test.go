package handler

import (
	"awesomeProject/main/internal/models"
	"awesomeProject/main/internal/repositories"
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetBooks(t *testing.T) {
	e := echo.New()
	mockBookRepository := new(repositories.MockBookRepository)
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

func TestCreateBooks(t *testing.T) {
	e := echo.New()
	mockBookRepository := new(repositories.MockBookRepository)
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
