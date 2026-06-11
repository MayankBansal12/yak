/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"

	"yak-cli/utils"
	"github.com/spf13/cobra"
)

type arguments struct {
	lastN int
	date string
	isDetail bool
}

var commandArgs arguments

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Aliases: []string{"view", "see", "get", "find"},
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			idArg := strings.TrimSpace(strings.TrimLeft(args[0], "#"))
			if idArg == "" {
				return fmt.Errorf("todo id is required for list command")
			}
			fmt.Println("list todo item with id:", idArg)
			return nil
		}

		if commandArgs.lastN > 0 {
			fmt.Printf("listing last %d todos\n", commandArgs.lastN)
		} else if commandArgs.date != "" {
			parsedDate, err := utils.FormatParsedDate(commandArgs.date)
			if err != nil {
				return fmt.Errorf("%w", err)
			}
			fmt.Printf("listing todos assigned to %v\n", parsedDate)
		}

		// fetch the todos stored in local store and present them to user
		if commandArgs.isDetail {
			fmt.Println("here's the todo in detailed format")
		}else {
			fmt.Println("here's the list of all the todos")
		}

		return nil
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().IntVarP(&commandArgs.lastN, "last-N", "n", -1, "List last n todos including completed ones")
	listCmd.Flags().StringVarP(&commandArgs.date, "datetime", "t", "", "List all todos assigned to a specific date")
	listCmd.Flags().BoolVarP(&commandArgs.isDetail, "detail", "d", false, "List todos in detailed format (title, description, priority, deadline, created at)")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
