package input

import (
	"bufio"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"
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

func MultiLineString(scanner *bufio.Scanner, prompt string) ([]string, error) {
	fmt.Print(prompt)
	var lines []string
	for {
		scanner.Scan()
		line := scanner.Text()
		if len(line) == 0 {
			break
		}
		lines = append(lines, line)
	}

	err := scanner.Err()
	if err != nil {
		return nil, err
	}
	return lines, nil
}

func Bool(scanner *bufio.Scanner, prompt string) (bool, error) {
	fmt.Printf("%s (y/n)", prompt)
	if !scanner.Scan() {
		return false, scanner.Err()
	}
	text := strings.ToLower(scanner.Text())
	switch text {
	case "y":
		return true, nil
	case "n":
		return false, nil
	default:
		return false, errors.New("Invalid input")
	}
}

func Choice(scanner *bufio.Scanner, prompt string, valid []string) (string, error) {
	fmt.Printf("%s (%s): ", prompt, strings.Join(valid, "/"))
	if !scanner.Scan() {
		return "", scanner.Err()
	}
	input := strings.TrimSpace(scanner.Text())
	for _, v := range valid {
		if strings.EqualFold(input, v) {
			return v, nil
		}
	}
	return "", fmt.Errorf("invalid choice: %s", input)
}

func Match(scanner *bufio.Scanner, prompt string, pattern string) (string, error) {
	val, err := String(scanner, prompt)
	if err != nil {
		return "", err
	}
	matched, _ := regexp.MatchString(pattern, val)
	if !matched {
		return "", fmt.Errorf("input does not match required format")
	}
	return val, nil
}

func IntRange(scanner *bufio.Scanner, prompt string, min int, max int) (int, error) {
	prompt = fmt.Sprintf("%s,(%d-%d)", prompt, min, max)
	n, err := Int(scanner, prompt)
	if err != nil {
		return 0, err
	}
	if (min < n) && (n < max) {
		return n, nil
	}
	return 0, errors.New("Integer not in range")
}
