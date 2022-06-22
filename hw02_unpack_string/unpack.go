package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var result, prev string
	var err error
	escape := false

	for _, char := range str {
		if string(char) == `\` {
			if escape {
				prev = string(char)
				escape = false
				continue
			}
			escape = true
			result += prev
			continue
		}

		if unicode.IsDigit(char) {
			if prev == "" {
				err = ErrInvalidString
				break
			}

			if escape {
				prev = string(char)
				escape = false
				continue
			}

			multiplicity, _ := strconv.Atoi(string(char))
			result += strings.Repeat(prev, multiplicity)
			prev = ""
			continue
		}

		if prev == "" {
			prev = string(char)
		} else {
			result += prev
			prev = string(char)
		}
		escape = false
	}
	if escape {
		err = ErrInvalidString
	}

	if prev != "" {
		result += prev
	}

	if err != nil {
		return "", err
	}

	return result, err
}
