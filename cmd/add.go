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

		hasTitle := len(args) > 0
		if hasTitle && strings.TrimSpace(args[0]) != "" {
			item.Title = strings.TrimSpace(args[0])
		} else {
			interactionModeForAdd(cmd, &item)
			return nil
		}

		if cmdFlags.interactive {
			interactionModeForAdd(cmd, &item)
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

		storage.AddTodo(item)
		return nil
	},
}

func interactionModeForAdd(cmd *cobra.Command, item *model.Todo) error {
	isTitleEmpty := item.Title == ""
	if isTitleEmpty {
		userInput := strings.TrimSpace(utils.GetUserInput("Title for the todo: "))

		// todo: should have retry instead of throwing error
		if userInput == "" {
			return fmt.Errorf("ERROR: Title for the ToDo can't be empty")
		}
		item.Title = userInput
	}

	userDesc := cmdFlags.desc
	if !cmd.Flags().Changed("detail") {
		userDesc = utils.GetUserInput("Any Description/Details for the todo (optional): ")
	}
	item.Desc = strings.TrimSpace(userDesc)

	err := getPriority(cmd, item)
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
	if !cmd.Flags().Changed("due") {
		dueDate = utils.GetUserInput("Deadline for todo item (optional): ")
	}

	trimmedDueDate := strings.TrimSpace(dueDate)
	if trimmedDueDate != "" {
		formattedDate, err := validateAndParseDueDate(trimmedDueDate)
		if err != nil {
			return err
		}
		item.DueDate = formattedDate
	}
	return nil
}

func validateAndParseDueDate(dueDate string) (string, error) {
	itemDue, err := utils.FormatParsedDate(dueDate)
	if err != nil {
		// todo: reprompt if invalid date instead of error
		return "", fmt.Errorf("ERROR: Invalid Date Format: %w", err)
	}
	return itemDue, nil
}

func getPriority(cmd *cobra.Command, item *model.Todo) error {
	itemPriority := cmdFlags.priority
	if !cmd.Flags().Changed("priority") {
		inputPriority := strings.TrimSpace(utils.GetUserInput("Want to add priority for the todo? By default it's none: "))
		if inputPriority == "" {
			return nil
		}
		parsedPriority, err := strconv.Atoi(inputPriority)
		if err != nil {
			return fmt.Errorf("ERROR: Invalid priority: %w", err)
		}
		itemPriority = parsedPriority
	}

	err := validatePriority(itemPriority)
	if err != nil {
		return err
	}

	item.Priority = &itemPriority
	return nil
}

func validatePriority(priority int) error {
	if priority < 0 || priority > 2 {
		// todo: reprompt user instead of throwing error
		return fmt.Errorf("ERROR: Invalid priority, please select priority b/w 0 and 2 (2 being lowest)")
	}
	return nil
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolVarP(&cmdFlags.interactive, "interactive", "i", false, "Add todo details in interactive mode")
	addCmd.Flags().StringVarP(&cmdFlags.desc, "detail", "d", "", "description/details for the todo item")
	addCmd.Flags().IntVarP(&cmdFlags.priority, "priority", "p", 0, "priority for the todo item")
	addCmd.Flags().StringVarP(&cmdFlags.dueDate, "due", "t", "", "due date for the todo item")
}
