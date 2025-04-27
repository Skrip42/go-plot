package goplot

import (
	"fmt"
	"os"
	"strings"

	"github.com/Skrip42/go-plot/internal/helpers"
)

const tmpPrefix = "go-gnuplot-"

type dataOption struct {
	compile func() string
}

func (o dataOption) Compile() string {
	return o.compile()
}

func (b *builder) RemoveData(name string) {
	delete(b.data, name)
}

func (b *builder) AddPoints(name string, points [][]float64, options ...dataOption) {
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
			return helpers.EscapeString(f.Name()) + " " + helpers.Compile(options), nil
		},
	}
}

func (b *builder) AddDataFile(name string, filename string, options ...dataOption) {
	b.data[name] = data{
		compile: func() (string, error) {
			return helpers.EscapeString(filename) + " " + helpers.Compile(options), nil
		},
	}
}

func (b *builder) AddFunction(name string, function string, options ...dataOption) {
	b.data[name] = data{
		compile: func() (string, error) {
			return function + " " + helpers.Compile(options), nil
		},
	}
}

func WithDataTitle(title titlespec) dataOption {
	return dataOption{
		compile: func() string {
			return title.compile()
		},
	}
}

func WithDataStyle(style Style) dataOption {
	return dataOption{
		compile: func() string {
			return "with " + string(style)
		},
	}
}
