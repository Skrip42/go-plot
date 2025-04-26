package goplot

import "fmt"

type colorspec struct {
	compile func() string
}

func (c colorspec) Compile() string {
	return c.compile()
}

func RGBColor(r, g, b int) colorspec {
	return colorspec{
		compile: func() string {
			return fmt.Sprintf("textcolor rgb \"#00%02X%02X%02X\"", r, g, b)
		},
	}
}
