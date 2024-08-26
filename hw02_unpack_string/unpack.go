package hw02unpackstring

import (
	"errors"
	"regexp"
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

		if num == 0 {
			newStr := builder.String()[:builder.Len()-1]
			builder.Reset()
			builder.WriteString(newStr)
			continue
		}

		builder.WriteString(
			strings.Repeat(string(strIn[i-1]), num-1),
		)
	}

	return builder.String(), nil
}

func isValidString(str string) bool {
	if ok, err := regexp.MatchString("^[A-z][A-z\\d]*$", str); err != nil || !ok {
		return false
	}

	for i := range str {
		curChar := string(str[i])
		_, err := strconv.Atoi(curChar)
		if err == nil && i == 0 {
			return false
		}

		var nextChar string
		if i+1 < len(str) {
			nextChar = string(str[i+1])
			if _, err2 := strconv.Atoi(nextChar); err == nil && err2 == nil {
				return false
			}
		}
	}

	return true
}
