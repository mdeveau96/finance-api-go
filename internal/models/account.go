package models

import (
	"time"
)

type Account struct {
	ID uint `json:"id" gorm:"primaryKey"`
	UserID uint `json:"user_id" gorm:"not null;index"`
	Name string `json:"name" gorm:"not null"`
	Type string `json:"type" gorm:"not null"`
	Balance float64 `json:"balance" gorm:"not null"`
	Currency string `json:"currency" gorm:"not null;default:'USD'"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	
	User User `json:"user,omitempty" gorm:"foreignKey:UserID"`
	Transactions []Transaction `json:"transactions,omitempty" gorm:"foreignKey:AccountID"`
}
