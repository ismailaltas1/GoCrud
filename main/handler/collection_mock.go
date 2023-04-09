package handler

import (
	"context"

	"github.com/stretchr/testify/mock"
)

type BookStoreMock struct {
	mock.Mock
}

func (m *BookStoreMock) Find(ctx context.Context, filter interface{}) ([]Book, error) {
	args := m.Called(ctx, filter)
	return args.Get(0).([]Book), args.Error(1)
}

// Implement the other methods for the BookStore interface...
