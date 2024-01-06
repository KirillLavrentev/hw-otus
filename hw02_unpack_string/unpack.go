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
	// Place your code here.

	var previousRune string
	var newString string

	for i, v := range myStr {
		switch {
		case IsDigit(v) && i == 0:
			return "error", ErrInvalidString
		case IsDigit(v):
			digit, err := strconv.Atoi(string(v))
			if err != nil {
				return "", err
			}
			a := myStr[i-1]
			b := "_"
			if i >= 2 {
				b = string(myStr[i-2])
			}
			if b == "\\" {
				previousRune = b + string(a)
			} else {
				previousRune = string(a)
			}
			if IsDigit(rune(a)) {
				return "error", ErrInvalidString
			}
			if digit == 0 {
				newString = newString[:len(newString)-len(previousRune)]
			} else {
				newString += strings.Repeat(previousRune, digit-1)
			}
		default:
			previousRune = string(myStr[i])
			newString += previousRune
		}
	}

	return newString, nil
}
