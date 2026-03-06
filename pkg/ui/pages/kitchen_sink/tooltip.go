package kitchensink

import (
	. "github.com/namzug16/gotags"
)

func TooltipSection() HTML {
	return Div(
		X.Class("p-4"),
		Div(
			X.Class("flex flex-wrap items-center gap-4"),
			Button(X.Type("button"), X.Class("btn-outline"), X.Attr("data-tooltip", "Top tooltip"), "Top"),
			Button(X.Type("button"), X.Class("btn-outline"), X.Attr("data-tooltip", "Right tooltip"), X.Attr("data-side", "right"), "Right"),
			Button(X.Type("button"), X.Class("btn-outline"), X.Attr("data-tooltip", "Bottom tooltip"), X.Attr("data-side", "bottom"), "Bottom"),
			Button(X.Type("button"), X.Class("btn-outline"), X.Attr("data-tooltip", "Left tooltip"), X.Attr("data-side", "left"), "Left"),
		),
	)
}
