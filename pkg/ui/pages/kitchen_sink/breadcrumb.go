package kitchensink

import (
	"github.com/mikestefanello/pagoda/pkg/ui/lucide"
	. "github.com/namzug16/gotags"
)

func BreadcrumbSection() HTML {
	return Div(
		X.Class("p-4"),
		Ol(
			X.Class("text-muted-foreground flex flex-wrap items-center gap-1.5 text-sm break-words sm:gap-2.5"),
			Li(
				X.Class("inline-flex items-center gap-1.5"),
				A(X.Href("#"), X.Class("hover:text-foreground transition-colors"), "Home"),
			),
			Li(lucide.ChevronRight(X.Class("size-3.5"))),
			Li(
				X.Class("inline-flex items-center gap-1.5"),
				Div(
					X.Id("demo-breadcrumb-menu"),
					X.Class("dropdown-menu"),
					X.Attr("data-dropdown-menu-initialized", "true"),
					Button(
						X.Type("button"),
						X.Id("demo-breadcrumb-menu-trigger"),
						X.Attr("aria-haspopup", "menu"),
						X.Attr("aria-controls", "demo-breadcrumb-menu-menu"),
						X.Attr("aria-expanded", "false"),
						X.Class("flex size-9 items-center justify-center h-4 w-4 hover:text-foreground cursor-pointer"),
						lucide.Ellipsis(),
					),
					Div(
						X.Id("demo-breadcrumb-menu-popover"),
						X.Attr("data-popover", ""),
						X.Attr("aria-hidden", "true"),
						X.Class("p-1"),
						Div(
							X.Attr("role", "menu"),
							X.Id("demo-breadcrumb-menu-menu"),
							X.Attr("aria-labelledby", "demo-breadcrumb-menu-trigger"),
							Div(X.Id("demo-breadcrumb-menu-items-1"), X.Attr("role", "menuitem"), "Documentation"),
							Div(X.Id("demo-breadcrumb-menu-items-2"), X.Attr("role", "menuitem"), "Themes"),
							Div(X.Id("demo-breadcrumb-menu-items-3"), X.Attr("role", "menuitem"), "GitHub"),
						),
					),
				),
			),
			Li(lucide.ChevronRight(X.Class("size-3.5"))),
			Li(
				X.Class("inline-flex items-center gap-1.5"),
				A(X.Href("#"), X.Class("hover:text-foreground transition-colors"), "Components"),
			),
			Li(lucide.ChevronRight(X.Class("size-3.5"))),
			Li(
				X.Class("inline-flex items-center gap-1.5"),
				Span(X.Class("text-foreground font-normal"), "Breadcrumb"),
			),
		),
	)
}
