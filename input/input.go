package input

import (
	"bufio"
	"fmt"
	"strconv"
)

func String(scanner *bufio.Scanner, prompt string) (string, error) {
	fmt.Print(prompt)
	if !scanner.Scan() {
		return "", scanner.Err()
	}
	return scanner.Text(), nil
}

func Int(scanner *bufio.Scanner, prompt string) (int, error) {
	text, err := String(scanner, prompt)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(text)
}

func Float(scanner *bufio.Scanner, prompt string) (float64, error) {
	text, err := String(scanner, prompt)
	if err != nil {
		return 0, err
	}
	return strconv.ParseFloat(text, 64)
}
