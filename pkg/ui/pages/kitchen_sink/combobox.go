package kitchensink

import (
	"github.com/mikestefanello/pagoda/pkg/ui/basecoat"
	. "github.com/namzug16/gotags"
)

func ComboboxSection() HTML {
	return Div(
		X.Class("p-4"),
		Div(
			X.Class("flex flex-wrap items-start gap-4"),
			basecoat.Select(basecoat.SelectParams{
				ID:                  "demo-combobox-frameworks",
				IsCombobox:          true,
				Selected:            []string{"Next.js"},
				PopoverExtraClasses: "w-48",
				ListboxAttrs: []HTML{
					X.Attr("data-empty", "No framework found."),
				},
				Items: []basecoat.SelectItem{
					{Type: "item", Value: "Next.js", Label: "Next.js"},
					{Type: "item", Value: "SvelteKit", Label: "SvelteKit"},
					{Type: "item", Value: "Nuxt.js", Label: "Nuxt.js"},
					{Type: "item", Value: "Remix", Label: "Remix"},
					{Type: "item", Value: "Astro", Label: "Astro"},
				},
			}),
			basecoat.Select(basecoat.SelectParams{
				ID:                  "demo-combobox-timezones",
				IsCombobox:          true,
				Selected:            []string{"America/New_York"},
				PopoverExtraClasses: "w-72",
				ListboxExtraClasses: "scrollbar overflow-y-auto max-h-64",
				ListboxAttrs: []HTML{
					X.Data("empty", "No timezone found."),
				},
				Items: []basecoat.SelectItem{
					{
						Type:  "group",
						ID:    "demo-combobox-timezones-group-0",
						Label: "Americas",
						Items: []basecoat.SelectItem{
							{Type: "item", Value: "America/New_York", Label: "(GMT-5) New York"},
							{Type: "item", Value: "America/Los_Angeles", Label: "(GMT-8) Los Angeles"},
							{Type: "item", Value: "America/Chicago", Label: "(GMT-6) Chicago"},
							{Type: "item", Value: "America/Toronto", Label: "(GMT-5) Toronto"},
							{Type: "item", Value: "America/Vancouver", Label: "(GMT-8) Vancouver"},
							{Type: "item", Value: "America/Sao_Paulo", Label: "(GMT-3) Sao Paulo"},
						},
					},
					{
						Type:  "group",
						ID:    "demo-combobox-timezones-group-1",
						Label: "Europe",
						Items: []basecoat.SelectItem{
							{Type: "item", Value: "Europe/London", Label: "(GMT+0) London"},
							{Type: "item", Value: "Europe/Paris", Label: "(GMT+1) Paris"},
							{Type: "item", Value: "Europe/Berlin", Label: "(GMT+1) Berlin"},
							{Type: "item", Value: "Europe/Rome", Label: "(GMT+1) Rome"},
							{Type: "item", Value: "Europe/Madrid", Label: "(GMT+1) Madrid"},
							{Type: "item", Value: "Europe/Amsterdam", Label: "(GMT+1) Amsterdam"},
						},
					},
					{
						Type:  "group",
						ID:    "demo-combobox-timezones-group-2",
						Label: "Asia/Pacific",
						Items: []basecoat.SelectItem{
							{Type: "item", Value: "Asia/Tokyo", Label: "(GMT+9) Tokyo"},
							{Type: "item", Value: "Asia/Shanghai", Label: "(GMT+8) Shanghai"},
							{Type: "item", Value: "Asia/Singapore", Label: "(GMT+8) Singapore"},
							{Type: "item", Value: "Asia/Hong_Kong", Label: "(GMT+8) Hong Kong"},
							{Type: "item", Value: "Asia/Bangkok", Label: "(GMT+7) Bangkok"},
							{Type: "item", Value: "Asia/Jakarta", Label: "(GMT+7) Jakarta"},
							{Type: "item", Value: "Asia/Seoul", Label: "(GMT+9) Seoul"},
							{Type: "item", Value: "Australia/Sydney", Label: "(GMT+10) Sydney"},
							{Type: "item", Value: "Australia/Melbourne", Label: "(GMT+10) Melbourne"},
						},
					},
				},
			}),
		),
	)
}
