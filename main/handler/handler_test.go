package handler_test

import (
	"awesomeProject/main/handler"
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	_ "go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"net/http/httptest"
	"testing"
)

type MockBookStore struct{}

func (m *MockBookStore) InsertOne(book *handler.Book) (primitive.ObjectID, error) {
	// Mock implementation to return a new ObjectID
	return primitive.NewObjectID(), nil
}

func (m *MockBookStore) GetAllBooks() ([]handler.Book, error) {
	// Mock implementation to return a list of books
	return []handler.Book{
		{
			ID:     primitive.NewObjectID(),
			Title:  "Book1",
			Author: "Author1",
		},
		{
			ID:     primitive.NewObjectID(),
			Title:  "Book2",
			Author: "Author2",
		},
	}, nil
}

func TestGetBooks(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()

	// Create a new book handler with the mock book store
	bookHandler := handler.NewBookHandler(&MockBookStore{})

	// Register the GetBooks endpoint
	e.GET("/books", bookHandler.GetBooks)

	// Create a new test request
	req := httptest.NewRequest(http.MethodGet, "/books", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Execute the handler
	if assert.NoError(t, bookHandler.GetBooks(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)

		var books []handler.Book
		err := json.Unmarshal(rec.Body.Bytes(), &books)
		assert.NoError(t, err)
		assert.Len(t, books, 2)
	}
}

func TestPostBook(t *testing.T) {
	// Create a new Echo instance
	e := echo.New()

	// Create a new book handler with the mock book store
	bookHandler := handler.NewBookHandler(&MockBookStore{})

	// Register the PostBook endpoint
	e.POST("/books", bookHandler.PostBook)

	// Create a new test request
	book := handler.Book{
		Title:  "New Book",
		Author: "New Author",
	}
	body, _ := json.Marshal(book)
	req := httptest.NewRequest(http.MethodPost, "/books", bytes.NewReader(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Execute the handler
	if assert.NoError(t, bookHandler.PostBook(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)

		var createdBook handler.Book
		err := json.Unmarshal(rec.Body.Bytes(), &createdBook)
		assert.NoError(t, err)
		assert.Equal(t, book.Title, createdBook.Title)
		assert.Equal(t, book.Author, createdBook.Author)
	}
}
