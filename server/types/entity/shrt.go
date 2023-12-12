package entity

import "time"

type Shrt struct {
	ID        uint      `json:"id" gorm:"primarykey"`
	LongURL   string    `json:"long_url" gorm:"type:varchar(255)"`
	Slug      string    `json:"slug" gorm:"type:varchar(255)"`
	Visit     int       `json:"visit"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Tabler interface {
	TableName() string
}

func (Shrt) TableName() string {
	return "Shrts"
}
