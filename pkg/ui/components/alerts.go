package components

import (
	"github.com/mikestefanello/pagoda/pkg/msg"
	"github.com/mikestefanello/pagoda/pkg/ui"
	"github.com/mikestefanello/pagoda/pkg/ui/icons"
	. "github.com/namzug16/gotags"
)

func FlashMessages(r *ui.Request) HTML {
	var g []HTML
	var color Color

	for _, typ := range []msg.Type{
		msg.TypeSuccess,
		msg.TypeInfo,
		msg.TypeWarning,
		msg.TypeError,
	} {
		for _, str := range msg.Get(r.Context, typ) {
			switch typ {
			case msg.TypeSuccess:
				color = ColorSuccess
			case msg.TypeInfo:
				color = ColorInfo
			case msg.TypeWarning:
				color = ColorWarning
			case msg.TypeError:
				color = ColorError
			}

			g = append(g, Alert(color, str))
		}
	}

	return Fragment(g...)
}

func Alert(color Color, text string) HTML {
	var class string

	switch color {
	case ColorSuccess:
		class = "alert-success"
	case ColorInfo:
		class = "alert-info"
	case ColorWarning:
		class = "alert-warning"
	case ColorError:
		class = "alert-error"
	}

	return Div(
		X.Role("alert"),
		X.Class("alert mb-2 "+class),
		X.Attr("x-data", "{show: true}"),
		X.Attr("x-show", "show"),
		Span(
			X.Attr("@click", "show = false"),
			X.Class("cursor-pointer"),
			icons.XCircle(),
		),
		Span(text),
	)
}
