package repository

import (
	"shrt-server/types/entity"
)

type ShrtRepository interface {
	FindBySlug(slug string) (*entity.Shrt, error)
	FindByLongURL(longURL string) (*entity.Shrt, error)
	UpdateVisit(id uint) error
	Create(shrt *entity.Shrt) error
}
