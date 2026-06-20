package cmd

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/mayankbansal12/yak/internal/display"
	"github.com/mayankbansal12/yak/internal/model"
	"github.com/mayankbansal12/yak/internal/storage"
	"github.com/mayankbansal12/yak/utils"

	"github.com/spf13/cobra"
)

type listArgs struct {
	lastN      int
	due        string
	showDetail bool
}

var commandArgs listArgs

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:     "list [todo-id]",
	Aliases: []string{"view", "see", "get", "find"},
	Short:   "List todo items",
	Long: `List all todo items, or view a specific todo by its ID.

By default, todos are sorted by most recently updated. Use -n to limit results,
-t to filter by due date, or pass a todo ID to view its details.

Examples:
  yak list
  yak list 3
  yak list -n 5
  yak list -t 25-03`,
	Args: cobra.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		storedTodos, err := storage.GetTodos()
		if err != nil {
			return err
		}

		if len(args) > 0 {
			idArg := strings.TrimSpace(strings.TrimLeft(args[0], "#"))
			id, err := strconv.Atoi(idArg)
			if err != nil {
				return fmt.Errorf("todo id is required for list command")
			}
			commandArgs.showDetail = true
			storedTodos = getTodoById(storedTodos, id)
		} else {
			slices.SortFunc(storedTodos, func(a, b model.Todo) int {
				return strings.Compare(b.UpdatedAt, a.UpdatedAt)
			})
		}

		if commandArgs.lastN > 0 {
			limit := min(commandArgs.lastN, len(storedTodos))
			storedTodos = storedTodos[:limit]
		} else if commandArgs.due != "" {
			dueDate, err := utils.FormatParsedDate(commandArgs.due)
			if err != nil {
				return err
			}

			dueTodos := []model.Todo{}
			for _, todo := range storedTodos {
				if todo.DueDate == dueDate {
					dueTodos = append(dueTodos, todo)
				}
			}
			storedTodos = dueTodos
		}

		display.PrintTodos(storedTodos, commandArgs.showDetail)
		return nil
	},
}

func getTodoById(todos []model.Todo, id int) []model.Todo {
	for _, todo := range todos {
		if todo.ID == id {
			return []model.Todo{todo}
		}
	}
	fmt.Printf("todo with id %v not found\n", id)
	return nil
}

func init() {
	rootCmd.AddCommand(listCmd)
	listCmd.Flags().IntVarP(&commandArgs.lastN, "last-N", "n", -1, "List last n todos including completed ones")
	listCmd.Flags().StringVarP(&commandArgs.due, "datetime", "t", "", "List all todos assigned to a specific date")
	listCmd.Flags().BoolVarP(&commandArgs.showDetail, "detail", "d", false, "List todos in detailed format (title, description, priority, deadline, created at)")
}
