package logic

import (
	"github.com/sudofrost/expense-tracker/internal/models"
)

var tracker *models.Tracker

func NewID() float64 {
	var max float64
	for _, expense := range tracker.Expenses {
		if expense.ID > max {
			max = expense.ID
		}
	}
	return max + 1
}

func init() {
	tracker = models.NewTracker()
}
