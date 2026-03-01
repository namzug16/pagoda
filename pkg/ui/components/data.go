package components

import (
	. "github.com/namzug16/gotags"
)

type (
	CardParams struct {
		Title  string
		Body   []HTML
		Footer []HTML
		Color  Color
		Size   Size
	}

	Stat struct {
		Title       string
		Value       string
		Description string
		Icon        HTML
	}
)

func Badge(color Color, text string) HTML {
	var class string

	switch color {
	case ColorSuccess:
		class = "badge-success"
	case ColorWarning:
		class = "badge-warning"
	}

	return Div(
		X.Class("badge "+class),
		text,
	)
}

func Divider(text string) HTML {
	return Div(
		X.Class("divider"),
		text,
	)
}

func Card(params CardParams) HTML {
	var colorClass, sizeClass string

	switch params.Color {
	case ColorSuccess:
		colorClass = "bg-success text-success-content"
	case ColorPrimary:
		colorClass = "bg-primary text-primary-content"
	case ColorAccent:
		colorClass = "bg-accent text-accent-content"
	case ColorNeutral:
		colorClass = "bg-neutral text-neutral-content"
	case ColorWarning:
		colorClass = "bg-warning text-warning-content"
	case ColorInfo:
		colorClass = "bg-info text-info-content"
	}

	switch params.Size {
	case SizeSmall:
		sizeClass = "card-sm"
	case SizeMedium:
		sizeClass = "card-md"
	case SizeLarge:
		sizeClass = "card-lg"
	}

	return Div(
		X.Class("cards mb-2 "+colorClass+" "+sizeClass),
		Div(
			X.Class("card-body"),
			If(len(params.Title) > 0, Span(
				X.Class("card-title"),
				params.Title,
			)),
			params.Body,
			If(len(params.Footer) > 0, Div(
				X.Class("card-actions justify-end"),
				Fragment(params.Footer...),
			)),
		),
	)
}

func Stats(stats ...Stat) HTML {
	g := make([]HTML, 0, len(stats))
	for _, stat := range stats {
		g = append(g, Div(
			X.Class("stat"),
			If(stat.Icon != nil,
				Div(
					X.Class("stat-figure text-secondary"),
					stat.Icon,
				),
			),
			Div(
				X.Class("stat-title"),
				stat.Title,
			),
			Div(
				X.Class("stat-value"),
				stat.Value,
			),
			Div(
				X.Class("stat-desc"),
				stat.Description,
			),
		))
	}
	return Div(
		X.Class("stats shadow"),
		g,
	)
}
