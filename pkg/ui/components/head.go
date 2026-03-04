package components

import (
	"strings"

	"github.com/mikestefanello/pagoda/pkg/ui"
	. "github.com/namzug16/gotags"
)

func JS() []HTML {
	return []HTML{
		Script(X.Src("https://unpkg.com/htmx.org@2.0.0/dist/htmx.min.js"), X.Defer()),
		Script(X.Src("https://unpkg.com/alpinejs@3.x.x/dist/cdn.min.js"), X.Defer()),
		Script(X.Src("js/basecoat.all.min.js"), X.Defer()),
	}
}

func CSS() []HTML {
	return []HTML{
		Link(
			X.Href(ui.StaticFile("main.css")),
			X.Rel("stylesheet"),
			X.Type("text/css"),
		),
	}
}

func Metatags(r *ui.Request) HTML {
	title := r.Config.App.Name
	if r.Title != "" {
		title += " | " + r.Title
	}

	return Fragment(
		Meta(X.Charset("utf-8")),
		Meta(X.Name("viewport"), X.Content("width=device-width, initial-scale=1")),
		Link(X.Rel("icon"), X.Href(ui.StaticFile("favicon.png"))),
		Title(title),
		If(r.Metatags.Description != "", Meta(X.Name("description"), X.Content(r.Metatags.Description))),
		If(len(r.Metatags.Keywords) > 0, Meta(X.Name("keywords"), X.Content(strings.Join(r.Metatags.Keywords, ", ")))),
	)
}
