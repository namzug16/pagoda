package kitchensink

import (
	. "github.com/namzug16/gotags"
)

func FormSection() HTML {
	return Div(
		X.Class("p-4"),
		Form(
			X.Class("form grid w-full max-w-sm gap-6"),
			Div(
				X.Class("grid gap-2"),
				Label(X.For("demo-form-text"), "Username"),
				Input(X.Type("text"), X.Id("demo-form-text"), X.Placeholder("hunvreus")),
				P(X.Class("text-muted-foreground text-sm"), "This is your public display name."),
			),
			Div(
				X.Class("grid gap-2"),
				Label(X.For("demo-form-select"), "Email"),
				Select(
					X.Id("demo-form-select"),
					Option(X.Value("bob@example.com"), "m@example.com"),
					Option(X.Value("alice@example.com"), "m@google.com"),
					Option(X.Value("john@example.com"), "m@support.com"),
				),
				P(X.Class("text-muted-foreground text-sm"), "You can manage email addresses in your email settings."),
			),
			Div(
				X.Class("grid gap-2"),
				Label(X.For("demo-form-textarea"), "Bio"),
				Textarea(X.Id("demo-form-textarea"), X.Placeholder("I like to..."), X.Attr("rows", "3")),
				P(X.Class("text-muted-foreground text-sm"), "You can @mention other users and organizations."),
			),
			Div(
				X.Class("flex flex-col gap-3"),
				Label(X.For("demo-form-radio"), "Notify me about..."),
				Fieldset(
					X.Id("demo-form-radio"),
					X.Class("grid gap-3"),
					Label(
						X.Class("font-normal"),
						Input(X.Type("radio"), X.Name("demo-form-radio"), X.Value("1")),
						"All new messages",
					),
					Label(
						X.Class("font-normal"),
						Input(X.Type("radio"), X.Name("demo-form-radio"), X.Value("2")),
						"Direct messages and mentions",
					),
					Label(
						X.Class("font-normal"),
						Input(X.Type("radio"), X.Name("demo-form-radio"), X.Value("3")),
						"Nothing",
					),
				),
			),
			Div(
				X.Class("flex flex-row items-start gap-3 rounded-md border p-4 shadow-xs"),
				Input(X.Type("checkbox"), X.Id("demo-form-checkbox")),
				Div(
					X.Class("flex flex-col gap-1"),
					Label(X.For("demo-form-checkbox"), X.Class("leading-snug"), "Use different settings for my mobile devices"),
					P(X.Class("text-muted-foreground text-sm leading-snug"), "You can manage your mobile notifications in the mobile settings page."),
				),
			),
			Div(
				X.Class("flex flex-col gap-4"),
				Header(
					Label(X.For("demo-form-checkboxes"), X.Class("text-base leading-normal"), "Sidebar"),
					P(X.Class("text-muted-foreground text-sm"), "Select the items you want to display in the sidebar."),
				),
				Fieldset(
					X.Id("demo-form-checkboxes"),
					X.Class("flex flex-col gap-2"),
					Label(
						X.Class("font-normal leading-tight"),
						Input(X.Type("checkbox"), X.Name("demo-form-checkboxes"), X.Value("1"), X.Checked()),
						"Recents",
					),
					Label(
						X.Class("font-normal leading-tight"),
						Input(X.Type("checkbox"), X.Name("demo-form-checkboxes"), X.Value("2"), X.Checked()),
						"Home",
					),
					Label(
						X.Class("font-normal leading-tight"),
						Input(X.Type("checkbox"), X.Name("demo-form-checkboxes"), X.Value("3")),
						"Applications",
					),
					Label(
						X.Class("font-normal leading-tight"),
						Input(X.Type("checkbox"), X.Name("demo-form-checkboxes"), X.Value("4")),
						"Desktop",
					),
					Label(
						X.Class("font-normal leading-tight"),
						Input(X.Type("checkbox"), X.Name("demo-form-checkboxes"), X.Value("5")),
						"Download",
					),
					Label(
						X.Class("font-normal leading-tight"),
						Input(X.Type("checkbox"), X.Name("demo-form-checkboxes"), X.Value("6")),
						"Documents",
					),
				),
			),
			Div(
				X.Class("grid gap-2"),
				Label(X.For("demo-form-date"), "Date of birth"),
				Input(X.Type("date"), X.Id("demo-form-date")),
				P(X.Class("text-muted-foreground text-sm"), "Your date of birth is used to calculate your age."),
			),
			Section(
				X.Class("grid gap-4"),
				H3(X.Class("text-lg font-medium"), "Email Notifications"),
				Div(
					X.Class("gap-2 flex flex-row items-start justify-between rounded-lg border p-4 shadow-xs"),
					Div(
						X.Class("flex flex-col gap-0.5"),
						Label(X.For("demo-form-switch"), X.Class("leading-normal"), "Marketing emails"),
						P(X.Class("text-muted-foreground text-sm"), "Receive emails about new products, features, and more."),
					),
					Input(X.Type("checkbox"), X.Id("demo-form-switch"), X.Attr("role", "switch")),
				),
				Div(
					X.Class("gap-2 flex flex-row items-start justify-between rounded-lg border p-4 shadow-xs"),
					Div(
						X.Class("flex flex-col gap-0.5 opacity-60"),
						Label(X.For("demo-form-switch-disabled"), X.Class("leading-normal"), "Marketing emails"),
						P(X.Class("text-muted-foreground text-sm"), "Receive emails about new products, features, and more."),
					),
					Input(X.Type("checkbox"), X.Id("demo-form-switch-disabled"), X.Attr("role", "switch"), X.Disabled()),
				),
			),
			Button(X.Type("submit"), X.Class("btn"), "Submit"),
		),
	)
}
