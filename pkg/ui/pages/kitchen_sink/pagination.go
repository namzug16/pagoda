package kitchensink

import (
	"github.com/mikestefanello/pagoda/pkg/ui/lucide"
	. "github.com/namzug16/gotags"
)

func PaginationSection() HTML {
	return Div(
		X.Class("p-4"),
		Div(
			X.Class("inline-flex"),
			Nav(
				X.Attr("role", "navigation"),
				X.Attr("aria-label", "pagination"),
				X.Class("mx-auto flex w-full justify-center"),
				Ul(
					X.Class("flex flex-row items-center gap-1"),
					Li(A(X.Href("#"), X.Class("btn-ghost"), lucide.ChevronLeft(), " Previous")),
					Li(A(X.Href("#"), X.Class("btn-ghost size-9"), "1")),
					Li(A(X.Href("#"), X.Class("btn-outline size-9"), "2")),
					Li(A(X.Href("#"), X.Class("btn-ghost size-9"), "3")),
					Li(A(X.Href("#"), X.Class("btn-icon-ghost"), lucide.Ellipsis())),
					Li(A(X.Href("#"), X.Class("btn-ghost"), "Next ", lucide.ChevronRight())),
				),
			),
		),
	)
}
