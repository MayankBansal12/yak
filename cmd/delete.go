/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"yak-cli/internal/storage"

	"github.com/spf13/cobra"
)

var isConfirm bool

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		idArgs := strings.TrimSpace(strings.TrimLeft(args[0], "#"))
		id, err := strconv.Atoi(idArgs)
		if err != nil {
			return fmt.Errorf("Invalid todo id arg: %s", args[0])
		}

		if !isConfirm {
			reader := bufio.NewReader(os.Stdin)
			fmt.Printf("Are you sure you want to delete todo #%v: (y or n) ", id)
			userInput, _ := reader.ReadString('\n')
			userInput = strings.TrimSpace(userInput)

			if userInput != "yes" && userInput != "y" {
				fmt.Printf("Delete operation cancelled for Todo with id: #%v\n", id)
				return nil
			}
		}

		err = storage.DeleteTodo(id)
		if err != nil {
			return err
		}
		fmt.Printf("Todo with id: #%v deleted successfully\n", id)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().BoolVarP(&isConfirm, "confirm", "y", false, "Delete todo without confirmation")
}
