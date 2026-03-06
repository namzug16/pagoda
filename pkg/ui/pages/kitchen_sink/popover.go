package kitchensink

import (
	"github.com/mikestefanello/pagoda/pkg/ui/basecoat"
	. "github.com/namzug16/gotags"
)

func PopoverSection() HTML {
	return Div(
		X.Class("p-4"),
		basecoat.Popover(basecoat.PopoverParams{
			ID:                  "demo-popover",
			Trigger:             "Open popover",
			TriggerExtraClasses: "btn-outline",
			PopoverExtraClasses: "w-80",
			Content: Div(
				X.Class("grid gap-4"),
				Header(
					X.Class("grid gap-1.5"),
					H4(X.Class("leading-none font-medium"), "Dimensions"),
					P(X.Class("text-muted-foreground text-sm"), "Set the dimensions for the layer."),
				),
				Form(
					X.Class("form grid gap-2"),
					Div(
						X.Class("grid grid-cols-3 items-center gap-4"),
						Label(X.For("demo-popover-width"), "Width"),
						Input(X.Type("text"), X.Id("demo-popover-width"), X.Value("100%"), X.Class("col-span-2 h-8")),
					),
					Div(
						X.Class("grid grid-cols-3 items-center gap-4"),
						Label(X.For("demo-popover-max-width"), "Max. width"),
						Input(X.Type("text"), X.Id("demo-popover-max-width"), X.Value("300px"), X.Class("col-span-2 h-8")),
					),
					Div(
						X.Class("grid grid-cols-3 items-center gap-4"),
						Label(X.For("demo-popover-height"), "Height"),
						Input(X.Type("text"), X.Id("demo-popover-height"), X.Value("25px"), X.Class("col-span-2 h-8")),
					),
					Div(
						X.Class("grid grid-cols-3 items-center gap-4"),
						Label(X.For("demo-popover-max-height"), "Max. height"),
						Input(X.Type("text"), X.Id("demo-popover-max-height"), X.Value("none"), X.Class("col-span-2 h-8")),
					),
				),
			),
		}),
	)
}
