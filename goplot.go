package goplot

import "context"

type GoPlot interface {
	SetXLabel(name string, options ...labelOption)
	AddPoints(name string, points [][]float64)
	Plot(ctx context.Context) error
}

type builder struct {
	commands []command
	data     map[string]data
}

func NewGoPlot() GoPlot {
	return &builder{
		commands: []command{},
		data:     map[string]data{},
	}
}

type command struct {
	compile func() string
}

type data struct {
	compile func() (string, error)
}
