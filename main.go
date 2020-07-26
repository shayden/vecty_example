package main

import (
	"strconv"
	component "vectyexample/components"

	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
	"github.com/gopherjs/vecty/event"
	"github.com/gopherjs/vecty/prop"
)

func main() {
	vecty.SetTitle("Hello Vecty!")
	vecty.AddStylesheet("//styles.css")
	vecty.AddStylesheet("https://stackpath.bootstrapcdn.com/bootstrap/4.5.0/css/bootstrap.min.css")
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
	container := component.NewContainer(
		component.NewRow(
			component.NewColumn(
				elem.Div(vecty.Text("Hi!")),
			),
			component.NewColumn(),
			component.NewColumn(),
		),
		component.NewRow(
			component.NewColumn(
				vecty.Text("Hello Vecty! Whee"),
			),
			component.NewColumn(
				vecty.Text(strconv.FormatBool(p.Clicked)),
			),
			component.NewColumn(
				elem.Input(
					vecty.Markup(
						vecty.Class("btn"),
						vecty.Class("btn-primary"),
						prop.Type(prop.TypeButton),
						event.Click(
							p.onClicked,
						),
					),
				),
			),
		),
	)
	return elem.Body(
		container,
	)
}
