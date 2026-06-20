package cmd

import (
	"fmt"
	// "strings"

	// "github.com/mayankbansal12/yak/internal/model"
	// "github.com/mayankbansal12/yak/utils"

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
	// Args: cobra.RangeArgs(1, 2),
	RunE: func(cmd *cobra.Command, args []string) error {
		// var item model.Todo

		// idArg := strings.TrimSpace(strings.TrimLeft(args[0], "#"))
		// if idArg == "" {
		// 	return fmt.Errorf("ToDo item ID is required for the update command")
		// }

		// hasTitle := len(args) > 1
		// if !hasTitle && cmd.Flags().NFlag() == 0 {
		// 	return fmt.Errorf("either a title or at least one flag is required")
		// }
		// if hasTitle && strings.TrimSpace(args[1]) != "" {
		// 	item.Title = strings.TrimSpace(args[1])
		// }

		// if updateArgs.desc != "" {
		// 	item.Desc = updateArgs.desc
		// }

		// if updateArgs.priority != -1 {
		// 	if updateArgs.priority < 0 || updateArgs.priority > 2 {
		// 		return fmt.Errorf("invalid priority %d: must be 0, 1, or 2", updateArgs.priority)
		// 	}
		// 	item.Priority = &updateArgs.priority
		// }

		// if updateArgs.dueDate != "" {
		// 	itemDue, err := utils.FormatParsedDate(updateArgs.dueDate)
		// 	if err != nil {
		// 		return fmt.Errorf("%w", err)
		// 	}
		// 	updateArgs.dueDate = itemDue
		// 	item.DueDate = updateArgs.dueDate
		// }

		// todo: update todo in local file
		// fmt.Printf("Updated the todo item %s : %v\n", idArg, updateArgs)
		//
		fmt.Println("Update Command isn't available yet! Trying to ship fast, please wait!")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVarP(&updateArgs.desc, "detail", "d", "", "description/details for the todo item")
	updateCmd.Flags().IntVarP(&updateArgs.priority, "priority", "p", 0, "priority for the todo item")
	updateCmd.Flags().StringVarP(&updateArgs.dueDate, "due", "t", "", "due date for the todo item")
}
