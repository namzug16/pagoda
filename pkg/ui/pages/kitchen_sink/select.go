package kitchensink

import (
	"github.com/mikestefanello/pagoda/pkg/ui/basecoat"
	"github.com/mikestefanello/pagoda/pkg/ui/lucide"
	. "github.com/namzug16/gotags"
)

func SelectSection() HTML {
	return Div(
		X.Class("p-4"),
		Div(
			X.Class("flex flex-col gap-4"),
			Div(
				X.Class("flex flex-wrap items-center gap-2 md:flex-row"),
				Select(
					X.Class("select w-[180px]"),
					Optgroup(X.Attr("label", "Fruits"), Option("Apple"), Option("Banana"), Option("Blueberry")),
					Optgroup(X.Attr("label", "Grapes"), Option("Pineapple")),
				),
				Select(X.Class("select w-[180px]"), X.Disabled(), Option("Disabled")),
			),
			Div(
				X.Class("flex flex-wrap items-center gap-4"),
				basecoat.Select(basecoat.SelectParams{
					ID:                  "select-default",
					Selected:            []string{"blueberry"},
					TriggerExtraClasses: "w-[180px]",
					Items: []basecoat.SelectItem{
						{Type: "group", Label: "Fruits", Items: []basecoat.SelectItem{{Type: "item", Value: "apple", Label: "Apple"}, {Type: "item", Value: "banana", Label: "Banana"}, {Type: "item", Value: "blueberry", Label: "Blueberry"}}},
						{Type: "group", Label: "Grapes", Items: []basecoat.SelectItem{{Type: "item", Value: "pineapple", Label: "Pineapple"}}},
					},
				}),
				basecoat.Select(basecoat.SelectParams{
					ID:                  "select-scrollbar",
					Selected:            []string{"item-0"},
					TriggerExtraClasses: "w-[180px]",
					ListboxExtraClasses: "scrollbar overflow-y-auto max-h-64",
					Items: []basecoat.SelectItem{
						{Type: "item", Value: "item-0", Label: "Item 0"}, {Type: "item", Value: "item-1", Label: "Item 1"}, {Type: "item", Value: "item-2", Label: "Item 2"}, {Type: "item", Value: "item-3", Label: "Item 3"}, {Type: "item", Value: "item-4", Label: "Item 4"},
						{Type: "item", Value: "item-5", Label: "Item 5"}, {Type: "item", Value: "item-6", Label: "Item 6"}, {Type: "item", Value: "item-7", Label: "Item 7"}, {Type: "item", Value: "item-8", Label: "Item 8"}, {Type: "item", Value: "item-9", Label: "Item 9"},
						{Type: "item", Value: "item-10", Label: "Item 10"}, {Type: "item", Value: "item-11", Label: "Item 11"}, {Type: "item", Value: "item-12", Label: "Item 12"}, {Type: "item", Value: "item-13", Label: "Item 13"}, {Type: "item", Value: "item-14", Label: "Item 14"},
						{Type: "item", Value: "item-15", Label: "Item 15"}, {Type: "item", Value: "item-16", Label: "Item 16"}, {Type: "item", Value: "item-17", Label: "Item 17"}, {Type: "item", Value: "item-18", Label: "Item 18"}, {Type: "item", Value: "item-19", Label: "Item 19"},
					},
				}),
				basecoat.Select(basecoat.SelectParams{
					ID:                  "select-disabled",
					Selected:            []string{"disabled"},
					TriggerExtraClasses: "w-[180px]",
					TriggerAttrs:        []HTML{X.Attr("disabled", "disabled")},
					Items:               []basecoat.SelectItem{{Type: "item", Value: "disabled", Label: "Disabled"}},
				}),
				basecoat.Select(basecoat.SelectParams{
					ID:                  "select-with-icon",
					Selected:            []string{"bar"},
					Name:                "chart-type",
					TriggerExtraClasses: "w-[180px]",
					Items: []basecoat.SelectItem{
						{Type: "item", Value: "bar", Label: Span(X.Class("flex items-center gap-2"), lucide.ChartColumn(X.Class("text-muted-foreground")), "Bar")},
						{Type: "item", Value: "line", Label: Span(X.Class("flex items-center gap-2"), lucide.ChartLine(X.Class("text-muted-foreground")), "Line")},
						{Type: "item", Value: "pie", Label: Span(X.Class("flex items-center gap-2"), lucide.ChartPie(X.Class("text-muted-foreground")), "Pie")},
					},
				}),
			),
		),
	)
}
