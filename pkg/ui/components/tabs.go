package components

import (
	"fmt"
	"math/rand"

	. "github.com/namzug16/gotags"
)

type Tab struct {
	Title, Body string
}

func Tabs(tabs []Tab) HTML {
	g := make([]HTML, 0, len(tabs)*2)
	id := fmt.Sprintf("tabs-%d", rand.Int())

	for i, tab := range tabs {
		g = append(g,
			Input(
				X.Type("radio"),
				X.Name(id),
				X.Class("tab"),
				X.Attr("aria-label", tab.Title),
				If(i == 0, X.Checked()),
			),
			Div(
				X.Class("tab-content bg-base-100 border-base-300 p-6"),
				Raw(tab.Body),
			))
	}

	return Div(
		X.Class("tabs tabs-lift"),
		g,
	)
}
