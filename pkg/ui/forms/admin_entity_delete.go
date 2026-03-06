package forms

import (
	"fmt"
	"net/http"

	"github.com/mikestefanello/pagoda/ent/admin"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/ui"
	. "github.com/mikestefanello/pagoda/pkg/ui/components"
	. "github.com/namzug16/gotags"
)

func AdminEntityDelete(r *ui.Request, entityType admin.EntityType) HTML {
	return Form(
		X.Class("form grid gap-6"),
		X.Method(http.MethodPost),
		P(
			fmt.Sprintf("Are you sure you want to delete this %s?", entityType.GetName()),
		),
		Div(
			X.Class("flex flex-col items-center gap-2"),
			Button(X.Class("btn btn-destructive w-full"), "Delete"),
			A(
				X.Href(r.Path(routenames.AdminEntityList(entityType.GetName()))),
				X.Class("btn-outline w-full"),
				"Cancel",
			),
		),
		CSRF(r),
	)
}
