package kitchensink

import (
	. "github.com/namzug16/gotags"
)

func SwitchSection() HTML {
	return Div(
		X.Class("p-4"),
		Div(
			X.Class("inline-flex flex-col gap-y-6"),
			Label(
				X.Class("label"),
				Input(X.Type("checkbox"), X.Name("switch"), X.Attr("role", "switch"), X.Class("input")),
				"Airplane Mode",
			),
			Label(
				X.Class("label"),
				Input(X.Type("checkbox"), X.Name("switch"), X.Attr("role", "switch"), X.Class("input checked:bg-blue-500 dark:checked:bg-blue-600"), X.Checked()),
				"Bluetooth",
			),
			Label(
				X.Class("label gap-6 leading-none border rounded-lg p-4 has-[input[type='checkbox']:checked]:border-blue-600"),
				Div(
					X.Class("grid gap-1"),
					H2(X.Class("font-medium text-sm"), "Share across devices"),
					P(X.Class("text-muted-foreground text-sm"), "Focus is shared across devices, and turns off when you leave the app."),
				),
				Input(X.Type("checkbox"), X.Name("switch"), X.Attr("role", "switch"), X.Class("input checked:bg-blue-500 dark:checked:bg-blue-600")),
			),
		),
	)
}
