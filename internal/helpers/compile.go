package helpers

import "strings"

type Compiled interface {
	Compile() string
}

func Compile[T Compiled](com []T) string {
	result := make([]string, len(com))
	for i, c := range com {
		result[i] = c.Compile()
	}
	return strings.Join(result, " ")
}
