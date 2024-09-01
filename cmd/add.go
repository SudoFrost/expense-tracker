package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Adds a new expense",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("#TODO: Add a new expense")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().StringP("description", "d", "", "Description of the expense")
	addCmd.Flags().Float64P("amount", "a", 0, "Amount of the expense")
}
