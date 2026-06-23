package display

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

var noColor = os.Getenv("NO_COLOR") != "" || strings.ToLower(os.Getenv("TERM")) == "dumb"

// Color wraps text in ANSI color codes when stdout is a terminal and NO_COLOR is unset.
func Color(text, color string) string {
	if noColor {
		return text
	}
	if !term.IsTerminal(int(os.Stdout.Fd())) {
		return text
	}

	switch color {
	case "red":
		return fmt.Sprintf("\033[31m%s\033[0m", text)
	case "green":
		return fmt.Sprintf("\033[32m%s\033[0m", text)
	case "yellow":
		return fmt.Sprintf("\033[33m%s\033[0m", text)
	default:
		return text
	}
}
