package generator

import (
	"strings"
)

func lowerCamelCase(target string) string {
	firstString := strings.ToLower(target[:1])
	dropedFirstString := target[1:]
	return firstString + dropedFirstString
}

func upperCamelCase(target string) string {
	firstString := strings.ToUpper(target[:1])
	dropedFirstString := target[1:]
	return firstString + dropedFirstString
}
