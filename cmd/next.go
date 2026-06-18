/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// nextCmd represents the next command
var nextCmd = &cobra.Command{
	Use:   "next",
	Short: "Show the next todo item",
	Long: `Show the next upcoming todo item from your list.

Examples:
  yak next`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("next called")
	},
}

func init() {
	rootCmd.AddCommand(nextCmd)
}
