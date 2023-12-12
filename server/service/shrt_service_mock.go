package service

import (
	"github.com/stretchr/testify/mock"
	"shrt-server/types"
)

type shrtServiceMock struct {
	mock.Mock
}

func NewShrtServiceMock() *shrtServiceMock {
	return &shrtServiceMock{}
}

func (m *shrtServiceMock) CreateShrtLink(body *types.CreateShortenLinkRequest) (*types.CreateShortenLinkResponse, error) {
	args := m.Called(body)
	return args.Get(0).(*types.CreateShortenLinkResponse), args.Error(1)
}

func (m *shrtServiceMock) GetOriginalURL(slug string) (*types.CreateShortenLinkResponse, error) {
	args := m.Called(slug)
	return args.Get(0).(*types.CreateShortenLinkResponse), args.Error(1)
}

func (m *shrtServiceMock) GetOriginalURLToRedirect(slug string) (string, error) {
	args := m.Called(slug)
	return args.String(0), args.Error(1)
}
