package display

import (
	"fmt"
	"strings"
	"time"

	"github.com/mayankbansal12/yak/internal/model"
	"github.com/mayankbansal12/yak/utils"
)

func printCards(todos []model.Todo, opts Options) {
	for _, t := range todos {
		printCard(t, opts)
	}
}

func cardWidth(t model.Todo, opts Options) int {
	maxLen := 0
	for _, line := range cardLines(t, opts) {
		if w := displayWidth(line); w > maxLen {
			maxLen = w
		}
	}
	w := max(maxLen + 4, 40)
	cap := max(terminalWidth() - 4, 40)

	if w > cap {
		w = cap
	}
	return w
}

func cardLines(t model.Todo, opts Options) []string {
	icon := Icon(t, opts.Now)
	head := icon + "  " + fmt.Sprintf("#%d", t.ID) + "  "
	lines := []string{head + t.Title}

	hasDue := t.DueDate != ""
	hasPri := t.Priority != nil
	if hasDue || hasPri {
		var meta string
		if hasDue && hasPri {
			meta = "  Due: " + utils.FormatDueDate(t.DueDate, opts.Now) + "    " +
				PriorityBars(*t.Priority) + "  " + priorityLabel(*t.Priority)
		} else if hasDue {
			meta = "  Due: " + utils.FormatDueDate(t.DueDate, opts.Now)
		} else {
			meta = "  " + PriorityBars(*t.Priority) + "  " + priorityLabel(*t.Priority)
		}
		lines = append(lines, meta)
	}

	word := "Pending"
	switch {
	case t.CompletedAt != "":
		word = "Done"
	case IsOverdue(t, opts.Now):
		word = "Overdue"
	}
	lines = append(lines, "  Status: "+icon+"  "+word)

	if opts.Detailed {
		if t.Desc != "" {
			lines = append(lines, "  "+t.Desc)
		}
		lines = append(lines, "  Created: "+formatDate(t.CreatedAt))
		lines = append(lines, "  Updated: "+formatDate(t.UpdatedAt))
	}
	return lines
}

func printCard(t model.Todo, opts Options) {
	width := cardWidth(t, opts)
	fmt.Println("╭" + strings.Repeat("─", width) + "╮")
	for _, line := range cardLines(t, opts) {
		printRow(line, width)
	}
	fmt.Println("╰" + strings.Repeat("─", width) + "╯")
	fmt.Println()
}

func printRow(content string, width int) {
	inner := width - 2
	if w := displayWidth(content); w > inner {
		content = truncateByWidth(content, inner)
	}
	padding := max(inner - displayWidth(content), 0)
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
