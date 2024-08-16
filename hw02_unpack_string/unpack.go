package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(strIn string) (string, error) {
	if strIn == "" {
		return "", nil
	}

	firstChar := string(strIn[0])
	if _, err := strconv.Atoi(firstChar); err == nil {
		return "", ErrInvalidString
	}

	if !isValidString(strIn) {
		return "", ErrInvalidString
	}

	var builder strings.Builder

	return builder.String(), nil
}

func isValidString(str string) bool {
	strLen := len(str) - 1
	for i := range str {
		if _, err := strconv.Atoi(string(str[i])); err != nil {
			continue
		}

		if i == 0 {
			return false
		}

		if strLen > i {
			if _, err := strconv.Atoi(string(str[i+1])); err != nil {
				continue
			}
			return false
		}
	}

	return true
}
