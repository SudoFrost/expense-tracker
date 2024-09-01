package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/sudofrost/expense-tracker/internal/logic"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new expense",
	Run: func(cmd *cobra.Command, args []string) {
		description, _ := cmd.Flags().GetString("description")
		amount, _ := cmd.Flags().GetFloat64("amount")

		id := logic.AddExpense(description, amount)

		fmt.Printf("Expense added successfully. (ID: %v)\n", id)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringP("description", "d", "", "Description of the expense")
	addCmd.Flags().Float64P("amount", "a", 0, "Amount of the expense")

	addCmd.MarkFlagRequired("description")
	addCmd.MarkFlagRequired("amount")
}
