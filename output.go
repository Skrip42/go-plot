package goplot

import (
	"strings"

	"github.com/Skrip42/go-plot/internal/helpers"
)

func (b *builder) SetOutput(filename string) {
	// if the terminal is not specified,
	// we try to detect it automatically
	if _, ok := b.commands["terminal"]; !ok {
		founded := false
		for format, term := range terminalFormats {
			if strings.Contains(filename, "."+format) {
				b.SetTerminal(term)
				founded = true
				break
			}
		}
		// if it does not work, we use the png terminal
		if !founded {
			b.SetTerminal(terminalFormats["png"])
		}
	}
	b.commands["output"] = command{
		compile: func() string {
			return "set output " + helpers.EscapeString(filename)
		},
	}
}
