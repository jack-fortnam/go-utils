package validate

import (
	"errors"
	"net/mail"
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

var phoneRegex = regexp.MustCompile(`^\+[1-9]\d{7,14}$`)
var PostCodeReg = regexp.MustCompile(`^[A-Z]{1,2}\d[A-Z\d]?\d[A-Z]{2}$`)

func Email(email string) error {
	_, err := mail.ParseAddress(email)
	return err
}

func PhoneNumber(phone string) error {
	if !phoneRegex.MatchString(phone) {
		return errors.New("invalid phone number")
	}
	return nil
}

func PostCode(code string) error {
	if !PostCodeReg.MatchString(NormalizePostcode(code)) {
		return errors.New("invalid postcode")
	}
	return nil
}

func NormalizePostcode(pc string) string {
	return strings.ToUpper(strings.ReplaceAll(pc, " ", ""))
}

func Card(number string) bool {
	number = strings.ReplaceAll(number, " ", "")
	number = strings.ReplaceAll(number, "-", "")

	if len(number) < 2 {
		return false
	}

	var sum int
	shouldDouble := false

	for i := len(number) - 1; i >= 0; i-- {
		digit, err := strconv.Atoi(string(number[i]))
		if err != nil {
			return false
		}

		if shouldDouble {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		shouldDouble = !shouldDouble
	}

	return sum%10 == 0
}

func Url(URL string) error {
	u, err := url.ParseRequestURI(URL)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return errors.New("invalid URL")
	}
	return nil
}
