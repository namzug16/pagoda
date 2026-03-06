package basecoat

import (
	"fmt"
	"math/rand"
	"strconv"

	gt "github.com/namzug16/gotags"
)

type TabsTabset struct {
	Tab        any
	Panel      any
	TabAttrs   []gt.HTML
	PanelAttrs []gt.HTML
}

type TabsParams struct {
	ID              string
	Tabsets         []TabsTabset
	DefaultTabIndex int

	MainExtraClasses    string
	TablistExtraClasses string

	MainAttrs    []gt.HTML
	TablistAttrs []gt.HTML
}

func Tabs(params TabsParams) gt.HTML {
	id := params.ID
	if id == "" {
		id = fmt.Sprintf("tabs-%d", rand.Intn(900000)+100000)
	}

	defaultTabIndex := params.DefaultTabIndex
	if defaultTabIndex <= 0 {
		defaultTabIndex = 1
	}

	tabs := make([]any, 0, len(params.Tabsets))
	panels := make([]any, 0, len(params.Tabsets))

	for i, tabset := range params.Tabsets {
		index := i + 1
		indexStr := strconv.Itoa(index)
		tabID := id + "-tab-" + indexStr
		panelID := id + "-panel-" + indexStr
		isActive := index == defaultTabIndex

		tabs = append(tabs, gt.Button(
			gt.X.Type("button"),
			gt.X.Attr("role", "tab"),
			gt.X.Id(tabID),
			gt.X.Attr("aria-controls", panelID),
			gt.X.Attr("aria-selected", strconv.FormatBool(isActive)),
			gt.X.Attr("tabindex", "0"),
			tabset.TabAttrs,
			normalizeAnyComponent(tabset.Tab),
		))

		if tabset.Panel == nil {
			continue
		}

		panels = append(panels, gt.Div(
			gt.X.Attr("role", "tabpanel"),
			gt.X.Id(panelID),
			gt.X.Attr("aria-labelledby", tabID),
			gt.X.Attr("tabindex", "-1"),
			gt.X.Attr("aria-selected", strconv.FormatBool(isActive)),
			gt.If(!isActive, gt.X.Attr("hidden", "")),
			tabset.PanelAttrs,
			normalizeAnyComponent(tabset.Panel),
		))
	}

	return gt.Div(
		gt.X.Class("tabs "+params.MainExtraClasses),
		gt.X.Id(id),
		params.MainAttrs,
		gt.Nav(
			gt.X.Attr("role", "tablist"),
			gt.X.Attr("aria-orientation", "horizontal"),
			gt.X.Class(params.TablistExtraClasses),
			params.TablistAttrs,
			gt.Fragment(tabs...),
		),
		gt.Fragment(panels...),
	)
}
