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

	if !isValidString(strIn) {
		return "", ErrInvalidString
	}

	var builder strings.Builder
	for i, r := range strIn {
		ch := string(r)
		num, err := strconv.Atoi(ch)
		if err != nil {
			builder.WriteString(ch)
			continue
		}

		if num > 0 {
			builder.WriteString(
				strings.Repeat(string(strIn[i-1]), num-1),
			)
			continue
		}

		if num == 0 {
			newStr := builder.String()[:builder.Len()-1]
			builder.Reset()
			builder.WriteString(newStr)
			continue
		}
	}

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
