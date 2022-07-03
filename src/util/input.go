package util

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func ReadInt() int {
	reader := bufio.NewReader(os.Stdin)
	inputText, err := reader.ReadString('\n')
	if err != nil {
		return 0
	}
	cleanInput := strings.TrimSpace(inputText)
	input, err := strconv.Atoi(cleanInput)
	if err != nil {
		return 0
	}

	return input
}

func ReadString() string {
	reader := bufio.NewReader(os.Stdin)
	inputText, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	return strings.TrimSpace(inputText)
}

func ReadStrings() (string, string) {
	reader := bufio.NewReader(os.Stdin)
	inputText, err := reader.ReadString('\n')
	if err != nil {
		return "", ""
	}
	cleanInput := strings.TrimSpace(inputText)
	strArray := strings.Split(cleanInput, "")
	return strArray[0], strArray[1]
}

func Valid(input int) bool {
	return input >= 1 && input <= 4
}
