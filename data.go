package goplot

import (
	"fmt"
	"os"
	"strings"
)

const tmpPrefix = "go-gnuplot-"

func (b *builder) AddPoints(name string, points [][]float64) {
	b.data[name] = data{
		compile: func() (string, error) {
			f, err := os.CreateTemp(os.TempDir(), tmpPrefix)
			if err != nil {
				return "", fmt.Errorf("can't create temp file: %w", err)
			}

			patternParts := make([]string, len(points))
			for i := range len(patternParts) {
				patternParts[i] = "%v"
			}
			pattern := strings.Join(patternParts, " ") + "\n"

			for i := range len(points[0]) {
				line := make([]any, len(points))
				for j := range len(points) {
					line[j] = points[j][i]
				}
				f.WriteString(fmt.Sprintf(pattern, line...))
			}
			f.Close()
			return "\"" + f.Name() + "\"", nil
		},
	}
}
