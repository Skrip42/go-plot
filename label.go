package goplot

import "github.com/Skrip42/go-plot/internal/helpers"

type labelOption struct {
	compile func() string
}

func (o labelOption) Compile() string {
	return o.compile()
}

func (b *builder) SetXLabel(name string, options ...labelOption) {
	b.commands = append(
		b.commands,
		command{
			compile: func() string {
				return "set xlabel \"" + name + "\" " + helpers.Compile(options)
			},
		},
	)
}

func WithLabelColor(color colorspec) labelOption {
	return labelOption{
		compile: func() string {
			return color.compile()
		},
	}
}
