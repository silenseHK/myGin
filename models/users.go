package models

import "time"

type Users struct{
	Id int64
	Phone string
	Password string
	CreatedAt time.Time `gorm:"default:null"`
	DeletedAt time.Time `gorm:"default:null"`
	Balance float64 `gorm:"default:0"`
	Status uint `gorm:"default:1"`
}
