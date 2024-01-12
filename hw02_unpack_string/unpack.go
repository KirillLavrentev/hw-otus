package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func IsDigit(digit rune) bool {
	if digit >= '0' && digit <= '9' {
		return true
	}
	return false
}

func Unpack(myStr string) (string, error) {
	var newString string

	runes := []rune(myStr)
	for i := 0; i < len(runes); i++ {
		switch {
		case IsDigit(runes[i]) && i == 0:
			return "error", ErrInvalidString
		case IsDigit(runes[i]):
			digit, err := strconv.Atoi(string(runes[i]))
			if err != nil {
				return "", err
			}
			previousRune := string(runes[i-1])
			if IsDigit(runes[i-1]) {
				return "error", ErrInvalidString
			}
			if digit == 0 {
				newString = newString[:len(newString)-len(previousRune)]
			} else {
				newString += strings.Repeat(previousRune, digit-1)
			}
		default:
			newString += string(runes[i])
		}
	}
	return newString, nil
}
