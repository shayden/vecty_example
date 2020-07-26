package components

import (
	"github.com/gopherjs/vecty"
	"github.com/gopherjs/vecty/elem"
)

type Row struct {
	vecty.Core
	children vecty.HTML
}

func NewRow(markup ...vecty.MarkupOrChild) vecty.ComponentOrHTML {
	return elem.Div(
		vecty.Markup(
			vecty.Class("row"),
		),
		elem.Div(markup...),
	)
}
