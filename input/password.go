package input

import (
	"bufio"
	"fmt"
	"os"

	"golang.org/x/term"
)

func GetPrivate(msg string) (string, error) {
	fmt.Print(msg)

	fd := int(os.Stdin.Fd())

	bytePrivate, err := term.ReadPassword(fd)
	if err != nil {
		return "", err
	}
	fmt.Println()

	return string(bytePrivate), nil
}

func GetLogin(scanner *bufio.Scanner) (string, string, error) {
	username, err := String(scanner, "Username: ")
	if err != nil {
		return "", "", err
	}
	password, err := GetPrivate("Password: ")
	if err != nil {
		return "", "", err
	}
	return username, password, nil
}
