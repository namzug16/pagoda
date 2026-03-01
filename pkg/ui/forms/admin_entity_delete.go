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
		X.Method(http.MethodPost),
		P(
			fmt.Sprintf("Are you sure you want to delete this %s?", entityType.GetName()),
		),
		ControlGroup(
			FormButton(ColorError, "Delete"),
			ButtonLink(
				ColorNone,
				r.Path(routenames.AdminEntityList(entityType.GetName())),
				"Cancel",
			),
		),
		CSRF(r),
	)
}
