package kitchensink

import (
	. "github.com/namzug16/gotags"
)

func RadioGroupSection() HTML {
	return Div(
		X.Class("p-4"),
		Div(
			X.Class("flex flex-col gap-y-6"),
			Fieldset(
				X.Class("grid gap-3"),
				Label(X.Class("label"), Input(X.Type("radio"), X.Name("demo-radio-group"), X.Value("default"), X.Class("input")), "Default"),
				Label(X.Class("label"), Input(X.Type("radio"), X.Name("demo-radio-group"), X.Value("comfortable"), X.Class("input")), "Comfortable"),
				Label(X.Class("label"), Input(X.Type("radio"), X.Name("demo-radio-group"), X.Value("compact"), X.Class("input")), "Compact"),
			),
			Fieldset(
				X.Class("grid gap-3 max-w-sm"),
				Label(
					X.Class("label gap-3 items-start hover:bg-accent/50 rounded-lg border p-4 has-[input[type='radio']:checked]:border-green-600 has-[input[type='radio']:checked]:bg-green-50 dark:has-[input[type='radio']:checked]:border-green-900 dark:has-[input[type='radio']:checked]:bg-green-950"),
					Input(X.Type("radio"), X.Name("demo-radio-group-styled"), X.Value("starter"), X.Class("input checked:bg-green-600 checked:border-green-600 dark:checked:bg-input/30 checked:before:bg-background dark:checked:before:bg-primary")),
					Div(
						X.Class("grid gap-1 font-normal"),
						H2(X.Class("font-medium"), "Starter Plan"),
						P(X.Class("text-muted-foreground leading-snug"), "Perfect for small businesses getting started with our platform"),
					),
				),
				Label(
					X.Class("label gap-3 items-start hover:bg-accent/50 rounded-lg border p-4 has-[input[type='radio']:checked]:border-green-600 has-[input[type='radio']:checked]:bg-green-50 dark:has-[input[type='radio']:checked]:border-green-900 dark:has-[input[type='radio']:checked]:bg-green-950"),
					Input(X.Type("radio"), X.Name("demo-radio-group-styled"), X.Value("pro"), X.Class("input checked:bg-green-600 checked:border-green-600 dark:checked:bg-input/30 checked:before:bg-background dark:checked:before:bg-primary")),
					Div(
						X.Class("grid gap-1 font-normal"),
						H2(X.Class("font-medium"), "Pro Plan"),
						P(X.Class("text-muted-foreground leading-snug"), "Advanced features for growing businesses with higher demands"),
					),
				),
			),
		),
	)
}
