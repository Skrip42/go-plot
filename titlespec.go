package goplot

import "github.com/Skrip42/go-plot/internal/helpers"

type titlespec struct {
	compile func() string
}

func (t titlespec) Compile() string {
	return t.compile()
}

func Title(title string) titlespec {
	return titlespec{
		compile: func() string {
			return "title " + helpers.EscapeString(title)
		},
	}
}

func NoTitle() titlespec {
	return titlespec{
		compile: func() string {
			return "notitle"
		},
	}
}

func TitleColumnheader() titlespec {
	return titlespec{
		compile: func() string {
			return "title columnheader"
		},
	}
}
