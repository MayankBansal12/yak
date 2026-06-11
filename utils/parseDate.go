
package utils

import (
	"fmt"
	"strings"
	"time"
)

const goReferenceTime = "02-01-06"

func FormatParsedDate(dateInput string) (string, error) {
	dateInput = strings.TrimSpace(dateInput)
	if dateInput != "" {
		parsedDate, err := time.Parse(goReferenceTime, dateInput)
		if err != nil {
			parsedDate, err = time.Parse("02-01", dateInput)

			if err != nil{
				return "", fmt.Errorf("invalid date format %s", dateInput)
			}

			parsedDate = parsedDate.AddDate(time.Now().Year()-parsedDate.Year(), 0, 0)
		}
		return parsedDate.Format(goReferenceTime), nil
	}
	return "", fmt.Errorf("Invalid or empty date input, please follow DD-MM format")
}
