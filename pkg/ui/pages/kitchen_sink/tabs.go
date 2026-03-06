package kitchensink

import (
	"github.com/mikestefanello/pagoda/pkg/ui/basecoat"
	"github.com/mikestefanello/pagoda/pkg/ui/lucide"
	. "github.com/namzug16/gotags"
)

func TabsSection() HTML {
	return Div(
		X.Class("p-4"),
		Div(
			X.Class("flex flex-col gap-6"),
			basecoat.Tabs(basecoat.TabsParams{
				ID:                  "demo-tabs-with-panels",
				MainExtraClasses:    "max-w-[300px]",
				TablistExtraClasses: "w-full",
				Tabsets: []basecoat.TabsTabset{
					{
						Tab: "Account",
						Panel: Div(
							X.Class("card"),
							Header(
								H2("Account"),
								P("Make changes to your account here. Click save when you're done."),
							),
							Section(
								Form(
									X.Class("form grid gap-6"),
									Div(X.Class("grid gap-3"), Label(X.For("demo-tabs-account-name"), "Name"), Input(X.Type("text"), X.Id("demo-tabs-account-name"), X.Value("Pedro Duarte"))),
									Div(X.Class("grid gap-3"), Label(X.For("demo-tabs-account-username"), "Username"), Input(X.Type("text"), X.Id("demo-tabs-account-username"), X.Value("@peduarte"))),
								),
							),
							Footer(Button(X.Type("button"), X.Class("btn"), "Save changes")),
						),
					},
					{
						Tab: "Password",
						Panel: Div(
							X.Class("card"),
							Header(
								H2("Password"),
								P("Change your password here. After saving, you'll be logged out."),
							),
							Section(
								Form(
									X.Class("form grid gap-6"),
									Div(X.Class("grid gap-3"), Label(X.For("demo-tabs-password-current"), "Current password"), Input(X.Type("password"), X.Id("demo-tabs-password-current"))),
									Div(X.Class("grid gap-3"), Label(X.For("demo-tabs-password-new"), "New password"), Input(X.Type("password"), X.Id("demo-tabs-password-new"))),
								),
							),
							Footer(Button(X.Type("button"), X.Class("btn"), "Save Password")),
						),
					},
				},
			}),
			basecoat.Tabs(basecoat.TabsParams{
				ID:      "demo-tabs-without-panels",
				Tabsets: []basecoat.TabsTabset{{Tab: "Home"}, {Tab: "Settings"}},
			}),
			basecoat.Tabs(basecoat.TabsParams{
				ID: "demo-tabs-disabled",
				Tabsets: []basecoat.TabsTabset{
					{Tab: "Home"},
					{Tab: "Disabled", TabAttrs: []HTML{X.Attr("disabled", "disabled")}},
				},
			}),
			basecoat.Tabs(basecoat.TabsParams{
				ID: "demo-tabs-with-icons",
				Tabsets: []basecoat.TabsTabset{
					{Tab: Fragment(lucide.Newspaper(), Span("Preview"))},
					{Tab: Fragment(lucide.Code(), Span("Code"))},
				},
			}),
		),
	)
}
