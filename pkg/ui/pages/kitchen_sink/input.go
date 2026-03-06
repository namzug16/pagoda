package kitchensink

import (
	. "github.com/namzug16/gotags"
)

func InputSection() HTML {
	return Div(
		X.Class("p-4"),
		Div(
			X.Class("flex flex-col gap-y-4"),
			Input(X.Type("text"), X.Class("input"), X.Placeholder("Text")),
			Input(X.Type("text"), X.Class("input"), X.Disabled(), X.Placeholder("Disabled")),
			Input(X.Type("text"), X.Class("input"), X.Attr("aria-invalid", "true"), X.Placeholder("Error")),
			Input(X.Type("email"), X.Class("input"), X.Placeholder("Email")),
			Input(X.Type("password"), X.Class("input"), X.Placeholder("Password")),
			Input(X.Type("number"), X.Class("input"), X.Placeholder("Number")),
			Input(X.Type("file"), X.Class("input")),
			Input(X.Type("tel"), X.Class("input"), X.Placeholder("Tel")),
			Input(X.Type("url"), X.Class("input"), X.Placeholder("URL")),
			Input(X.Type("search"), X.Class("input"), X.Placeholder("Search")),
			Input(X.Type("date"), X.Class("input")),
			Input(X.Type("datetime-local"), X.Class("input")),
			Input(X.Type("month"), X.Class("input")),
			Input(X.Type("week"), X.Class("input")),
			Input(X.Type("time"), X.Class("input")),
		),
	)
}
