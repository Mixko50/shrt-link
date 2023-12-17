package entity

import "time"

type Shrt struct {
	ID        uint      `gorm:"primarykey"`
	LongURL   string    `gorm:"type:varchar(255)"`
	Slug      string    `gorm:"type:varchar(255)"`
	Visit     int       `gorm:"visit"`
	CreatedAt time.Time `gorm:"column:createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt"`
}

type Tabler interface {
	TableName() string
}

func (Shrt) TableName() string {
	return "Shrts"
}
