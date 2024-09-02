/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"github.com/sudofrost/expense-tracker/cmd"
	"github.com/sudofrost/expense-tracker/internal/logic"
)

func main() {
	cmd.Execute()
	defer logic.Save()
}
