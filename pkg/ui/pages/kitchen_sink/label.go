package kitchensink

import (
	. "github.com/namzug16/gotags"
)

func LabelSection() HTML {
	return Div(
		X.Class("p-4"),
		Div(
			X.Class("flex flex-col gap-y-4"),
			Div(
				X.Class("grid w-full max-w-sm gap-6"),
				Div(
					X.Class("flex items-center gap-3"),
					Input(X.Type("checkbox"), X.Id("label-demo-terms"), X.Class("input")),
					Label(X.For("label-demo-terms"), X.Class("label"), "Accept terms and conditions"),
				),
				Div(
					X.Class("grid gap-3"),
					Label(X.For("label-demo-text"), X.Class("label"), "Username"),
					Input(X.Type("text"), X.Id("label-demo-text"), X.Class("input"), X.Placeholder("Username")),
				),
				Div(
					X.Class("grid gap-3"),
					Label(X.For("label-demo-disabled"), X.Class("label"), "Disabled"),
					Input(X.Type("text"), X.Id("label-demo-disabled"), X.Class("peer input"), X.Placeholder("Disabled"), X.Disabled()),
				),
				Div(
					X.Class("grid gap-3"),
					Label(X.For("label-demo-textarea"), X.Class("label"), "Message"),
					Textarea(X.Id("label-demo-textarea"), X.Class("textarea"), X.Placeholder("Message")),
				),
			),
		),
	)
}
