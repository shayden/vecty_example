package main

import (
	"strconv"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
)

func main() {
	vecty.SetTitle("Hello Vecty!")
	vecty.RenderBody(&PageView{})
}

type PageView struct {
	vecty.Core

	Clicked bool
}

func (p *PageView) onClicked(event *vecty.Event) {
	if p.Clicked {
		println("clicked")
	} else {
		println("not clicked")
	}
	p.Clicked = !p.Clicked

	// Unlike react, you get to figure out when to re-render.
	vecty.Rerender(p)
}

func (p *PageView) Render() vecty.ComponentOrHTML {
	return elem.Body(
		elem.Div(vecty.Text("Hi!")),
		elem.Div(vecty.Text("Hello Vecty! Whee")),
		elem.Div(vecty.Text(strconv.FormatBool(p.Clicked))),
		elem.Div(
			elem.Input(
				vecty.Markup(
					prop.Type(prop.TypeButton),
					event.Click(
						p.onClicked,
					),
				),
			),
		),
	)
}
