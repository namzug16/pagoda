package kitchensink

import (
	. "github.com/namzug16/gotags"
)

func TextareaSection() HTML {
	return Div(
		X.Class("p-4"),
		Div(
			X.Class("flex flex-col gap-y-10"),
			Textarea(X.Class("textarea"), X.Placeholder("Type your message here")),
			Textarea(X.Class("textarea"), X.Placeholder("Type your message here"), X.Attr("aria-invalid", "true")),
			Div(
				X.Class("grid gap-3"),
				Label(X.For("textarea-demo-label"), X.Class("label"), "Label"),
				Textarea(X.Id("textarea-demo-label"), X.Class("textarea"), X.Placeholder("Type your message here")),
			),
			Div(
				X.Class("grid gap-3"),
				Label(X.For("textarea-demo-label-and-description"), X.Class("label"), "With label and description"),
				Textarea(X.Id("textarea-demo-label-and-description"), X.Class("textarea"), X.Placeholder("Type your message here")),
				P(X.Class("text-muted-foreground text-sm"), "Type your message and press enter to send."),
			),
			Div(
				X.Class("grid gap-3"),
				Label(X.For("textarea-demo-disabled"), X.Class("label"), "Disabled"),
				Textarea(X.Id("textarea-demo-disabled"), X.Class("textarea"), X.Placeholder("Type your message here"), X.Disabled()),
			),
		),
	)
}
