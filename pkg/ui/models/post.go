package models

import (
	"fmt"

	"github.com/mikestefanello/pagoda/pkg/pager"
	"github.com/mikestefanello/pagoda/pkg/ui"
	. "github.com/mikestefanello/pagoda/pkg/ui/components"
	. "github.com/namzug16/gotags"
)

type (
	Posts struct {
		Posts []Post
		Pager pager.Pager
	}

	Post struct {
		ID          int
		Title, Body string
	}
)

func (p *Posts) Render(path string) HTML {
	g := make([]HTML, 0)
	for i, post := range p.Posts {
		g = append(g, post.Render())
		if i < len(p.Posts)-1 {
			g = append(g, Separator())
		}
	}

	return Div(
		X.Id("posts"),
		X.Class("grid gap-4"),
		Ul(
			X.Class("flex flex-col gap-3"),
			g,
		),
		Pager(p.Pager.Page, path, !p.Pager.IsEnd(), "#posts"),
	)
}

func Separator() HTML {
	return Div(X.Class("h-px w-full bg-border"))
}

func (p *Post) Render() HTML {
	return Li(
		X.Class("flex flex-row gap-2 items-center"),
		Div(
			X.Class("text-4xl font-thin opacity-30 tabular-nums"),
			fmt.Sprintf("%02d", p.ID),
		),
		Div(
			X.Class("min-w-10 min-h-10"),
			Img(
				X.Class("size-10 rounded-box"),
				X.Src(ui.StaticFile("gopher.png")),
				X.Alt("Gopher"),
			),
		),
		Div(
			X.Class("list-col-grow"),
			Div(
				p.Title,
			),
			Div(
				X.Class("text-xs font-semibold opacity-60"),
				p.Body,
			),
		),
	)
}
