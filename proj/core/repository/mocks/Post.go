package mocks

import (
	"github.com/stretchr/testify/mock"
	"go-training/proj/core/model"
)

type PostRepositoryMock struct {
	mock.Mock
}

func CreatePostRepositoryMock() *PostRepositoryMock {
	return &PostRepositoryMock{}
}

func (m *PostRepositoryMock) GetPostList(categoryId int32) (*[]model.Post, error) {
	args := m.Called(categoryId)

	return args.Get(0).(*[]model.Post), args.Error(1)
}

func (m *PostRepositoryMock) GetPostById(id int32) (*model.Post, error) {
	args := m.Called(id)

	return args.Get(0).(*model.Post), args.Error(1)
}

func (m *PostRepositoryMock) DeletePost(post *model.Post) error {
	args := m.Called(post)

	return args.Error(0)
}

func (m *PostRepositoryMock) UpdatePost(post *model.Post) error {
	args := m.Called(post)

	return args.Error(0)
}

func (m *PostRepositoryMock) AddPost(post *model.Post, userId int32) error {
	args := m.Called(post, userId)

	return args.Error(0)
}
