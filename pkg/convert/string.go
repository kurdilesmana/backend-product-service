package convert

import (
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

var (
	matchFirstCap = regexp.MustCompile("(.)([A-Z][a-z]+)")
	matchAllCap   = regexp.MustCompile("([a-z0-9])([A-Z])")
)

func NormalizePhoneNumber(phoneNumber string) string {
	strNo := ""
	for _, c := range phoneNumber {
		if unicode.IsDigit(c) {
			strNo += string(c)
		}
	}

	if string(strNo[0]) == "0" {
		strNo = fmt.Sprintf("62%s", strNo[1:])
	}

	return strNo
}

func ToSnakeCase(str string) string {
	snake := matchFirstCap.ReplaceAllString(str, "${1}_${2}")
	snake = matchAllCap.ReplaceAllString(snake, "${1}_${2}")
	return strings.ToLower(snake)
}
