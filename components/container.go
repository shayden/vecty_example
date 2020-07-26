package components

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type Container struct {
	vecty.Core
	children vecty.HTML
}

func NewContainer(markup ...vecty.MarkupOrChild) vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("container"),
		),
		elem.Div(markup...),
	)
}

func (c *Container) Render() vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("container"),
		),
		&c.children,
	)
}
