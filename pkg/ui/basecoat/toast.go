package basecoat

import (
	gt "github.com/namzug16/gotags"

	"github.com/mikestefanello/pagoda/pkg/ui/lucide"
)

type ToastAction struct {
	Label   string
	OnClick string
	Href    string
	Attrs   []gt.HTML
}

type ToastParams struct {
	Category    string
	Title       string
	Description string
	Action      *ToastAction
	Cancel      *ToastAction

	ExtraClasses string
	Attrs        []gt.HTML
}

type ToasterParams struct {
	ID string

	Toasts  []ToastParams
	Content gt.HTML

	ExtraClasses string
	Attrs        []gt.HTML
}

func Toaster(params ToasterParams) gt.HTML {
	id := params.ID
	if id == "" {
		id = "toaster"
	}

	content := make([]any, 0, len(params.Toasts)+1)
	for _, toast := range params.Toasts {
		content = append(content, Toast(toast))
	}

	if params.Content != nil {
		content = append(content, params.Content)
	}

	return gt.Div(
		gt.X.Id(id),
		gt.X.Class("toaster "+params.ExtraClasses),
		params.Attrs,
		gt.Fragment(content...),
	)
}

func Toast(params ToastParams) gt.HTML {
	category := params.Category
	if category == "" {
		category = "success"
	}

	role := "status"
	if category == "error" {
		role = "alert"
	}

	footerItems := make([]any, 0, 2)
	if params.Action != nil {
		if params.Action.Href != "" {
			footerItems = append(footerItems, gt.A(
				gt.X.Href(params.Action.Href),
				gt.X.Class("btn-sm"),
				gt.X.Attr("data-toast-action", ""),
				params.Action.Attrs,
				params.Action.Label,
			))
		} else {
			footerItems = append(footerItems, gt.Button(
				gt.X.Type("button"),
				gt.X.Class("btn"),
				gt.X.Attr("data-toast-action", ""),
				gt.If(params.Action.OnClick != "", gt.X.Attr("onclick", params.Action.OnClick)),
				params.Action.Attrs,
				params.Action.Label,
			))
		}
	}

	if params.Cancel != nil {
		footerItems = append(footerItems, gt.Button(
			gt.X.Type("button"),
			gt.X.Class("btn-sm-outline"),
			gt.X.Attr("data-toast-cancel", ""),
			gt.If(params.Cancel.OnClick != "", gt.X.Attr("onclick", params.Cancel.OnClick)),
			params.Cancel.Attrs,
			params.Cancel.Label,
		))
	}

	return gt.Div(
		gt.X.Class("toast "+params.ExtraClasses),
		gt.X.Attr("role", role),
		gt.X.Attr("aria-atomic", "true"),
		gt.X.Attr("aria-hidden", "false"),
		gt.If(category != "", gt.X.Attr("data-category", category)),
		params.Attrs,
		gt.Div(
			gt.X.Class("toast-content"),
			toastCategoryIcon(category),
			gt.Section(
				gt.If(params.Title != "", gt.H2(params.Title)),
				gt.If(params.Description != "", gt.P(params.Description)),
			),
			gt.If(
				params.Action != nil || params.Cancel != nil,
				gt.Footer(
					gt.Fragment(footerItems...),
				),
			),
		),
	)
}

func toastCategoryIcon(category string) gt.HTML {
	switch category {
	case "success":
		return lucide.CircleCheck(gt.X.Attr("aria-hidden", "true"))
	case "error":
		return lucide.CircleAlert(gt.X.Attr("aria-hidden", "true"))
	case "info":
		return lucide.Info(gt.X.Attr("aria-hidden", "true"))
	case "warning":
		return lucide.TriangleAlert(gt.X.Attr("aria-hidden", "true"))
	default:
		return nil
	}
}
