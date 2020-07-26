package components

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type Column struct {
	vecty.Core
	children vecty.HTML
}

func NewColumn(markup ...vecty.MarkupOrChild) vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("col-4"),
		),
		elem.Div(markup...),
	)
}
