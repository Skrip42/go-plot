package helpers

import "strings"

func EscapeString(str string) string {
	str = strings.ReplaceAll(str, "\"", "\\\"")
	return "\"" + str + "\""
}
