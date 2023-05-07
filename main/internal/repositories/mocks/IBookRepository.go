// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	models "awesomeProject/main/internal/models"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// IBookRepository is an autogenerated mock type for the IBookRepository type
type IBookRepository struct {
	mock.Mock
}

// CreateBooks provides a mock function with given fields: ctx, book
func (_m *IBookRepository) CreateBooks(ctx context.Context, book models.Book) error {
	ret := _m.Called(ctx, book)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, models.Book) error); ok {
		r0 = rf(ctx, book)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetBookById provides a mock function with given fields: background, id
func (_m *IBookRepository) GetBookById(background context.Context, id string) (models.Book, error) {
	ret := _m.Called(background, id)

	var r0 models.Book
	if rf, ok := ret.Get(0).(func(context.Context, string) models.Book); ok {
		r0 = rf(background, id)
	} else {
		r0 = ret.Get(0).(models.Book)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(background, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetBooks provides a mock function with given fields: ctx
func (_m *IBookRepository) GetBooks(ctx context.Context) ([]models.Book, error) {
	ret := _m.Called(ctx)

	var r0 []models.Book
	if rf, ok := ret.Get(0).(func(context.Context) []models.Book); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]models.Book)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}