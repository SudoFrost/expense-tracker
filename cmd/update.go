/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/sudofrost/expense-tracker/internal/logic"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update [id]",
	Short: "Update an expense",
	Long:  `Update an expense`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Fprintln(os.Stderr, "ID is required")
			return
		}
		id, err := strconv.ParseInt(args[0], 10, 64)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Invalid ID")
			return
		}

		expense, err := logic.GetExpense(id)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			return
		}

		var description = expense.Description
		var amount = expense.Amount

		if cmd.Flags().Changed("description") {
			description, _ = cmd.Flags().GetString("description")
		}

		if cmd.Flags().Changed("amount") {
			amount, _ = cmd.Flags().GetFloat64("amount")
		}

		err = logic.UpdateExpense(id, description, amount)
		if err != nil {
			fmt.Fprintln(os.Stderr, err.Error())
			return
		}

		fmt.Printf("Expense updated successfully. (ID: %v)\n", id)
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringP("description", "d", "", "Description of the expense")
	updateCmd.Flags().Int64P("amount", "a", 0, "Amount of the expense")

	updateCmd.MarkFlagsOneRequired("description", "amount")
}
