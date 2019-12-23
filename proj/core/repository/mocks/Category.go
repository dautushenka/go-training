package mocks

import (
	"github.com/stretchr/testify/mock"
	"go-training/proj/core/model"
)

type CategoryRepositoryMock struct {
	mock.Mock
}

func CreateCategoryRepositoryMock() *CategoryRepositoryMock {
	return &CategoryRepositoryMock{}
}

func (m *CategoryRepositoryMock) GetCategoryList() (*[]model.Category, error) {
	args := m.Called()

	return args.Get(0).(*[]model.Category), args.Error(1)
}
