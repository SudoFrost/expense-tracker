/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
	"github.com/sudofrost/expense-tracker/internal/logic"
	"github.com/sudofrost/expense-tracker/internal/models"
)

// summaryCmd represents the summary command
var summaryCmd = &cobra.Command{
	Use:   "summary",
	Short: "Summary of expenses",
	Long:  `Summary of expenses`,
	Run: func(cmd *cobra.Command, args []string) {
		var expenses []*models.Expense
		var year = time.Now().Year()
		var month, _ = cmd.Flags().GetInt64("month")

		if month < 0 || month > 12 {
			fmt.Fprintln(os.Stderr, "Invalid month")
			return
		}

		for _, expense := range logic.GetExpenses() {
			if expense.CreatedAt.Year() == year {
				if month == 0 || expense.CreatedAt.Month() == time.Month(month) {
					expenses = append(expenses, &expense)
				}
			}
		}

		var total float64
		for _, expense := range expenses {
			total += expense.Amount
		}
		if month == 0 {
			fmt.Printf("Total expenses: $%.2f\n", total)
		} else {
			fmt.Printf("Total expenses for %s: $%.2f\n", time.Month(month), total)
		}
	},
}

func init() {
	rootCmd.AddCommand(summaryCmd)

	summaryCmd.Flags().Int64P("month", "m", 0, "Month of the year")
}
