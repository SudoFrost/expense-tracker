/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/sudofrost/expense-tracker/internal/logic"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete an expense",
	Long:  `Delete an expense`,
	Run: func(cmd *cobra.Command, args []string) {
		id, _ := cmd.Flags().GetInt64("id")
		err := logic.DeleteExpense(id)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		fmt.Printf("Expense deleted successfully. (ID: %v)\n", id)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	deleteCmd.Flags().Int64P("id", "i", 0, "ID of the expense")

	deleteCmd.MarkFlagRequired("id")
}
