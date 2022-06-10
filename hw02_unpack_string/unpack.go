package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var res, prev string
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
			res = res + prev
			continue
		} else if unicode.IsLetter(char) {
			if prev == "" {
				prev = string(char)
			} else {
				res = res + prev
				prev = string(char)
			}
			escape = false
		} else if unicode.IsDigit(char) {
			if escape {
				prev = string(char)
				escape = false
				continue
			} else if prev == "" {
				err = ErrInvalidString
				break
			} else {
				multiplicity, _ := strconv.Atoi(string(char))
				res = res + strings.Repeat(prev, multiplicity)
				prev = ""
			}
		}
	}
	if prev != "" {
		res = res + prev
	}

	if escape {
		res = ""
		err = ErrInvalidString
	}

	return res, err
}
