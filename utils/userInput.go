package utils

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func GetUserInput(prompt string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(prompt)
	userInput, _ := reader.ReadString('\n')
	return strings.TrimSpace(userInput)
}
