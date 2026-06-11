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

var updateArgs todo

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update [todo-id]",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.RangeArgs(1,2),
	RunE: func(cmd *cobra.Command, args []string) error {
		idArg := strings.TrimSpace(strings.TrimLeft(args[0], "#"))
		if idArg == "" {
			return fmt.Errorf("ToDo item ID is required for the update command")
		}

		hasTitle := len(args) > 1
		if !hasTitle && cmd.Flags().NFlag() == 0 {
		    return fmt.Errorf("either a title or at least one flag is required")
		}

		if hasTitle && strings.TrimSpace(args[1]) != "" {
			updateArgs.title = strings.TrimSpace(args[1])
			fmt.Printf("Updated the todo item title to %s\n", updateArgs.title)
		}

		if updateArgs.desc != "" {
			fmt.Println("Updated the todo item description")
		}

		if updateArgs.priority != -1 {
			if updateArgs.priority < 0 || updateArgs.priority > 2 {
			    return fmt.Errorf("invalid priority %d: must be 0, 1, or 2", updateArgs.priority)
			}
			fmt.Println("Updated the todo item priority")
		}

		if updateArgs.dueDate != "" {
			itemDue, err := utils.FormatParsedDate(updateArgs.dueDate)
			if err != nil {
				return fmt.Errorf("%w", err)
			}
			updateArgs.dueDate = itemDue
			fmt.Printf("Updated the todo item due date to %s\n", itemDue)
		}

		fmt.Printf("Updated the todo item %s : %v\n", idArg, updateArgs)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVarP(&updateArgs.desc, "detail", "d", "", "description/details for the todo item")
	updateCmd.Flags().IntVarP(&updateArgs.priority, "priority", "p", -1, "priority for the todo item")
	updateCmd.Flags().StringVarP(&updateArgs.dueDate, "due", "t", "", "due date for the todo item")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// updateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// updateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
