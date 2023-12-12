package repository

import (
	"gorm.io/gorm"
	"shrt-server/types/entity"
)

type shrtRepository struct {
	db *gorm.DB
}

func NewShrtRepository(db *gorm.DB) shrtRepository {
	return shrtRepository{db: db}
}

func (r shrtRepository) FindBySlug(slug string) (*entity.Shrt, error) {
	var shrt *entity.Shrt
	err := r.db.Where("slug = ?", slug).First(&shrt).Error
	if err != nil {
		return nil, err
	}

	return shrt, nil
}

func (r shrtRepository) FindByLongURL(longURL string) (*entity.Shrt, error) {
	var shrt *entity.Shrt
	err := r.db.Where("long_url = ?", longURL).First(&shrt).Error
	if err != nil {
		return nil, err
	}

	return shrt, nil
}

func (r shrtRepository) UpdateVisit(id uint) error {
	return r.db.Model(&entity.Shrt{}).Where("id = ?", int(id)).Update("visit", gorm.Expr("visit + ?", 1)).Error
}

func (r shrtRepository) Create(shrt *entity.Shrt) error {
	return r.db.Create(shrt).Error
}
