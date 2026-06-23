package display

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/mayankbansal12/yak/internal/model"
	"github.com/mayankbansal12/yak/utils"
	"golang.org/x/term"
)

func printCompact(todos []model.Todo, opts Options) {
	width := terminalWidth()
	for _, t := range todos {
		line := compactLine(t, opts)
		if len(line) > width {
			line = line[:width-3] + "..."
		}
		fmt.Println(line)
	}
}

func compactLine(t model.Todo, opts Options) string {
	parts := []string{
		Color(Icon(t, opts.Now), iconColor(t, opts.Now)),
		fmt.Sprintf("#%d", t.ID),
		t.Title,
	}
	if t.DueDate != "" {
		parts = append(parts, utils.FormatDueDate(t.DueDate, opts.Now))
	}
	if t.Priority != nil {
		parts = append(parts, PriorityBars(*t.Priority))
	}
	return strings.Join(parts, "  ")
}

func iconColor(t model.Todo, now time.Time) string {
	if t.CompletedAt != "" {
		return "green"
	}
	if IsOverdue(t, now) {
		return "red"
	}
	return ""
}

func terminalWidth() int {
	w, _, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil || w < 40 {
		return 100
	}
	return w
}
