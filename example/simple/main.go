package main

import (
	"context"
	"fmt"

	goplot "github.com/Skrip42/go-plot"
)

func main() {
	points := make([][]float64, 2)
	points[0] = make([]float64, 10)
	points[1] = make([]float64, 10)

	for i := range 10 {
		points[0][i] = float64(i)
		points[1][i] = float64(i * 2)
	}

	plot := goplot.NewGoPlot()
	plot.SetXLabel(
		"X label",
		goplot.WithLabelColor(goplot.RGBColor(12, 89, 200)),
	)
	plot.AddPoints(
		"line",
		points,
		goplot.WithDataStyle(goplot.StyleLines),
		goplot.WithDataTitle(goplot.Title("line!")),
	)
	plot.AddFunction("sinus", "sin(x)")
	plot.SetOutput("./example/simple/simple.svg")
	plot.SetTerminal(goplot.TerminalSvg(goplot.WithSvgSize(1000, 1000, goplot.SvgSizeDynamic)))

	err := plot.DebugPlot()
	fmt.Println(err)
	err = plot.Plot(context.Background())
	fmt.Println(err)
}
