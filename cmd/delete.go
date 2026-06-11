/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strings"
	"os"
	"bufio"

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
		if idArgs == "" {
			return fmt.Errorf("invalid todo id arg: %s", args[0])
		}

		if isConfirm {
			fmt.Println("deleting todo item with id:", idArgs)
			return nil
		}

		reader := bufio.NewReader(os.Stdin)
		fmt.Printf("Are you sure you want to delete todo #%s: ", idArgs)
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimSpace(userInput)
		if userInput == "yes" || userInput == "y" {
			fmt.Printf("deleting todo item with id: %s\n", idArgs)
			return nil
		}

		fmt.Printf("todo delete operation cancelled for: %s\n", idArgs)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().BoolVarP(&isConfirm, "skip confirmation", "y", false, "Delete todo without confirmation")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
