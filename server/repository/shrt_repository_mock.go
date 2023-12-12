package repository

import (
	"github.com/stretchr/testify/mock"
	"shrt-server/types/entity"
)

type shrtRepositoryMock struct {
	mock.Mock
}

func NewShrtRepositoryMock() *shrtRepositoryMock {
	return &shrtRepositoryMock{}
}

func (m *shrtRepositoryMock) FindBySlug(slug string) (*entity.Shrt, error) {
	args := m.Called(slug)
	return args.Get(0).(*entity.Shrt), args.Error(1)
}

func (m *shrtRepositoryMock) FindByLongURL(longURL string) (*entity.Shrt, error) {
	args := m.Called(longURL)
	return args.Get(0).(*entity.Shrt), args.Error(1)
}

func (m *shrtRepositoryMock) UpdateVisit(id uint) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *shrtRepositoryMock) Create(shrt *entity.Shrt) error {
	args := m.Called(shrt)
	return args.Error(0)
}
