package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/mock"
)

type EchoContext interface {
	echo.Context
	JSON(int, interface{}) error
	Param(string) string
}

type echoContextMock struct {
	mock.Mock
	echo.Context
}

func (m *echoContextMock) JSON(code int, i interface{}) error {
	args := m.Called(code, i)
	return args.Error(0)
}

func (m *echoContextMock) Param(key string) string {
	args := m.Called(key)
	return args.String(0)
}
