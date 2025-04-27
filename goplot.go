package goplot

import "context"

type GoPlot interface {
	SetXLabel(name string, options ...labelOption)
	SetYLabel(name string, options ...labelOption)
	SetZLabel(name string, options ...labelOption)
	AddPoints(name string, points [][]float64, options ...dataOption)
	AddDataFile(name string, filename string, options ...dataOption)
	AddFunction(name string, function string, options ...dataOption)
	SetOutput(filename string)
	SetTerminal(term terminal)
	Plot(ctx context.Context) error
	DebugPlot() error
}

type builder struct {
	commands map[string]command
	data     map[string]data
}

func NewGoPlot() GoPlot {
	return &builder{
		commands: map[string]command{},
		data:     map[string]data{},
	}
}

type command struct {
	compile func() string
}

type data struct {
	compile func() (string, error)
}
