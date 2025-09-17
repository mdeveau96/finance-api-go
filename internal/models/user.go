package models

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Email     string    `json:"email" gorm:"unique;not null"`
	Password  string    `json:"-" gorm:"not null"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Accounts     []Account     `json:"accounts,omitempty" gorm:"foreignKey:UserID"`
	Transactions []Transaction `json:"transactions,omitempty" gorm:"foreignKey:UserID"`
	Budgets      []Budget      `json:"budgets,omitempty" gorm:"foreignKey:UserID"`
}
