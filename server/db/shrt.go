package db

import "gorm.io/gorm"

type Shrts struct {
	gorm.Model
	LongURL string `json:"long_url" gorm:"type:varchar(255)"`
	Slug    string `json:"slug" gorm:"type:varchar(255)"`
	Visit   int    `json:"visit"`
}

type Tabler interface {
	TableName() string
}

func (Shrts) TableName() string {
	return "Shrts"
}
