package models

import (
	. "github.com/namzug16/gotags"
)

type SearchResult struct {
	Title string
	URL   string
}

func (s *SearchResult) Render() HTML {
	return Li(
		X.Class("list-row"),
		A(
			X.Href(s.URL),
			s.Title,
		),
	)
}
