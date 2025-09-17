package models

import (
	"time"
)

type Transaction struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
    UserID      uint      `json:"user_id"`
    AccountID   uint      `json:"account_id"`
    CategoryID  uint      `json:"category_id"`
    Amount      float64   `json:"amount" gorm:"not null"`
    Type        string    `json:"type" gorm:"not null"` // income, expense, transfer
    Description string    `json:"description"`
    Date        time.Time `json:"date" gorm:"not null"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
    
    User     User     `json:"user,omitempty"`
    Account  Account  `json:"account,omitempty"`
    Category Category `json:"category,omitempty"`
}
