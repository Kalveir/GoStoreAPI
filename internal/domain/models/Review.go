package models

import (
	"time"
)

type Review struct {
	ID        uint      `json:"id" gorm:"primaryKey;autoIncrement"`
	UserID    uint      `json:"-" gorm:"not null"`
	User      User      `json:"User" gorm:"foreignKey:UserID"`
	ProductID uint      `json:"product_id" gorm:"not null"`
	Rating    int       `json:"rating" gorm:"not null;check:rating >= 1 AND rating <= 5"`
	Comment   string    `json:"comment" gorm:"type:text"`
	CreatedAt time.Time `json:"CreatedAt" gorm:"autoCreateTime"`
}
