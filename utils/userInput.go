package utils

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func GetUserInput(prompt string) (string, error) {
	fmt.Print(prompt)
	userInput, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(userInput), nil
}

func PromptUntilInput(prompt string, warningText string) (string, error) {
	var userInput string
	var err error

	for userInput == "" {
		fmt.Print(prompt)
		userInput, err = reader.ReadString('\n')
		if err != nil {
			return "", err
		}
		userInput = strings.TrimSpace(userInput)
		if userInput == "" && warningText != "" {
			fmt.Println(warningText)
		}
	}
	return userInput, nil
}
