/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"yak-cli/internal/model"
	"yak-cli/internal/storage"
	"yak-cli/utils"

	"github.com/spf13/cobra"
)

type addFlags struct {
	desc        string
	priority    int
	dueDate     string
	interactive bool
}

var cmdFlags addFlags

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "A brief description of your command",
	Args:  cobra.RangeArgs(0, 1),
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var item model.Todo

		title := ""
		hasTitle := len(args) > 0
		if hasTitle {
			title = strings.TrimSpace(args[0])
		}

		item.Title = title

		if title == "" || cmdFlags.interactive {
			err := interactionModeForAdd(cmd, &item)
			if err != nil {
				return err
			}
			return nil
		}

		item.Desc = strings.TrimSpace(cmdFlags.desc)

		err := getPriority(cmd, &item)
		if err != nil {
			return err
		}

		err = getDueDate(cmd, &item)
		if err != nil {
			return err
		}

		err = storage.AddTodo(item)
		if err != nil {
			return err
		}
		return nil
	},
}

func interactionModeForAdd(cmd *cobra.Command, item *model.Todo) error {
	isTitleEmpty := item.Title == ""
	if isTitleEmpty {
		userInput, err := utils.PromptUntilInput("Title for the todo: ", "Title for the ToDo can't be empty")
		if err != nil {
			return err
		}
		item.Title = userInput
	}

	userDesc := cmdFlags.desc
	var err error
	if !cmd.Flags().Changed("detail") {
		userDesc, err = utils.GetUserInput("Any Description/Details for the todo (optional): ")
		if err != nil {
			return err
		}
	}
	item.Desc = strings.TrimSpace(userDesc)

	err = getPriority(cmd, item)
	if err != nil {
		return err
	}

	err = getDueDate(cmd, item)
	if err != nil {
		return err
	}
	return nil
}

func getDueDate(cmd *cobra.Command, item *model.Todo) error {
	dueDate := cmdFlags.dueDate
	var err error

	if !cmd.Flags().Changed("due") {
		dueDate, err = utils.GetUserInput("Deadline for todo item (optional): ")
		if err != nil {
			return err
		}
		if dueDate == "" {
			fmt.Println("Skipping Due Date...")
			return nil
		}
	}

	for {
		formattedDate, err := validateAndParseDueDate(dueDate)
		if err == nil {
			item.DueDate = formattedDate
			return nil
		}

		fmt.Println("Warning: " + err.Error())
		dueDate, err = utils.GetUserInput("Deadline for todo item (optional): ")
		if err != nil {
			return err
		}
		if dueDate == "" {
			fmt.Println("Skipping Due Date...")
			return nil
		}
	}
}

func validateAndParseDueDate(dueDate string) (string, error) {
	itemDue, err := utils.FormatParsedDate(dueDate)
	if err != nil {
		return "", err
	}
	return itemDue, nil
}

func getPriority(cmd *cobra.Command, item *model.Todo) error {
	itemPriority := strconv.Itoa(cmdFlags.priority)
	var err error

	if !cmd.Flags().Changed("priority") {
		itemPriority, err = utils.GetUserInput("Add priority for the todo (0-2)...by default it's none: ")
		if err != nil {
			return err
		}
		if itemPriority == "" {
			fmt.Println("Skipping Priority...")
			return nil
		}
	}

	for {
	    parsed, err := validatePriority(itemPriority)
	    if err == nil {
	        item.Priority = &parsed
	        return nil
	    }

		fmt.Println("Warning: " + err.Error())
	    itemPriority, err = utils.GetUserInput("Priority for item (0-2, or Enter to skip): ")
		if err != nil {
			return err
		}
	    if itemPriority == "" {
			fmt.Println("Skipping Priority...")
	        return nil
	    }
	}
}

func validatePriority(priority string) (int, error) {
	parsedPriority, err := strconv.Atoi(priority)
	if err != nil {
		return 0, fmt.Errorf("Invalid priority: %w", err)
	}
	if parsedPriority < 0 || parsedPriority > 2 {
		return 0, fmt.Errorf("Invalid priority, please select priority b/w 0 and 2 (2 being lowest)")
	}
	return parsedPriority, nil
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolVarP(&cmdFlags.interactive, "interactive", "i", false, "Add todo details in interactive mode")
	addCmd.Flags().StringVarP(&cmdFlags.desc, "detail", "d", "", "description/details for the todo item")
	addCmd.Flags().IntVarP(&cmdFlags.priority, "priority", "p", 0, "priority for the todo item")
	addCmd.Flags().StringVarP(&cmdFlags.dueDate, "due", "t", "", "due date for the todo item")
}
