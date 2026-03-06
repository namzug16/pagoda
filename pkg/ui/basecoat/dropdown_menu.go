package basecoat

import (
	"fmt"
	"math/rand"
	"strconv"

	gt "github.com/namzug16/gotags"
)

type DropdownMenuItem struct {
	Type  string
	ID    string
	Label any
	Icon  gt.HTML
	URL   string
	Items []DropdownMenuItem
	Attrs []gt.HTML
}

type DropdownMenuParams struct {
	ID      string
	Trigger any
	Items   []DropdownMenuItem
	Content gt.HTML

	MainExtraClasses    string
	TriggerExtraClasses string
	PopoverExtraClasses string
	MenuExtraClasses    string

	MainAttrs    []gt.HTML
	TriggerAttrs []gt.HTML
	PopoverAttrs []gt.HTML
	MenuAttrs    []gt.HTML
}

func DropdownMenu(params DropdownMenuParams) gt.HTML {
	id := params.ID
	if id == "" {
		id = fmt.Sprintf("dropdown-menu-%d", rand.Intn(900000)+100000)
	}

	menuContent := params.Content
	if len(params.Items) > 0 {
		menuContent = renderDropdownMenuItems(params.Items, id+"-items")
	}

	return gt.Div(
		gt.X.Id(id),
		gt.X.Class("dropdown-menu "+params.MainExtraClasses),
		params.MainAttrs,
		gt.Button(
			gt.X.Type("button"),
			gt.X.Id(id+"-trigger"),
			gt.X.Attr("aria-haspopup", "menu"),
			gt.X.Attr("aria-controls", id+"-menu"),
			gt.X.Attr("aria-expanded", "false"),
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
			gt.Div(
				gt.X.Attr("role", "menu"),
				gt.X.Id(id+"-menu"),
				gt.X.Attr("aria-labelledby", id+"-trigger"),
				gt.X.Class(params.MenuExtraClasses),
				params.MenuAttrs,
				menuContent,
			),
		),
	)
}

func renderDropdownMenuItems(items []DropdownMenuItem, parentIDPrefix string) gt.HTML {
	entries := make([]any, 0, len(items))

	for i, item := range items {
		itemID := parentIDPrefix + "-" + strconv.Itoa(i+1)

		switch item.Type {
		case "group":
			groupLabelID := item.ID
			if groupLabelID == "" {
				groupLabelID = "group-label-" + itemID
			}

			groupItems := gt.HTML(nil)
			if len(item.Items) > 0 {
				groupItems = renderDropdownMenuItems(item.Items, itemID)
			}

			entries = append(entries, gt.Div(
				gt.X.Attr("role", "group"),
				gt.X.Attr("aria-labelledby", groupLabelID),
				item.Attrs,
				gt.Div(
					gt.X.Attr("role", "heading"),
					gt.X.Id(groupLabelID),
					normalizeAnyComponent(item.Label),
				),
				groupItems,
			))
		case "separator":
			entries = append(entries, gt.Hr(gt.X.Attr("role", "separator")))
		default:
			if item.URL != "" {
				entries = append(entries, gt.A(
					gt.X.Id(itemID),
					gt.X.Attr("role", "menuitem"),
					gt.X.Href(item.URL),
					item.Attrs,
					normalizeAnyComponent(item.Icon),
					normalizeAnyComponent(item.Label),
				))
				continue
			}

			entries = append(entries, gt.Div(
				gt.X.Id(itemID),
				gt.X.Attr("role", "menuitem"),
				item.Attrs,
				normalizeAnyComponent(item.Icon),
				normalizeAnyComponent(item.Label),
			))
		}
	}

	return gt.Fragment(entries...)
}
