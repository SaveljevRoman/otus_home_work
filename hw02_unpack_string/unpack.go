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

	if err := validStringNums(strIn); err != nil {
		return "", err
	}

	var builder strings.Builder
	// todo
	return builder.String(), nil
}

func validStringNums(str string) error {
	strLen := len(str) - 1
	for i := range str {
		_, err := strconv.Atoi(string(str[i]))
		if err != nil {
			continue
		}

		if strLen > i {
			_, err = strconv.Atoi(string(str[i+1]))
			if err != nil {
				continue
			}
			return ErrInvalidString
		}
	}

	return nil
}
