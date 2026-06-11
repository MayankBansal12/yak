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

type todo struct {
	title string
	desc string
	priority int
	dueDate string
}

var item todo

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Args: cobra.ExactArgs(1),
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		item.title = strings.TrimSpace(args[0])
		if item.title == "" {
			return fmt.Errorf("Invalid argument for the command %s", args[0])
		}

		item.desc = strings.TrimSpace(item.desc)
		if item.desc != "" {
			fmt.Printf("Description Added\n");
		}

		if item.priority != -1 {
			if item.priority < 0 || item.priority > 2 {
			    return fmt.Errorf("invalid priority %d: must be 0, 1, or 2", item.priority)
			}
			fmt.Printf("Priority Added\n");
		}

		itemDue, err := utils.FormatParsedDate(item.dueDate)
		if err != nil {
			return fmt.Errorf("%w", err)
		}

		item.dueDate = itemDue

		// ToDo: save the todo in the local user data

		fmt.Printf("Added a new ToDo item: %v\n", item)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().StringVarP(&item.desc, "detail", "d", "", "description/details for the todo item")
	addCmd.Flags().IntVarP(&item.priority, "priority", "p", -1, "priority for the todo item")
	addCmd.Flags().StringVarP(&item.dueDate, "due", "t", "", "due date for the todo item")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
