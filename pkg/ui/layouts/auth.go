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
				X.Class("flex items-center justify-center min-h-screen"),
				FlashMessages(r),
				content,
			),
			HtmxListeners(r),
		),
	)
}
