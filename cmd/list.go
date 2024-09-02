/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/sudofrost/expense-tracker/internal/logic"
	"github.com/sudofrost/expense-tracker/pkg/table"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all expenses",
	Long:  `Lists all expenses`,
	Run: func(cmd *cobra.Command, args []string) {
		expenses := logic.GetExpenses()
		if len(expenses) == 0 {
			fmt.Println("No expenses found")
			return
		}
		columns := []string{"ID", "Description", "Amount", "Created At", "Updated At"}
		data := make([]map[string]string, 0, len(expenses))
		for _, expense := range expenses {
			data = append(data, map[string]string{
				"ID":          fmt.Sprintf("%.0f", expense.ID),
				"Description": expense.Description,
				"Amount":      fmt.Sprintf("%.2f", expense.Amount),
				"Created At":  expense.CreatedAt.Format("2006-01-02 15:04:05"),
				"Updated At":  expense.UpdatedAt.Format("2006-01-02 15:04:05"),
			})
		}
		table.Print(columns, data, " | ")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
