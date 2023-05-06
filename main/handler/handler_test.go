package handler

import (
	"awesomeProject/main/internal/models"
	"bytes"
	"context"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	context2 "golang.org/x/net/context"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockBookRepository struct {
	mock.Mock
}

func (m *MockBookRepository) CreateBooks(ctx context2.Context, book models.Book) (err error) {
	args := m.Called(ctx, book)
	return args.Error(0)
}

func (m *MockBookRepository) GetBooks(ctx context.Context) ([]models.Book, error) {
	args := m.Called(ctx)
	return args.Get(0).([]models.Book), args.Error(1)
}

func TestGetBooks(t *testing.T) {
	e := echo.New()
	mockBookRepository := new(MockBookRepository)
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
	mockBookRepository := new(MockBookRepository)
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
