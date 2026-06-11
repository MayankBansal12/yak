/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

type todayArguments struct {
	lastN int
	showAll bool
}

var todayArgs todayArguments

// todayCmd represents the today command
var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if todayArgs.showAll {
			fmt.Println("all todos for today are following: ")
		} else if todayArgs.lastN > 0 {
			fmt.Printf("listing last %d todos for today\n", todayArgs.lastN)
		}else {
			fmt.Println("3 for today are following: ")
		}
	},
}

func init() {
	rootCmd.AddCommand(todayCmd)
	todayCmd.Flags().IntVarP(&todayArgs.lastN, "last-N", "n", -1, "List last n todos for today")
	todayCmd.Flags().BoolVarP(&todayArgs.showAll, "all", "a", false, "List all todos for today")


	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// todayCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// todayCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
