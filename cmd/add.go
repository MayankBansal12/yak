package cmd

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/mayankbansal12/yak/internal/display"
	"github.com/mayankbansal12/yak/internal/model"
	"github.com/mayankbansal12/yak/internal/storage"
	"github.com/mayankbansal12/yak/utils"

	"github.com/spf13/cobra"
)

type addFlags struct {
	desc        string
	priority    int
	dueDate     string
	interactive bool
	jsonOutput  bool
}

var cmdFlags addFlags

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add [title]",
	Short: "Add a new todo item",
	Args:  cobra.RangeArgs(0, 1),
	Long: `Add a new todo to your list.

The title can be passed as the first argument, or entered when prompted.
Optional fields — description, priority, and due date — can be supplied
as flags.
Add todo interactively with -i or --interactive. Use -j to output the created
todo as JSON.

Examples:
  yak add "Buy milk"
  yak add "Pay rent" -p 1 -d "Include utilities" -t 25-03
  yak add -i (interactive mode)
  yak add "Buy milk" -j`,
	RunE: func(cmd *cobra.Command, args []string) error {
		var item model.Todo

		title := ""
		hasTitle := len(args) > 0
		if hasTitle {
			title = strings.TrimSpace(args[0])
		}

		item.Title = title

		if cmdFlags.interactive {
			err := interactionModeForAdd(cmd, &item)
			if err != nil {
				return err
			}
			err = storage.AddTodo(item)
			if err != nil {
				return err
			}

			display.PrintTodoJSON(item, time.Now())
			return nil
		}

		if title == "" {
			fmt.Println("did you forget -i flag?\n\nEither provide a title or run in interactive mode with -i flag")
			return nil
		}

		if cmd.Flags().Changed("detail") {
			item.Desc = strings.TrimSpace(cmdFlags.desc)
		}

		if cmd.Flags().Changed("priority") {
			err := getPriority(cmd, &item)
			if err != nil {
				return err
			}
		}

		if cmd.Flags().Changed("due") {
			err := getDueDate(cmd, &item)
			if err != nil {
				return err
			}
		}

		err := storage.AddTodo(item)
		if err != nil {
			return err
		}

		display.PrintTodoJSON(item, time.Now())
		return nil
	},
}

func interactionModeForAdd(cmd *cobra.Command, item *model.Todo) error {
	isTitleEmpty := item.Title == ""
	if isTitleEmpty {
		userInput, err := utils.PromptUntilInput("Title for the todo: ", "Title for the ToDo can't be empty")
		if err != nil {
			return err
		}
		item.Title = userInput
	}

	userDesc := cmdFlags.desc
	var err error
	if !cmd.Flags().Changed("detail") {
		userDesc, err = utils.GetUserInput("Any Description/Details for the todo (optional): ")
		if err != nil {
			return err
		}
	}
	item.Desc = strings.TrimSpace(userDesc)

	err = getPriority(cmd, item)
	if err != nil {
		return err
	}

	err = getDueDate(cmd, item)
	if err != nil {
		return err
	}
	return nil
}

func getDueDate(cmd *cobra.Command, item *model.Todo) error {
	dueDate := cmdFlags.dueDate
	var err error

	if !cmd.Flags().Changed("due") {
		dueDate, err = utils.GetUserInput("Deadline for todo item (optional): ")
		if err != nil {
			return err
		}
		if dueDate == "" {
			fmt.Println("Skipping Due Date...")
			return nil
		}
	}

	for {
		formattedDate, err := utils.FormatParsedDate(dueDate)
		if err == nil {
			item.DueDate = formattedDate
			return nil
		}

		fmt.Println("Warning: " + err.Error())
		dueDate, err = utils.GetUserInput("Deadline for todo item (optional): ")
		if err != nil {
			return err
		}
		if dueDate == "" {
			fmt.Println("Skipping Due Date...")
			return nil
		}
	}
}

func getPriority(cmd *cobra.Command, item *model.Todo) error {
	itemPriority := strconv.Itoa(cmdFlags.priority)
	var err error

	if !cmd.Flags().Changed("priority") {
		itemPriority, err = utils.GetUserInput("Add priority for the todo (0-2)...by default it's none: ")
		if err != nil {
			return err
		}
		if itemPriority == "" {
			fmt.Println("Skipping Priority...")
			return nil
		}
	} else if cmdFlags.priority == -1 {
		item.Priority = nil
		return nil
	}

	for {
		parsed, err := validatePriority(itemPriority)
		if err == nil {
			item.Priority = &parsed
			return nil
		}

		fmt.Println("Warning: " + err.Error())
		itemPriority, err = utils.GetUserInput("Priority for item (0-2, or Enter to skip): ")
		if err != nil {
			return err
		}
		if itemPriority == "" {
			fmt.Println("Skipping Priority...")
			return nil
		}
	}
}

func validatePriority(priority string) (int, error) {
	parsedPriority, err := strconv.Atoi(priority)
	if err != nil {
		return 0, fmt.Errorf("Invalid priority: %w", err)
	}
	if parsedPriority < 0 || parsedPriority > 2 {
		return 0, fmt.Errorf("Invalid priority, please select priority b/w 0 and 2 (2 being lowest)")
	}
	return parsedPriority, nil
}

func init() {
	rootCmd.AddCommand(addCmd)
	addCmd.Flags().BoolVarP(&cmdFlags.interactive, "interactive", "i", false, "Add todo details in interactive mode")
	addCmd.Flags().StringVarP(&cmdFlags.desc, "detail", "d", "", "description/details for the todo item")
	addCmd.Flags().IntVarP(&cmdFlags.priority, "priority", "p", -1, "priority for the todo item")
	addCmd.Flags().StringVarP(&cmdFlags.dueDate, "due", "t", "", "due date for the todo item")
	addCmd.Flags().BoolVarP(&cmdFlags.jsonOutput, "json", "j", false, "Output the created todo as JSON")
}
