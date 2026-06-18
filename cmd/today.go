/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type todayArguments struct {
	lastN   int
	showAll bool
}

var todayArgs todayArguments

// todayCmd represents the today command
var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "List today's todos",
	Long: `List all todo items due today.

Use -a or --all to show all todos for today, or -n to limit results.

Examples:
  yak today
  yak today -a
  yak today -n 5`,
	Run: func(cmd *cobra.Command, args []string) {
		if todayArgs.showAll {
			fmt.Println("all todos for today are following: ")
		} else if todayArgs.lastN > 0 {
			fmt.Printf("listing last %d todos for today\n", todayArgs.lastN)
		} else {
			fmt.Println("3 for today are following: ")
		}
	},
}

func init() {
	rootCmd.AddCommand(todayCmd)
	todayCmd.Flags().IntVarP(&todayArgs.lastN, "last-N", "n", -1, "List last n todos for today")
	todayCmd.Flags().BoolVarP(&todayArgs.showAll, "all", "a", false, "List all todos for today")
}
