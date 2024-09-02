package logic

import (
	"fmt"
	"time"

	"github.com/sudofrost/expense-tracker/internal/models"
	"github.com/sudofrost/expense-tracker/internal/storage"
)

var tracker *models.Tracker

func NewID() int64 {
	var max int64
	for _, expense := range tracker.Expenses {
		if expense.ID > max {
			max = expense.ID
		}
	}
	return max + 1
}

func AddExpense(description string, amount float64) int64 {
	expense := models.NewExpense(NewID(), amount, description)
	tracker.Expenses = append(tracker.Expenses, expense)
	return expense.ID
}

func GetExpenses() []models.Expense {
	var expenses = make([]models.Expense, 0, len(tracker.Expenses))
	for _, expense := range tracker.Expenses {
		expenses = append(expenses, *expense)
	}
	return expenses
}

func GetExpense(id int64) (models.Expense, error) {
	for _, expense := range tracker.Expenses {
		if expense.ID == id {
			return *expense, nil
		}
	}
	return models.Expense{}, fmt.Errorf("expense not found")
}

func UpdateExpense(id int64, description string, amount float64) error {
	for _, expense := range tracker.Expenses {
		if expense.ID == id {
			expense.Description = description
			expense.Amount = amount
			expense.UpdatedAt = time.Now()
			return nil
		}
	}
	return fmt.Errorf("expense not found")
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
