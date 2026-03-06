package basecoat

import (
	"fmt"
	"math/rand"
	"strconv"

	gt "github.com/namzug16/gotags"

	"github.com/mikestefanello/pagoda/pkg/ui/lucide"
)

type CommandItem struct {
	Type     string
	ID       string
	Label    any
	Icon     any
	URL      string
	Keywords string
	Disabled bool
	Items    []CommandItem
	Attrs    []gt.HTML
}

type CommandParams struct {
	ID          string
	Items       []CommandItem
	Placeholder string
	EmptyText   string
	Content     gt.HTML

	MainExtraClasses  string
	InputExtraClasses string
	MenuExtraClasses  string

	MainAttrs  []gt.HTML
	InputAttrs []gt.HTML
	MenuAttrs  []gt.HTML
}

type CommandDialogParams struct {
	ID                  string
	Items               []CommandItem
	Placeholder         string
	EmptyText           string
	Content             gt.HTML
	Open                bool
	CloseButton         *bool
	CloseOnOverlayClick *bool

	DialogExtraClasses  string
	CommandExtraClasses string
	InputExtraClasses   string
	MenuExtraClasses    string

	DialogAttrs []gt.HTML
	InputAttrs  []gt.HTML
	MenuAttrs   []gt.HTML
}

func Command(params CommandParams) gt.HTML {
	id := params.ID
	if id == "" {
		id = fmt.Sprintf("command-%d", rand.Intn(900000)+100000)
	}

	placeholder := params.Placeholder
	if placeholder == "" {
		placeholder = "Type a command or search..."
	}

	emptyText := params.EmptyText
	if emptyText == "" {
		emptyText = "No results found."
	}

	menuContent := params.Content
	if len(params.Items) > 0 {
		menuContent = renderCommandItems(params.Items, id+"-items")
	}

	return gt.Div(
		gt.X.Id(id),
		gt.X.Class("command "+params.MainExtraClasses),
		gt.X.Attr("aria-label", "Command menu"),
		params.MainAttrs,
		gt.Header(
			lucide.Search(),
			gt.Input(
				gt.X.Type("text"),
				gt.X.Id(id+"-input"),
				gt.X.Placeholder(placeholder),
				gt.X.Attr("autocomplete", "off"),
				gt.X.Attr("autocorrect", "off"),
				gt.X.Attr("spellcheck", "false"),
				gt.X.Attr("aria-autocomplete", "list"),
				gt.X.Attr("role", "combobox"),
				gt.X.Attr("aria-expanded", "true"),
				gt.X.Attr("aria-controls", id+"-menu"),
				gt.X.Class(params.InputExtraClasses),
				params.InputAttrs,
			),
		),
		gt.Div(
			gt.X.Attr("role", "menu"),
			gt.X.Id(id+"-menu"),
			gt.X.Attr("aria-orientation", "vertical"),
			gt.X.Attr("data-empty", emptyText),
			gt.X.Class("scrollbar "+params.MenuExtraClasses),
			params.MenuAttrs,
			menuContent,
		),
	)
}

func CommandDialog(params CommandDialogParams) gt.HTML {
	closeButton := true
	if params.CloseButton != nil {
		closeButton = *params.CloseButton
	}

	closeOnOverlayClick := true
	if params.CloseOnOverlayClick != nil {
		closeOnOverlayClick = *params.CloseOnOverlayClick
	}

	id := params.ID
	if id == "" {
		id = fmt.Sprintf("command-dialog-%d", rand.Intn(900000)+100000)
	}

	placeholder := params.Placeholder
	if placeholder == "" {
		placeholder = "Type a command or search..."
	}

	emptyText := params.EmptyText
	if emptyText == "" {
		emptyText = "No results found."
	}

	menuContent := params.Content
	if len(params.Items) > 0 {
		menuContent = renderCommandItems(params.Items, id+"-items")
	}

	return gt.Dialog(
		gt.X.Id(id),
		gt.X.Class("dialog command-dialog "+params.DialogExtraClasses),
		gt.X.Attr("aria-label", "Command menu"),
		gt.If(params.Open, gt.X.Attr("open", "")),
		gt.If(closeOnOverlayClick, gt.X.Attr("onclick", "if (event.target === this) this.close()")),
		params.DialogAttrs,
		gt.Div(
			gt.X.Class("command "+params.CommandExtraClasses),
			gt.Header(
				lucide.Search(),
				gt.Input(
					gt.X.Type("text"),
					gt.X.Id(id+"-input"),
					gt.X.Placeholder(placeholder),
					gt.X.Attr("autocomplete", "off"),
					gt.X.Attr("autocorrect", "off"),
					gt.X.Attr("spellcheck", "false"),
					gt.X.Attr("aria-autocomplete", "list"),
					gt.X.Attr("role", "combobox"),
					gt.X.Attr("aria-expanded", "true"),
					gt.X.Attr("aria-controls", id+"-menu"),
					gt.X.Class(params.InputExtraClasses),
					params.InputAttrs,
				),
			),
			gt.Div(
				gt.X.Attr("role", "menu"),
				gt.X.Id(id+"-menu"),
				gt.X.Attr("aria-orientation", "vertical"),
				gt.X.Attr("data-empty", emptyText),
				gt.X.Class("scrollbar "+params.MenuExtraClasses),
				params.MenuAttrs,
				menuContent,
			),
			gt.If(
				closeButton,
				gt.Button(
					gt.X.Type("button"),
					gt.X.Attr("aria-label", "Close dialog"),
					gt.X.Attr("onclick", "this.closest('dialog').close()"),
					lucide.X(),
				),
			),
		),
	)
}

func renderCommandItems(items []CommandItem, parentIDPrefix string) gt.HTML {
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
				groupItems = renderCommandItems(item.Items, itemID)
			}

			entries = append(entries, gt.Div(
				gt.X.Attr("role", "group"),
				gt.X.Attr("aria-labelledby", groupLabelID),
				item.Attrs,
				gt.Span(
					gt.X.Attr("role", "heading"),
					gt.X.Id(groupLabelID),
					item.Label,
				),
				groupItems,
			))
		case "separator":
			entries = append(entries, gt.Hr(gt.X.Attr("role", "separator")))
		default:
			if item.URL != "" {
				entries = append(entries, gt.A(
					gt.X.Id(itemID),
					gt.X.Href(item.URL),
					gt.X.Attr("role", "menuitem"),
					gt.If(item.Keywords != "", gt.X.Attr("data-keywords", item.Keywords)),
					gt.If(item.Disabled, gt.X.Attr("aria-disabled", "true")),
					item.Attrs,
					item.Icon,
					item.Label,
				))
				continue
			}

			entries = append(entries, gt.Div(
				gt.X.Id(itemID),
				gt.X.Attr("role", "menuitem"),
				gt.If(item.Keywords != "", gt.X.Attr("data-keywords", item.Keywords)),
				gt.If(item.Disabled, gt.X.Attr("aria-disabled", "true")),
				item.Attrs,
				item.Icon,
				item.Label,
			))
		}
	}

	return gt.Fragment(entries...)
}
