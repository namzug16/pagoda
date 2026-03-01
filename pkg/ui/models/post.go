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
	g := make([]HTML, len(p.Posts))
	for i, post := range p.Posts {
		g[i] = post.Render()
	}

	return Div(
		X.Id("posts"),
		Ul(
			X.Class("list bg-base-100 rounded-box shadow-md not-prose"),
			g,
		),
		Div(X.Class("mb-4")),
		Pager(p.Pager.Page, path, !p.Pager.IsEnd(), "#posts"),
	)
}

func (p *Post) Render() HTML {
	return Li(
		X.Class("list-row"),
		Div(
			X.Class("text-4xl font-thin opacity-30 tabular-nums"),
			fmt.Sprintf("%02d", p.ID),
		),
		Div(
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
