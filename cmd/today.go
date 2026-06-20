package cmd

import (
	"cmp"
	"fmt"
	"github.com/mayankbansal12/yak/internal/display"
	"github.com/mayankbansal12/yak/internal/model"
	"github.com/mayankbansal12/yak/internal/storage"
	"slices"
	"time"

	"github.com/spf13/cobra"
)

type todayArguments struct {
	lastN   int
	showAll bool
}

var todayArgs todayArguments

// todayCmd represents the today command
var todayCmd = &cobra.Command{
	Use:   "today",
	Short: "List today's todos",
	Long: `List all todo items due today.

Use -a or --all to show all todos for today, or -n to limit results.

Examples:
  yak today
  yak today -a
  yak today -n 5`,
	RunE: func(cmd *cobra.Command, args []string) error {
		storedTodos, err := storage.GetTodos()
		if err != nil {
			return err
		}

		todayTime := time.Now().Format("02-01-2006")
		var upcomingTodos []model.Todo
		for _, todo := range storedTodos {
			t, _ := time.Parse(time.RFC3339, todo.DueDate)
			dueDateOnly := t.Format("02-01-2006")
			if dueDateOnly == todayTime {
				upcomingTodos = append(upcomingTodos, todo)
			}
		}

		slices.SortFunc(upcomingTodos, func(a, b model.Todo) int {
			if a.Priority == nil && b.Priority == nil {
				return 0
			}
			if a.Priority == nil {
				return 1
			}
			if b.Priority == nil {
				return -1
			}
			return cmp.Compare(*a.Priority, *b.Priority)
		})

		todosLen := len(upcomingTodos)
		if todosLen == 0 {
			fmt.Println("No upcoming todos found for today")
			return nil
		}

		listLen := todosLen
		if todayArgs.lastN > 0 {
			listLen = min(todayArgs.lastN, todosLen)
		} else if !todayArgs.showAll {
			listLen = 3
		}
		display.PrintTodos(upcomingTodos[:listLen], true)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(todayCmd)
	todayCmd.Flags().IntVarP(&todayArgs.lastN, "last-N", "n", -1, "List last n todos for today")
	todayCmd.Flags().BoolVarP(&todayArgs.showAll, "all", "a", false, "List all todos for today")
}
