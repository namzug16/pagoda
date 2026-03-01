package pages

import (
	"fmt"
	"net/url"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/ent/admin"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/ui"
	. "github.com/mikestefanello/pagoda/pkg/ui/components"
	"github.com/mikestefanello/pagoda/pkg/ui/forms"
	"github.com/mikestefanello/pagoda/pkg/ui/layouts"
	. "github.com/namzug16/gotags"
)

func AdminEntityDelete(ctx echo.Context, entityType admin.EntityType) error {
	r := ui.NewRequest(ctx)
	r.Title = fmt.Sprintf("Delete %s", entityType.GetName())

	return r.Render(
		layouts.Primary,
		forms.AdminEntityDelete(r, entityType),
	)
}

func AdminEntityInput(ctx echo.Context, entityType admin.EntityType, values url.Values) error {
	r := ui.NewRequest(ctx)
	if values == nil {
		r.Title = fmt.Sprintf("Add %s", entityType.GetName())
	} else {
		r.Title = fmt.Sprintf("Edit %s", entityType.GetName())
	}

	return r.Render(
		layouts.Primary,
		forms.AdminEntity(r, entityType, values),
	)
}

func AdminEntityList(
	ctx echo.Context,
	entityType admin.EntityType,
	entityList *admin.EntityList,
) error {
	r := ui.NewRequest(ctx)
	r.Title = entityType.GetName()

	genHeader := func() HTML {
		g := make([]HTML, 0, len(entityList.Columns)+2)
		g = append(g, Th("ID"))
		for _, h := range entityList.Columns {
			g = append(g, Th(h))
		}
		g = append(g, Th())
		return Fragment(g...)
	}

	genRow := func(row admin.EntityValues) HTML {
		g := make([]HTML, 0, len(row.Values)+3)
		g = append(g, Th(fmt.Sprint(row.ID)))
		for _, h := range row.Values {
			g = append(g, Td(h))
		}
		g = append(g,
			Td(
				ButtonLink(
					ColorInfo,
					r.Path(routenames.AdminEntityEdit(entityType.GetName()), row.ID),
					"Edit",
				),
				Span(X.Class("mr-2")),
				ButtonLink(
					ColorError,
					r.Path(routenames.AdminEntityDelete(entityType.GetName()), row.ID),
					"Delete",
				),
			),
		)
		return Fragment(g...)
	}

	genRows := func() HTML {
		g := make([]HTML, 0, len(entityList.Entities))
		for _, row := range entityList.Entities {
			g = append(g, Tr(genRow(row)))
		}
		return Fragment(g...)
	}

	return r.Render(layouts.Primary, Fragment(
		Div(
			X.Class("form-control mb-2"),
			ButtonLink(
				ColorAccent,
				r.Path(routenames.AdminEntityAdd(entityType.GetName())),
				fmt.Sprintf("Add %s", entityType.GetName()),
			),
		),
		Table(
			X.Class("table table-zebra mb-2"),
			Thead(
				Tr(genHeader()),
			),
			Tbody(genRows()),
		),
		Pager(
			entityList.Page,
			r.Path(routenames.AdminEntityAdd(entityType.GetName())),
			entityList.HasNextPage,
			"",
		),
	))
}
