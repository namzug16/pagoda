package layouts

import (
	"github.com/mikestefanello/pagoda/pkg/ui"
	. "github.com/mikestefanello/pagoda/pkg/ui/components"
	. "github.com/namzug16/gotags"
)

func Auth(r *ui.Request, content HTML) HTML {
	return Doc(
		Head(
			Metatags(r),
			CSS(),
			JS(),
		),
		Body(
			X.Attr("data-theme", "dark"),
			Div(
				X.Class("hero flex items-center justify-center min-h-screen"),
				Div(
					X.Class("flex-col hero-content"),
					Div(
						X.Class("card shadow-md bg-base-200 w-96"),
						Div(
							X.Class("card-body"),
							If(len(r.Title) > 0, H1(X.Class("text-2xl font-bold"), r.Title)),
							FlashMessages(r),
							content,
						),
					),
				),
			),
			HtmxListeners(r),
		),
	)
}
