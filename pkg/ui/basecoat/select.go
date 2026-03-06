package basecoat

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"strings"

	gt "github.com/namzug16/gotags"

	"github.com/mikestefanello/pagoda/pkg/ui/lucide"
)

type SelectItem struct {
	Type  string
	ID    string
	Value string
	Label any
	Items []SelectItem
	Attrs []gt.HTML
}

type SelectParams struct {
	ID                string
	Selected          []string
	Name              string
	Items             []SelectItem
	Multiple          bool
	Placeholder       string
	CloseOnSelect     bool
	SearchPlaceholder string
	IsCombobox        bool
	Content           gt.HTML

	MainExtraClasses    string
	TriggerExtraClasses string
	PopoverExtraClasses string
	ListboxExtraClasses string

	MainAttrs    []gt.HTML
	TriggerAttrs []gt.HTML
	PopoverAttrs []gt.HTML
	ListboxAttrs []gt.HTML
	InputAttrs   []gt.HTML
}

func Select(params SelectParams) gt.HTML {
	id := params.ID
	if id == "" {
		id = fmt.Sprintf("select-%d", rand.Intn(900000)+100000)
	}

	searchPlaceholder := params.SearchPlaceholder
	if searchPlaceholder == "" {
		searchPlaceholder = "Search entries..."
	}

	selectedSet := make(map[string]struct{}, len(params.Selected))
	for _, value := range params.Selected {
		selectedSet[value] = struct{}{}
	}

	firstOption, selectedOptions := collectSelectOptions(params.Items, selectedSet)
	defaultOption := firstOption
	if len(selectedOptions) > 0 {
		defaultOption = selectedOptions[0]
	}

	defaultLabel := ""
	if params.Multiple {
		if len(selectedOptions) > 0 {
			labels := make([]string, 0, len(selectedOptions))
			for _, option := range selectedOptions {
				labels = append(labels, optionLabelText(*option))
			}
			defaultLabel = strings.Join(labels, ", ")
		} else {
			defaultLabel = params.Placeholder
		}
	} else if defaultOption != nil {
		defaultLabel = optionLabelText(*defaultOption)
	}

	hiddenName := params.Name
	if hiddenName == "" {
		hiddenName = id + "-value"
	}

	hiddenValue := ""
	if params.Multiple {
		encoded, err := json.Marshal(params.Selected)
		if err == nil {
			hiddenValue = string(encoded)
		}
	} else if defaultOption != nil {
		hiddenValue = defaultOption.Value
	}

	listContent := params.Content
	if len(params.Items) > 0 {
		listContent = renderSelectItems(params.Items, selectedSet, id+"-items")
	}

	return gt.Div(
		gt.X.Id(id),
		gt.X.Class("select "+params.MainExtraClasses),
		gt.If(params.Multiple && params.Placeholder != "", gt.X.Attr("data-placeholder", params.Placeholder)),
		gt.If(params.Multiple && params.CloseOnSelect, gt.X.Attr("data-close-on-select", "true")),
		params.MainAttrs,
		gt.Button(
			gt.X.Type("button"),
			gt.X.Class("btn-outline "+params.TriggerExtraClasses),
			gt.X.Id(id+"-trigger"),
			gt.X.Attr("aria-haspopup", "listbox"),
			gt.X.Attr("aria-expanded", "false"),
			gt.X.Attr("aria-controls", id+"-listbox"),
			params.TriggerAttrs,
			gt.Span(gt.X.Class("truncate"), defaultLabel),
			gt.If(
				params.IsCombobox,
				lucide.ChevronsUpDown(gt.X.Class("text-muted-foreground opacity-50 shrink-0")),
			),
			gt.If(
				!params.IsCombobox,
				lucide.ChevronDown(gt.X.Class("text-muted-foreground opacity-50 shrink-0")),
			),
		),
		gt.Div(
			gt.X.Id(id+"-popover"),
			gt.X.Attr("data-popover", ""),
			gt.X.Attr("aria-hidden", "true"),
			gt.X.Class(params.PopoverExtraClasses),
			params.PopoverAttrs,
			gt.If(
				params.IsCombobox,
				gt.Header(
					lucide.Search(),
					gt.Input(
						gt.X.Type("text"),
						gt.X.Value(""),
						gt.X.Placeholder(searchPlaceholder),
						gt.X.Attr("autocomplete", "off"),
						gt.X.Attr("autocorrect", "off"),
						gt.X.Attr("spellcheck", "false"),
						gt.X.Attr("aria-autocomplete", "list"),
						gt.X.Attr("role", "combobox"),
						gt.X.Attr("aria-expanded", "false"),
						gt.X.Attr("aria-controls", id+"-listbox"),
						gt.X.Attr("aria-labelledby", id+"-trigger"),
					),
				),
			),
			gt.Div(
				gt.X.Attr("role", "listbox"),
				gt.X.Id(id+"-listbox"),
				gt.X.Attr("aria-orientation", "vertical"),
				gt.X.Attr("aria-labelledby", id+"-trigger"),
				gt.If(params.Multiple, gt.X.Attr("aria-multiselectable", "true")),
				gt.X.Class(params.ListboxExtraClasses),
				params.ListboxAttrs,
				listContent,
			),
		),
		gt.Input(
			gt.X.Type("hidden"),
			gt.X.Name(hiddenName),
			gt.X.Value(hiddenValue),
			params.InputAttrs,
		),
	)
}

func collectSelectOptions(items []SelectItem, selectedSet map[string]struct{}) (*SelectItem, []*SelectItem) {
	var firstOption *SelectItem
	selectedOptions := make([]*SelectItem, 0)

	for i := range items {
		item := &items[i]
		if item.Type == "group" {
			groupFirst, groupSelected := collectSelectOptions(item.Items, selectedSet)
			if firstOption == nil {
				firstOption = groupFirst
			}
			selectedOptions = append(selectedOptions, groupSelected...)
			continue
		}

		if item.Type == "separator" {
			continue
		}

		if firstOption == nil {
			firstOption = item
		}

		if _, ok := selectedSet[item.Value]; ok {
			selectedOptions = append(selectedOptions, item)
		}
	}

	return firstOption, selectedOptions
}

func renderSelectItems(items []SelectItem, selectedSet map[string]struct{}, parentIDPrefix string) gt.HTML {
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
				groupItems = renderSelectItems(item.Items, selectedSet, itemID)
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
			entries = append(entries, gt.Div(
				gt.X.Id(itemID),
				gt.X.Attr("role", "option"),
				gt.If(item.Value != "", gt.X.Attr("data-value", item.Value)),
				gt.If(isSelected(selectedSet, item.Value), gt.X.Attr("aria-selected", "true")),
				item.Attrs,
				normalizeAnyComponent(item.Label),
			))
		}
	}

	return gt.Fragment(entries...)
}

func optionLabelText(item SelectItem) string {
	if label, ok := item.Label.(string); ok {
		return label
	}

	if item.Label == nil {
		return ""
	}

	return fmt.Sprint(item.Label)
}

func isSelected(selectedSet map[string]struct{}, value string) bool {
	_, ok := selectedSet[value]
	return ok
}
