package basecoat

import (
	"fmt"
	"math/rand"

	gt "github.com/namzug16/gotags"

	"github.com/mikestefanello/pagoda/pkg/ui/lucide"
)

type DialogParams struct {
	ID                  string
	DialogExtraClasses  string
	Trigger             any // HTML or string
	Title               string
	Description         string
	Body                gt.HTML
	Footer              gt.HTML
	DialogAttrs         []gt.HTML
	TriggerAttrs        []gt.HTML
	HeaderAttrs         []gt.HTML
	Open                bool
	CloseButton         *bool
	CloseOnOverlayClick *bool
}

func Dialog(params DialogParams) gt.HTML {
	// Default CloseButton to true if not set
	closeButton := true
	if params.CloseButton != nil {
		closeButton = *params.CloseButton
	}

	// Default CloseOnOverlayClick to true if not set
	closeOnOverlayClick := true
	if params.CloseOnOverlayClick != nil {
		closeOnOverlayClick = *params.CloseOnOverlayClick
	}

	id := params.ID
	if id == "" {
		id = fmt.Sprintf("dialog-%d", rand.Int())
	}

	return gt.Fragment(
		gt.IfLazy(
			params.Trigger != nil,
			func() gt.HTML {
				return gt.Button(
					gt.X.Type("button"),
					gt.X.Attr("onclick", fmt.Sprintf("document.getElementById('%s').showModal()", id)),
					params.TriggerAttrs,
					normalizeAnyComponent(params.Trigger),
				)
			},
		),
		gt.Dialog(
			gt.X.Id(id),
			gt.X.Class("dialog "+params.DialogExtraClasses),
			gt.X.Attr("aria-labelledby", id+"-title"),
			gt.If(params.Description != "", gt.X.Attr("aria-describedby", id+"-description")),
			gt.If(closeOnOverlayClick, gt.X.Attr("onclick", "if (event.target === this) this.close()")),
			params.DialogAttrs,
			gt.Div(
				gt.If(
					params.Title != "" || params.Description != "",
					gt.Header(
						params.HeaderAttrs,
						gt.H2(
							gt.X.Id(id+"-title"),
							params.Title,
						),
						gt.If(
							params.Description != "",
							gt.P(gt.X.Id(id+"-description"), params.Description),
						),
					),
				),
				gt.If(
					params.Body != nil,
					params.Body,
				),
				gt.If(
					params.Footer != nil,
					gt.Fragment(params.Footer),
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
		),
	)
}
