/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
	"time"
	"yak-cli/internal/display"
	"yak-cli/internal/model"
	"yak-cli/internal/storage"

	"github.com/spf13/cobra"
)

// nextCmd represents the next command
var nextCmd = &cobra.Command{
	Use:   "next",
	Short: "Show the next todo item",
	Long: `Show the next upcoming todo item from your list.

Examples:
  yak next`,
	RunE: func(cmd *cobra.Command, args []string) error {
		storedTodos, err := storage.GetTodos()
		if err != nil {
			return err
		}

		todayTime := time.Now().Format("02-01-2006")
		var upcomingTodo []model.Todo

		for _, todo := range storedTodos {
			if todo.CompletedAt != "" || todo.DueDate != "" {
				continue
			}
			t, _ := time.Parse(time.RFC3339, todo.DueDate)
			dueDateOnly := t.Format("02-01-2006")
			if dueDateOnly < todayTime {
				continue
			}
			upcomingTodo = append(upcomingTodo, todo)
		}

		slices.SortFunc(upcomingTodo, func(a, b model.Todo) int {
			dateCmp := strings.Compare(a.DueDate, b.DueDate)
			if dateCmp != 0 {
				return dateCmp
			}
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

		todosLen := len(upcomingTodo)
		if todosLen == 0 {
			fmt.Println("No next todos found")
			return nil
		}

		listLen := min(todosLen, 3)
		display.PrintTodos(upcomingTodo[:listLen], true)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(nextCmd)
}
