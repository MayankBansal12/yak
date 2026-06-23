package display

import (
	"time"

	"github.com/mayankbansal12/yak/internal/model"
	"github.com/mayankbansal12/yak/utils"
)

const (
	iconCompleted = "●"
	iconPending   = "○"
	iconOverdue   = "!"
)

func Icon(t model.Todo, now time.Time) string {
	if t.CompletedAt != "" {
		return iconCompleted
	}
	if IsOverdue(t, now) {
		return iconOverdue
	}
	return iconPending
}

func IsOverdue(t model.Todo, now time.Time) bool {
	if t.CompletedAt != "" || t.DueDate == "" {
		return false
	}
	due, err := utils.ParseStringToTime(t.DueDate)
	if err != nil {
		return false
	}
	today := utils.GetDateFromTimestamp(now, time.Local)
	return due.Before(today)
}

func IsDueToday(t model.Todo, now time.Time) bool {
	if t.DueDate == "" {
		return false
	}
	due, err := utils.ParseStringToTime(t.DueDate)
	if err != nil {
		return false
	}
	today := utils.GetDateFromTimestamp(now, time.Local)
	return due.Equal(today)
}
