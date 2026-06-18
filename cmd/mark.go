/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"yak-cli/internal/storage"

	"github.com/spf13/cobra"
)

// markCmd represents the mark command
var markCmd = &cobra.Command{
	Use:     "mark [todo-id]",
	Aliases: []string{"done", "donee", "doneee"},
	Short:   "Mark a todo as completed",
	Long: `Mark a todo item as completed by its ID.

Examples:
  yak mark 3
  yak done 3`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		idArg := strings.TrimSpace(strings.TrimLeft(args[0], "#"))
		id, err := strconv.Atoi(idArg)
		if err != nil {
			return fmt.Errorf("todo id is required for mark command")
		}

		err = storage.MarkTodo(id)
		if err != nil {
			return err
		}

		fmt.Printf("Todo with id: %v marked as completed\n", id)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(markCmd)
}
