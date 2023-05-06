package repositories

import (
	"awesomeProject/main/internal/models"
	"context"
	"github.com/stretchr/testify/mock"
	context2 "golang.org/x/net/context"
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
