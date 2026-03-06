package basecoat

import (
	"fmt"
	"strconv"

	gt "github.com/namzug16/gotags"
)

type SidebarItem struct {
	Type    string
	ID      string
	Label   any
	Icon    gt.HTML
	URL     string
	Current bool
	Open    bool
	Items   []SidebarItem
	Attrs   []gt.HTML
}

type SidebarParams struct {
	ID      string
	Label   string
	Open    *bool
	Side    string
	Header  any
	Footer  any
	Menu    []SidebarItem
	Content gt.HTML

	MainExtraClasses    string
	HeaderExtraClasses  string
	ContentExtraClasses string
	FooterExtraClasses  string

	MainAttrs    []gt.HTML
	HeaderAttrs  []gt.HTML
	ContentAttrs []gt.HTML
	FooterAttrs  []gt.HTML
}

func Sidebar(params SidebarParams) gt.HTML {
	open := true
	if params.Open != nil {
		open = *params.Open
	}

	label := params.Label
	if label == "" {
		label = "Sidebar navigation"
	}

	side := params.Side
	if side == "" {
		side = "left"
	}

	content := params.Content
	if len(params.Menu) > 0 {
		prefix := "content"
		if params.ID != "" {
			prefix = params.ID + "-content"
		}
		content = renderSidebarContent(params.Menu, prefix)
	}

	return gt.Aside(
		gt.If(params.ID != "", gt.X.Id(params.ID)),
		gt.X.Class("sidebar "+params.MainExtraClasses),
		gt.X.Attr("data-side", side),
		gt.X.Attr("aria-hidden", strconv.FormatBool(!open)),
		gt.If(!open, gt.X.Attr("inert", "")),
		params.MainAttrs,
		gt.Nav(
			gt.X.Attr("aria-label", label),
			gt.If(
				params.Header != nil,
				gt.Header(
					gt.X.Class(params.HeaderExtraClasses),
					params.HeaderAttrs,
					normalizeAnyComponent(params.Header),
				),
			),
			gt.Section(
				gt.X.Class(params.ContentExtraClasses),
				params.ContentAttrs,
				content,
			),
			gt.If(
				params.Footer != nil,
				gt.Footer(
					gt.X.Class(params.FooterExtraClasses),
					params.FooterAttrs,
					normalizeAnyComponent(params.Footer),
				),
			),
		),
	)
}

func renderSidebarContent(items []SidebarItem, parentIDPrefix string) gt.HTML {
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
				groupItems = renderSidebarContent(item.Items, itemID)
			}

			entries = append(entries, gt.Div(
				gt.X.Attr("role", "group"),
				gt.If(hasSidebarLabel(item.Label), gt.X.Attr("aria-labelledby", groupLabelID)),
				item.Attrs,
				gt.If(
					hasSidebarLabel(item.Label),
					gt.H3(
						gt.X.Id(groupLabelID),
						normalizeAnyComponent(item.Label),
					),
				),
				gt.Ul(groupItems),
			))
		case "separator":
			entries = append(entries, gt.Hr(gt.X.Attr("role", "separator")))
		case "submenu":
			submenuID := "submenu-" + itemID
			submenuContentID := submenuID + "-content"

			submenuItems := gt.HTML(nil)
			if len(item.Items) > 0 {
				submenuItems = renderSidebarContent(item.Items, itemID)
			}

			entries = append(entries, gt.Li(
				gt.Details(
					gt.X.Id(submenuID),
					gt.If(item.Open, gt.X.Attr("open", "")),
					item.Attrs,
					gt.Summary(
						gt.X.Attr("aria-controls", submenuContentID),
						normalizeAnyComponent(item.Icon),
						normalizeAnyComponent(item.Label),
					),
					gt.Ul(
						gt.X.Id(submenuContentID),
						submenuItems,
					),
				),
			))
		default:
			entries = append(entries, gt.Li(
				gt.A(
					gt.If(item.URL != "", gt.X.Href(item.URL)),
					gt.If(item.Current, gt.X.Attr("aria-current", "page")),
					item.Attrs,
					normalizeAnyComponent(item.Icon),
					gt.Span(normalizeAnyComponent(item.Label)),
				),
			))
		}
	}

	return gt.Fragment(entries...)
}

func hasSidebarLabel(label any) bool {
	if label == nil {
		return false
	}

	return fmt.Sprint(label) != ""
}
