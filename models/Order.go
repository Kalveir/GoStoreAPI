package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Order struct {
	ID uuid.UUID `json:"id" gorm:"primaryKey;not null;unique"`
	// ID         uint        `json:"id" gorm:"PrimaryKey;autoIncrement"`
	UserID     uint        `json:"user_id" gorm:"not null" validate:"required"`
	TotalPrice uint        `json:"total_price" gorm:"not null"`
	TotalItem  uint        `json:"total_item" gorm:"not null"`
	Status     string      `json:"status" gorm:"type:default:'pending'"`
	OrderItems []OrderItem `json:"order_items,omitempty" gorm:"foreignKey:OrderID;constraint:OnDelete:CASCADE"`
	CreatedAt  *time.Time   `json:"created_at" gorm:"autoCreateTime"`
}

type OrderItem struct {
	ID      *uint     		`json:"id" gorm:"PrimaryKey;autoIncrement"`
	OrderID uuid.UUID   	`json:"order_id" gorm:"not null;references:ID"`
	// OrderID   uint      `json:"order_id" gorm:"not null"`
	ProductID uint      	`json:"product_id" gorm:"not null"`
	Product   Product   	`json:"product,omitempty"`
	Quantity  uint      	`json:"quantity" gorm:"not null"`
	TotalCost uint      	`json:"total_cost" gorm:"not null"`
	CreatedAt time.Time 	`json:"created_at" gorm:"autoCreateTime"`
}

func (o *Order) BeforeCreate(tx *gorm.DB) error {
	o.ID, _ = uuid.NewV7()
	return nil
}
