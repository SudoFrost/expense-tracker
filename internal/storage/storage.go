package storage

import (
	"encoding/json"
	"os"
	"path"
	"time"

	"github.com/sudofrost/expense-tracker/internal/models"
)

const (
	// Environment variable for get the storage file path
	FilePathEnv = "EXPENSE_TRACKER_STORAGE_FILE_PATH"
	// Default storage file path when no environment variable is set in home directory
	DefaultFilePath = "expense-tracker/expenses.json"
)

var filepath string

type ExpenseJson struct {
	CreatedAt int64
	UpdatedAt int64
	*models.Expense
}

func Save(expenses []models.Expense) error {
	expensesJson := make([]ExpenseJson, 0, len(expenses))

	for _, expense := range expenses {
		expensesJson = append(expensesJson, ExpenseJson{
			CreatedAt: expense.CreatedAt.Unix(),
			UpdatedAt: expense.UpdatedAt.Unix(),
			Expense:   &expense,
		})
	}

	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	encoder := json.NewEncoder(file)
	return encoder.Encode(expensesJson)
}

func Load() ([]models.Expense, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var expensesJson []ExpenseJson
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&expensesJson)
	if err != nil && err.Error() != "EOF" {
		return nil, err
	}
	expenses := make([]models.Expense, 0, len(expensesJson))
	for _, expenseJson := range expensesJson {
		expenseJson.Expense.CreatedAt = time.Unix(expenseJson.CreatedAt, 0)
		expenseJson.Expense.UpdatedAt = time.Unix(expenseJson.UpdatedAt, 0)
		expenses = append(expenses, *expenseJson.Expense)
	}

	return expenses, nil
}

func init() {
	filepath = os.Getenv(FilePathEnv)
	if filepath == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}
		filepath = home + "/" + DefaultFilePath
	}

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		dir := path.Dir(filepath)
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			os.MkdirAll(dir, 0755)
		}
		path.Clean(filepath)
		file, err := os.Create(filepath)
		if err != nil {
			panic(err)
		}
		defer file.Close()
	}
}
