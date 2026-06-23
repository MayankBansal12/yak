package display

import "github.com/mattn/go-runewidth"

func displayWidth(s string) int {
	return runewidth.StringWidth(s)
}

func truncateByWidth(s string, max int) string {
	if max <= 0 {
		return ""
	}
	if displayWidth(s) <= max {
		return s
	}
	if max == 1 {
		return "…"
	}
	w := 0
	out := make([]rune, 0, len(s))
	for _, r := range s {
		rw := runewidth.RuneWidth(r)
		if w+rw+1 > max {
			break
		}
		out = append(out, r)
		w += rw
	}
	return string(out) + "…"
}
