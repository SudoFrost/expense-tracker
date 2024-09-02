package logic

import (
	"github.com/sudofrost/expense-tracker/internal/models"
	"github.com/sudofrost/expense-tracker/internal/storage"
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

func AddExpense(description string, amount float64) float64 {
	expense := models.NewExpense(NewID(), amount, description)
	tracker.Expenses = append(tracker.Expenses, expense)
	return expense.ID
}

func init() {
	tracker = models.NewTracker()

	expenses, err := storage.Load()
	if err != nil {
		panic(err)
	}
	for _, expense := range expenses {
		tracker.Expenses = append(tracker.Expenses, &expense)
	}
}

func Save() error {
	expenses := make([]models.Expense, 0, len(tracker.Expenses))
	for _, expense := range tracker.Expenses {
		expenses = append(expenses, *expense)
	}
	return storage.Save(expenses)
}
