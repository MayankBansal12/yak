/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"yak-cli/internal/display"
	"yak-cli/internal/model"
	"yak-cli/internal/storage"
	"yak-cli/utils"

	"github.com/spf13/cobra"
)

type listArgs struct {
	lastN int
	due string
	showDetail bool
}

var commandArgs listArgs

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Aliases: []string{"view", "see", "get", "find"},
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.RangeArgs(0, 1),
	RunE: func(cmd *cobra.Command, args []string) error {
		storedTodos, err := storage.GetTodos()
		if err != nil {
			return fmt.Errorf("%w", err)
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
				return fmt.Errorf("%w", err)
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

func getTodoById (todos []model.Todo, id int) []model.Todo {
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
