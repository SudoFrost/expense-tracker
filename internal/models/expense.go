package models

import (
	"time"
)

type Expense struct {
	ID          int64
	Description string
	Amount      float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewExpense(id int64, amount float64, description string) *Expense {
	return &Expense{
		ID:          id,
		Description: description,
		Amount:      amount,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
