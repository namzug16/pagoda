package components

import (
	"fmt"

	"github.com/mikestefanello/pagoda/pkg/pager"
	"github.com/mikestefanello/pagoda/pkg/ui"
	"github.com/mikestefanello/pagoda/pkg/ui/lucide"
	. "github.com/namzug16/gotags"
)

func MenuLink(r *ui.Request, icon HTML, title, routeName string, routeParams ...any) HTML {
	href := r.Path(routeName, routeParams...)

	return Li(
		X.Class("ml-2"),
		A(
			X.Href(href),
			icon,
			title,
			X.Class(
				"p-2",
				X.If(href == r.CurrentPath, "menu-active"),
			),
		),
	)
}

func Pager(page int, path string, hasNext bool, hxTarget string) HTML {
	href := func(page int) string {
		return fmt.Sprintf("%s?%s=%d",
			path,
			pager.QueryKey,
			page,
		)
	}

	return Div(
		X.Class("w-full flex flex-row justify-center gap-1"),
		A(
			X.Class("btn-ghost"),
			If(page <= 1, X.Disabled()),
			X.Href(href(page-1)),
			If(
				len(hxTarget) > 0,
				X.Attr("hx-get", href(page-1)),
				X.Attr("hx-swap", "outerHTML"),
				X.Attr("hx-target", hxTarget),
			),
			lucide.ChevronLeft(),
		),
		Button(
			X.Class("btn-outline"),
			fmt.Sprintf("Page %d", page),
		),
		A(
			X.Href(href(page+1)),
			X.Class("btn-ghost"),
			If(!hasNext, X.Disabled()),
			If(len(hxTarget) > 0,
				X.Attr("hx-get", href(page+1)),
				X.Attr("hx-swap", "outerHTML"),
				X.Attr("hx-target", hxTarget),
			),
			lucide.ChevronRight(),
		),
	)
}
