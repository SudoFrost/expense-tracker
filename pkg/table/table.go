package table

import (
	"fmt"
	"strings"
)

func Print(columns []string, rows []map[string]string, separator string) {
	columnsLengths := make([]int, len(columns))

	for i, column := range columns {
		columnsLengths[i] = len(column)
	}

	for _, row := range rows {
		for i, column := range columns {
			if len(row[column]) > columnsLengths[i] {
				columnsLengths[i] = len(row[column])
			}
		}
	}

	for i, column := range columns {
		fmt.Printf("%s%s", strings.Repeat(" ", columnsLengths[i]-len(column)), column)
		if i < len(columns)-1 {
			fmt.Print(separator)
		}
	}
	fmt.Println()

	for _, row := range rows {
		for i, column := range columns {
			fmt.Printf("%s%s", strings.Repeat(" ", columnsLengths[i]-len(row[column])), row[column])
			if i < len(columns)-1 {
				fmt.Print(separator)
			}
		}
		fmt.Println()
	}
}
