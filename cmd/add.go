/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strings"
	"strconv"

	"yak-cli/utils"
	"github.com/spf13/cobra"
)

type todo struct {
	title string
	desc string
	priority int
	dueDate string
}

var interactive bool
var item todo

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Args: cobra.RangeArgs(0, 1),
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		hasTitle := len(args) > 0
		if hasTitle && strings.TrimSpace(args[0]) != "" {
			item.title = strings.TrimSpace(args[0])
		} else {
			interactionModeForAdd()
			return nil
		}

		if interactive {
			interactionModeForAdd()
			return nil
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

		if(item.dueDate != ""){
			itemDue, err := utils.FormatParsedDate(item.dueDate)
			if err != nil {
				return fmt.Errorf("%w", err)
			}
			item.dueDate = itemDue
		}

		// ToDo: save the todo in the local user data

		fmt.Printf("Added a new ToDo item: %v\n", item)
		return nil
	},
}

func interactionModeForAdd() {
	if item.title == "" {
		userInput := utils.GetUserInput("Title for the todo: ")
		if userInput == "" {
			// should have retry instead of throwing error
			fmt.Println("ERROR: Title for the ToDo can't be empty");
			return;
		}
		item.title = userInput
	}

	if item.desc == "" {
		userInput := utils.GetUserInput("Any Description/Details for the todo (optional): ")
		item.desc = userInput
	}

	if item.priority == -1 {
		userInput := utils.GetUserInput("Want to add priority for the todo? By default it's none: ")
		priority, err := strconv.Atoi(userInput)

		if userInput != "" && (err != nil || priority < 0 || priority > 2) {
			fmt.Println("ERROR: Invalid priority, please select priority b/w 0 and 2 (2 being lowest)")
			return;
		}
		item.priority = priority
	}

	if item.dueDate == "" {
		userInput := utils.GetUserInput("Deadline for todo item (optional): ")
		if userInput != "" {
			dueDate, err := utils.FormatParsedDate(userInput)
			if err != nil {
				fmt.Printf("ERROR: Invalid due date format: %s, please use DD-MM or DD-MM-YY\n", userInput)
				return;
			}
			item.dueDate = dueDate
		}
	}

	fmt.Printf("Interactive mode for adding todo item: %v\n", item)
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolVarP(&interactive, "interactive", "i", false, "Add todo details in interactive mode")
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
