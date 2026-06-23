package utils

import (
	"fmt"
	"strings"
	"time"
)

const goReferenceTime = "02-01-06"

func ParseStringToTime(raw string) (time.Time, error) {
	t, err := time.Parse(time.RFC3339, raw)
	if err != nil {
		return time.Time{}, err
	}
	return GetDateFromTimestamp(t, time.Local), nil
}

func GetDateFromTimestamp(t time.Time, loc *time.Location) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, loc)
}

func FormatParsedDate(dateInput string) (string, error) {
	dateInput = strings.TrimSpace(dateInput)
	if dateInput != "" {
		parsedDate, err := time.Parse(goReferenceTime, dateInput)
		if err != nil {
			parsedDate, err = time.Parse("02-01", dateInput)

			if err != nil {
				return "", fmt.Errorf("invalid date format %s", dateInput)
			}

			parsedDate = parsedDate.AddDate(time.Now().Year()-parsedDate.Year(), 0, 0)
		}
		return parsedDate.Format(time.RFC3339), nil
	}
	return "", fmt.Errorf("Invalid or empty date input, please follow DD-MM format")
}

func FormatDueDate(dueDate string, now time.Time) string {
	if dueDate == "" {
		return ""
	}

	due, err := ParseStringToTime(dueDate)
	if err != nil {
		return dueDate
	}

	today := GetDateFromTimestamp(now, time.Local)
	days := int(due.Sub(today).Hours() / 24)

	switch {
	case days == 0:
		return "today"
	case days == 1:
		return "tomorrow"
	case days == -1:
		return "yesterday"
	case days > 1 && days <= 7:
		return fmt.Sprintf("in %d days", days)
	case days < -1 && days >= -7:
		return fmt.Sprintf("%d days overdue", -days)
	case due.Year() != today.Year():
		return due.Format("Jan 02, 2006")
	default:
		return due.Format("Jan 02")
	}
}
