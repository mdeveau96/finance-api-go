package models

import (
	"time"
)

type Budget struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	UserID     uint      `json:"user_id"`
	CategoryID uint      `json:"category_id"`
	Amount     float64   `json:"amount" gorm:"not null"`
	Period     string    `json:"period" gorm:"not null"` // monthly, weekly
	StartDate  time.Time `json:"start_date"`
	EndDate    time.Time `json:"end_date"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`

	User     User     `json:"user,omitempty"`
	Category Category `json:"category,omitempty"`
}
