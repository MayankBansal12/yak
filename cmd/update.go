package cmd

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/mayankbansal12/yak/internal/model"
	"github.com/mayankbansal12/yak/internal/storage"
	"github.com/mayankbansal12/yak/utils"

	"github.com/spf13/cobra"
)

type updateFlags struct {
	id       int
	desc     string
	priority int
	dueDate  string
}

var updateArgs updateFlags

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update [todo-id] [title]",
	Short: "Update an existing todo item",
	Long: `Update an existing todo item by its ID.

A new title can be passed as the second argument. Optional fields — description,
priority, and due date — can be supplied as flags.

Examples:
  yak update 3 "New title"
  yak update 3 -p 1 -d "Updated description" -t 25-03`,
	Args: cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		var item model.Todo

		idArg := strings.TrimSpace(strings.TrimLeft(args[0], "#"))
		todoId, err := strconv.Atoi(idArg)
		if err != nil {
			return fmt.Errorf("todo id is required for update command")
		}

		hasTitle := len(args) > 1
		if !hasTitle && cmd.Flags().NFlag() == 0 {
			return fmt.Errorf("either a title or at least one update flag is required")
		}

		if hasTitle && strings.TrimSpace(args[1]) != "" {
			item.Title = strings.TrimSpace(args[1])
		}

		if updateArgs.desc != "" {
			item.Desc = updateArgs.desc
		}

		if updateArgs.priority != -1 {
			err := getUpdatedPriority(&item)
			if err != nil {
				return nil
			}
		}

		if updateArgs.dueDate != "" {
			err := getUpdatedDueDate(&item)
			if err != nil {
				return err
			}
		}

		err = storage.UpdateTodo(todoId, item)
		if err != nil {
			return err
		}
		fmt.Printf("Updated the todo item: #%v\n", todoId)
		return nil
	},
}

func getUpdatedDueDate(item *model.Todo) error {
	dueDate := updateArgs.dueDate
	for {
		formattedDate, err := utils.FormatParsedDate(dueDate)
		if err == nil {
			item.DueDate = formattedDate
			return nil
		}

		fmt.Println("Warning: " + err.Error())
		userInput, err := utils.GetUserInput("Deadline for todo item (optional): ")
		if err != nil {
			return err
		}
		if userInput == "" {
			fmt.Println("Skipping Due Date...")
			return nil
		}
		dueDate = userInput
	}
}

func getUpdatedPriority(item *model.Todo) error {
	priority := strconv.Itoa(updateArgs.priority)
	for {
		parsedPriority, err := validatePriority(priority)
		if err == nil {
			item.Priority = &parsedPriority
			return nil
		}

		fmt.Println("Warning: " + err.Error())
		itemPriority, err := utils.GetUserInput("Priority for item (0-2, or Enter to skip): ")
		if err != nil {
			return err
		}
		if itemPriority == "" {
			fmt.Println("Skipping Priority...")
			return nil
		}
	 	priority = itemPriority
	}
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVarP(&updateArgs.desc, "detail", "d", "", "description/details for the todo item")
	updateCmd.Flags().IntVarP(&updateArgs.priority, "priority", "p", -1, "priority for the todo item")
	updateCmd.Flags().StringVarP(&updateArgs.dueDate, "due", "t", "", "due date for the todo item")
}
