package kitchensink

import (
	. "github.com/namzug16/gotags"
)

func CheckboxSection() HTML {
	return Div(
		X.Class("p-4"),
		Div(
			X.Class("flex flex-col gap-6 max-w-lg"),
			Label(
				X.Class("label gap-3"),
				Input(X.Type("checkbox"), X.Class("input")),
				"Accept terms and conditions",
			),
			Div(
				X.Class("flex items-start gap-3"),
				Input(X.Type("checkbox"), X.Id("demo-checkbox-label-and-description"), X.Class("input")),
				Div(
					X.Class("grid gap-2"),
					Label(X.For("demo-checkbox-label-and-description"), X.Class("label"), "Accept terms and conditions"),
					P(X.Class("text-muted-foreground text-sm"), "By clicking this checkbox, you agree to the terms and conditions."),
				),
			),
			Label(
				X.Class("label gap-3"),
				Input(X.Type("checkbox"), X.Class("input"), X.Disabled()),
				"Enable notifications",
			),
			Label(
				X.Class("flex items-start gap-3 border p-3 hover:bg-accent/50 rounded-lg has-[input[type='checkbox']:checked]:border-blue-600 has-[input[type='checkbox']:checked]:bg-blue-50 dark:has-[input[type='checkbox']:checked]:border-blue-900 dark:has-[input[type='checkbox']:checked]:bg-blue-950"),
				Input(X.Type("checkbox"), X.Class("input checked:bg-blue-600 checked:border-blue-600 dark:checked:bg-blue-700 dark:checked:border-blue-700 checked:after:bg-white"), X.Checked()),
				Div(
					X.Class("grid gap-2"),
					H2(X.Class("text-sm leading-none font-medium"), "Enable notifications"),
					P(X.Class("text-muted-foreground text-sm"), "You can enable or disable notifications at any time."),
				),
			),
		),
	)
}
