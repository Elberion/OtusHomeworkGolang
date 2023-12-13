package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(str string) (string, error) {
	var prevStr, curStr string
	builder := strings.Builder{}
	resultStr := strings.Builder{}
	for _, v := range str {
		builder.WriteRune(v)
		curStr = builder.String()
		if i, err := strconv.Atoi(curStr); err == nil {
			if prevStr != "" {
				resultStr.WriteString(strings.Repeat(prevStr, i))
				prevStr = ""
			} else {
				return "", ErrInvalidString
			}
		} else {
			resultStr.WriteString(prevStr)
			prevStr = curStr
		}
		builder.Reset()
	}
	if _, err := strconv.Atoi(curStr); err != nil {
		resultStr.WriteString(curStr)
	}
	return resultStr.String(), nil
}
