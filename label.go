package goplot

import "github.com/Skrip42/go-plot/internal/helpers"

type labelOption struct {
	compile func() string
}

func (o labelOption) Compile() string {
	return o.compile()
}

func (b *builder) SetXLabel(name string, options ...labelOption) {
	b.commands["xlabel"] = command{
		compile: func() string {
			return "set xlabel " + helpers.EscapeString(name) + " " + helpers.Compile(options)
		},
	}
}

func (b *builder) SetYLabel(name string, options ...labelOption) {
	b.commands["ylabel"] = command{
		compile: func() string {
			return "set ylabel " + helpers.EscapeString(name) + " " + helpers.Compile(options)
		},
	}
}

func (b *builder) SetZLabel(name string, options ...labelOption) {
	b.commands["zlabel"] = command{
		compile: func() string {
			return "set zlabel " + helpers.EscapeString(name) + " " + helpers.Compile(options)
		},
	}
}

func WithLabelColor(color colorspec) labelOption {
	return labelOption{
		compile: func() string {
			return color.compile()
		},
	}
}
