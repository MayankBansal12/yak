package display

import (
	"fmt"
	"strings"
	"time"

	"yak-cli/internal/model"
)

func PrintTodos(todos []model.Todo, detailed bool) {
	if len(todos) == 0 {
		fmt.Println("No todos found.")
		return
	}

	width := 40
	for _, t := range todos {
		if len(t.Title)+6 > width {
			width = len(t.Title) + 6
		}
	}

	for _, t := range todos {
		printCard(t, width, detailed)
	}
}

func printCard(t model.Todo, width int, detailed bool) {
	top := "┌" + strings.Repeat("─", width) + "┐"
	bottom := "└" + strings.Repeat("─", width) + "┘"

	fmt.Println(top)
	printRow(fmt.Sprintf("#%-4d %s", t.ID, t.Title), width)

	if detailed {
		if t.Desc != "" {
			printRow(t.Desc, width)
		}
	}

	if t.DueDate != "" {
		printRow(fmt.Sprintf("Due: %s", formatDate(t.DueDate)), width)
	}

	if t.Priority != nil {
		printRow(fmt.Sprintf("Priority: %s", priorityLabel(*t.Priority)), width)
	}

	if detailed {
		printRow(fmt.Sprintf("Created: %s", formatDate(t.CreatedAt)), width)
		printRow(fmt.Sprintf("Updated: %s", formatDate(t.UpdatedAt)), width)
	}

	if t.CompletedAt != "" {
		printRow(fmt.Sprintf("Status: Done (%s)", formatDate(t.CompletedAt)), width)
	} else {
		printRow("Status: Pending", width)
	}

	fmt.Println(bottom)
	fmt.Println()
}

func printRow(content string, width int) {
	inner := width - 2
	if len(content) > inner {
		content = content[:inner-3] + "..."
	}
	padding := inner - len(content)
	fmt.Printf("│ %s%s │\n", content, strings.Repeat(" ", padding))
}

func formatDate(raw string) string {
	t, err := time.Parse(time.RFC3339, raw)
	if err != nil {
		return raw
	}
	return t.Format("Jan 02, 2006")
}

func priorityLabel(p int) string {
	switch p {
	case 0:
		return "High"
	case 1:
		return "Medium"
	case 2:
		return "Low"
	default:
		return "None"
	}
}
