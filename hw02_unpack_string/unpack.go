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
	var escape bool = false

	for _, char := range str {
		if string(char) == `\` {
			if escape {
				prev = string(char)
				escape = false
				continue
			}
			escape = true
			result = result + prev
			continue
		}

		if unicode.IsLetter(char) {
			if prev == "" {
				prev = string(char)
			} else {
				result = result + prev
				prev = string(char)
			}
			escape = false
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
			result = result + strings.Repeat(prev, multiplicity)
			prev = ""
		}
	}
	if escape {
		err = ErrInvalidString
	}

	if prev != "" {
		result = result + prev
	}

	if err != nil {
		return "", err
	}

	return result, err
}
