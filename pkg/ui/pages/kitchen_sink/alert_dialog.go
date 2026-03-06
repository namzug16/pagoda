package kitchensink

import (
	"github.com/mikestefanello/pagoda/pkg/ui/basecoat"
	. "github.com/namzug16/gotags"
)

func AlertDialogSection() HTML {
	return Div(
		X.Class("p-4"),
		basecoat.Dialog(basecoat.DialogParams{
			ID:          "alert-dialog-demo",
			Title:       "Are you absolutely sure?",
			Description: "This action cannot be undone. This will permanently delete your account and remove your data from our servers.",
			Trigger:     "Open dialog",
			TriggerAttrs: []HTML{
				X.Class("btn-outline"),
			},
			Footer: Footer(
				Button(
					X.Type("button"),
					X.Class("btn-outline"),
					X.Attr("onclick", "document.getElementById('alert-dialog-demo').close()"),
					"Cancel",
				),
				Button(
					X.Type("button"),
					X.Class("btn-primary"),
					X.Attr("onclick", "document.getElementById('alert-dialog-demo').close()"),
					"Continue",
				),
			),
		}),
	)
}
