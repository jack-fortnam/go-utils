package input

import (
	"bufio"
)

func retry[T any](fn func() (T, error)) T {
	for {
		v, err := fn()
		if err == nil {
			return v
		}
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
