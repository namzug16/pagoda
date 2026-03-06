package kitchensink

import (
	"github.com/mikestefanello/pagoda/pkg/ui/lucide"
	. "github.com/namzug16/gotags"
)

func BadgeSection() HTML {
	return Div(
		X.Class("p-4"),
		Div(
			X.Class("flex flex-col gap-2"),
			Div(
				X.Class("flex flex-wrap items-center gap-2 md:flex-row"),
				Span(X.Class("badge"), "Primary"),
				Span(X.Class("badge-secondary"), "Secondary"),
				Span(X.Class("badge-outline"), "Outline"),
				Span(X.Class("badge-destructive"), "Destructive"),
				Span(
					X.Class("badge-outline"),
					lucide.Check(),
					"Badge",
				),
				Span(
					X.Class("badge-destructive"),
					lucide.CircleAlert(),
					"Alert",
				),
				Span(X.Class("badge rounded-full min-w-5 px-1"), "8"),
				Span(X.Class("badge-destructive rounded-full min-w-5 px-1"), "99"),
				Span(X.Class("badge-outline rounded-full min-w-5 px-1 font-mono tabular-nums"), "20+"),
			),
			Div(
				X.Class("flex flex-wrap items-center gap-2 md:flex-row"),
				A(
					X.Href("#"),
					X.Class("badge"),
					"Link",
					lucide.ArrowRight(),
				),
				A(
					X.Href("#"),
					X.Class("badge-secondary"),
					"Link",
					lucide.ArrowRight(),
				),
				A(
					X.Href("#"),
					X.Class("badge-destructive"),
					"Link",
					lucide.ArrowRight(),
				),
				A(
					X.Href("#"),
					X.Class("badge-outline"),
					"Link",
					lucide.ArrowRight(),
				),
			),
		),
	)
}
