package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "yak",
	Version: "0.1.0",
	Short: "A fast and minimal CLI to manage todos",
	Long: `yak is a lightweight CLI todo manager that makes managing your daily tasks simple and fast.
Run the help command for instructions or go through README.md to know more about CLI.

Quick start:
  yak add "Buy groceries"           Add a todo
  yak add "Pay rent" -p 0 -t 25-06  Add with high priority & deadline 25 June 2026
  yak list                          List all pending todos
  yak today                         Show today's top 3 todos
  yak next                          Show upcoming todos by deadline
  yak mark 1                       Mark todo #1 as done`,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}
