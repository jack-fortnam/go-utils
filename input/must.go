package input

import (
	"bufio"
	"fmt"
)

func retry[T any](fn func() (T, error)) T {
	for {
		v, err := fn()
		if err == nil {
			return v
		}
		fmt.Println("error inputting")
	}
}

func MustString(scanner *bufio.Scanner, prompt string) string {
	return retry(func() (string, error) {
		return String(scanner, prompt)
	})
}

func MustInt(scanner *bufio.Scanner, prompt string) int {
	return retry(func() (int, error) {
		return Int(scanner, prompt)
	})
}

func MustFloat(scanner *bufio.Scanner, prompt string) float64 {
	return retry(func() (float64, error) {
		return Float(scanner, prompt)
	})
}

func MustChoice(scanner *bufio.Scanner, prompt string, valid []string) string {
	return retry(func() (string, error) {
		return Choice(scanner, prompt, valid)
	})
}

func MustMatch(scanner *bufio.Scanner, prompt string, pattern string) string {
	return retry(func() (string, error) {
		return Match(scanner, prompt, pattern)
	})
}

func MustIntRange(scanner *bufio.Scanner, prompt string, min int, max int) int {
	return retry(func() (int, error) {
		return IntRange(scanner, prompt, min, max)
	})
}
