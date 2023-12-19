package entity

import "time"

type Shrt struct {
	ID        uint      `gorm:"primarykey,autoIncrement"`
	LongURL   string    `gorm:"type:varchar(255),not null"`
	Slug      string    `gorm:"type:varchar(255),not null;unique"`
	Visit     int       `gorm:"visit,default:0,not null"`
	CreatedAt time.Time `gorm:"column:createdAt"`
	UpdatedAt time.Time `gorm:"column:updatedAt"`
}

type Tabler interface {
	TableName() string
}

func (Shrt) TableName() string {
	return "Shrts"
}
