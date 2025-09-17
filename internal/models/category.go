package models

import (
	"time"
)

type Category struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Name string `json:"name" gorm:"not null;unique"`
	Type string `json:"type" gorm:"not null"` // e.g., "income" or "expense"
	Color string `json:"color"`
	CreatedAt time.Time `json:"created_at"`

	Transactions []Transaction `json:"transactions,omitempty" gorm:"foreignKey:CategoryID"`
}
