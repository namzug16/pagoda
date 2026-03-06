package kitchensink

import (
	"github.com/mikestefanello/pagoda/pkg/ui/lucide"
	. "github.com/namzug16/gotags"
)

func ButtonSection() HTML {
	return Div(
		X.Class("p-4"),
		Div(
			X.Class("flex flex-col gap-6"),
			Div(
				X.Class("flex flex-wrap items-center gap-2 md:flex-row"),
				Button(X.Type("button"), X.Class("btn-primary"), "Primary"),
				Button(X.Type("button"), X.Class("btn-outline"), "Outline"),
				Button(X.Type("button"), X.Class("btn-ghost"), "Ghost"),
				Button(X.Type("button"), X.Class("btn-destructive"), lucide.Send(), "Danger"),
				Button(X.Type("button"), X.Class("btn-secondary"), "Secondary"),
				Button(X.Type("button"), X.Class("btn-link"), "Link"),
				Button(X.Type("button"), X.Class("btn-outline"), lucide.Send(), "Send"),
				Button(X.Type("button"), X.Class("btn-outline"), "Learn more", lucide.ArrowRight()),
				Button(X.Type("button"), X.Class("btn-outline"), X.Disabled(), lucide.Loader(X.Class("animate-spin")), "Loading"),
			),
			Div(
				X.Class("flex flex-wrap items-center gap-2 md:flex-row"),
				Button(X.Type("button"), X.Class("btn-sm-primary"), "Primary"),
				Button(X.Type("button"), X.Class("btn-sm-outline"), "Outline"),
				Button(X.Type("button"), X.Class("btn-sm-ghost"), "Ghost"),
				Button(X.Type("button"), X.Class("btn-sm-destructive"), "Danger"),
				Button(X.Type("button"), X.Class("btn-sm-secondary"), "Secondary"),
				Button(X.Type("button"), X.Class("btn-sm-link"), "Link"),
				Button(X.Type("button"), X.Class("btn-sm-outline"), lucide.Send(), "Send"),
				Button(X.Type("button"), X.Class("btn-sm-outline"), "Learn more", lucide.ArrowRight()),
				Button(X.Type("button"), X.Class("btn-sm-outline"), X.Disabled(), lucide.Loader(X.Class("animate-spin")), "Loading"),
			),
			Div(
				X.Class("flex flex-wrap items-center gap-2 md:flex-row"),
				Button(X.Type("button"), X.Class("btn-lg-primary"), "Primary"),
				Button(X.Type("button"), X.Class("btn-lg-outline"), "Outline"),
				Button(X.Type("button"), X.Class("btn-lg-ghost"), "Ghost"),
				Button(X.Type("button"), X.Class("btn-lg-destructive"), lucide.Send(), "Danger"),
				Button(X.Type("button"), X.Class("btn-lg-secondary"), "Secondary"),
				Button(X.Type("button"), X.Class("btn-lg-link"), "Link"),
				Button(X.Type("button"), X.Class("btn-lg-outline"), lucide.Send(), "Send"),
				Button(X.Type("button"), X.Class("btn-lg-outline"), "Learn more", lucide.ArrowRight()),
				Button(X.Type("button"), X.Class("btn-lg-outline"), X.Disabled(), lucide.Loader(X.Class("animate-spin")), "Loading"),
			),
			Div(
				X.Class("flex flex-wrap items-center gap-2 md:flex-row"),
				Button(X.Type("button"), X.Class("btn-icon-primary"), lucide.Download()),
				Button(X.Type("button"), X.Class("btn-icon-secondary"), lucide.Upload()),
				Button(X.Type("button"), X.Class("btn-icon-outline"), lucide.ArrowRight()),
				Button(X.Type("button"), X.Class("btn-icon-ghost"), lucide.Ellipsis()),
				Button(X.Type("button"), X.Class("btn-icon-destructive"), lucide.Trash2()),
				Button(X.Type("button"), X.Class("btn-icon-outline"), X.Disabled(), lucide.Loader(X.Class("animate-spin"))),
			),
		),
	)
}
