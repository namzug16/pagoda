package basecoat

import (
	"fmt"
	"math/rand"

	gt "github.com/namzug16/gotags"
)

type PopoverParams struct {
	ID      string
	Trigger any
	Content gt.HTML

	MainExtraClasses    string
	TriggerExtraClasses string
	PopoverExtraClasses string

	MainAttrs    []gt.HTML
	TriggerAttrs []gt.HTML
	PopoverAttrs []gt.HTML
}

func Popover(params PopoverParams) gt.HTML {
	id := params.ID
	if id == "" {
		id = fmt.Sprintf("popover-%d", rand.Intn(900000)+100000)
	}

	return gt.Div(
		gt.X.Id(id),
		gt.X.Class("popover "+params.MainExtraClasses),
		params.MainAttrs,
		gt.Button(
			gt.X.Id(id+"-trigger"),
			gt.X.Type("button"),
			gt.X.Attr("aria-expanded", "false"),
			gt.X.Attr("aria-controls", id+"-popover"),
			gt.X.Class(params.TriggerExtraClasses),
			params.TriggerAttrs,
			normalizeAnyComponent(params.Trigger),
		),
		gt.Div(
			gt.X.Id(id+"-popover"),
			gt.X.Attr("data-popover", ""),
			gt.X.Attr("aria-hidden", "true"),
			gt.X.Class(params.PopoverExtraClasses),
			params.PopoverAttrs,
			params.Content,
		),
	)
}
