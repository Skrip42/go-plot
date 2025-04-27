package goplot

import (
	"fmt"

	"github.com/Skrip42/go-plot/internal/helpers"
)

var terminalFormats = map[string]terminal{
	"png": TerminalPng(),
	"svg": TerminalSvg(),
}

type terminal struct {
	compile func() string
}

func (b *builder) SetTerminal(term terminal) {
	b.commands["terminal"] = command{
		compile: func() string {
			return "set terminal " + term.compile()
		},
	}
}

///SVG

type svgOption struct {
	compile func() string
}

func (o svgOption) Compile() string {
	return o.compile()
}

func TerminalSvg(options ...svgOption) terminal {
	return terminal{
		compile: func() string {
			return "svg " + helpers.Compile(options)
		},
	}
}

type SvgSizeModifer string

const (
	SvgSizeFixed   SvgSizeModifer = "fixed"
	SvgSizeDynamic SvgSizeModifer = "dynamic"
)

func WithSvgSize(w, h int, modifer SvgSizeModifer) svgOption {
	return svgOption{
		compile: func() string {
			return fmt.Sprintf("size %d,%d %s", w, h, modifer)
		},
	}
}

/// PNG

type pngOption struct {
	compile func() string
}

func (o pngOption) Compile() string {
	return o.compile()
}

func TerminalPng(options ...pngOption) terminal {
	return terminal{
		compile: func() string {
			return "png " + helpers.Compile(options)
		},
	}
}

func WithPngSize(w, h int) pngOption {
	return pngOption{
		compile: func() string {
			return fmt.Sprintf("size %d,%d", w, h)
		},
	}
}

func WithPngTransparent(isTransparent bool) pngOption {
	return pngOption{
		compile: func() string {
			if isTransparent {
				return "transparent"
			} else {
				return "notransparent"
			}
		},
	}
}
