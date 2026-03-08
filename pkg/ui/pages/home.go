package pages

import (
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/ui"
	"github.com/mikestefanello/pagoda/pkg/ui/layouts"
	"github.com/mikestefanello/pagoda/pkg/ui/lucide"
	"github.com/mikestefanello/pagoda/pkg/ui/models"
	. "github.com/namzug16/gotags"
)

func Home(ctx echo.Context, posts *models.Posts) error {
	r := ui.NewRequest(ctx)
	r.Metatags.Description = "This is the home page."
	r.Metatags.Keywords = []string{"Software", "Coding", "Go"}

	// This pages helps to illustrate the different options you can take when using HTMX to introduce interactivity
	// to your web application. The following three options are available, but here, we're opting for the first one.
	// 1) Highly-optimized and progressive enhancement:
	//    This is highly-optimized because the server is doing the least amount of work possible, only rendering
	//    the least amount possible based on the incoming request. It's possible that even your route handler would
	//    want to check the HTMX request in order to limit what it does. With HTMX, it's possible to still return a
	//    normal, full page, but use hx-select to pluck out only the part you want to re-render. It requires some extra
	//    condition checks and code but performance is improved. Progressive enhancement refers to having a fully
	//    functional web app, even if JS was disabled, but providing the enhancement if JS is enabled. All of these
	//    examples should continue to work fine without JS.
	// 2) Not optimized and progressive enhancement:
	//    As mentioned previously, you can remove all of these conditions, re-render the entire page for every request,
	//    and rely on HTMX's hx-select to only replace what you want to (ie, the posts).
	// 3) Optimized and partial renderings:
	//    You could have a separate route that is only for fetching posts while paging, and that would render only
	//    that partial HTML, which HTMX would then use to inject in to this page.

	html := Div(
		X.Class("flex flex-col gap-4 w-full"),
		If(
			r.Htmx.Target != "posts",
			Div(
				X.Class("grid grid-cols-1 md:grid-cols-3 gap-4"),
				Div(
					X.Class("card w-full"),
					Header(
						P("User name"),
						Div(
							X.Class("flex flex-row items-center justify-between"),
							H2(
								X.Class("text-2xl font-semibold tabular-nums @[250px]/card:text-3xl"),
								func() string {
									if r.IsAuth {
										return r.AuthUser.Name
									}
									return "(not logged in)"
								}(),
							),
							lucide.CircleUser(),
						),
					),
					Footer(
						X.Class("text-sm"),
						P(
							X.Class("text-muted-foreground"),
							"Trending up this month",
						),
					),
				),
				Div(
					X.Class("card w-full"),
					Header(
						P("Admin status"),
						Div(
							X.Class("flex flex-row gap-0.5 items-center justify-between"),
							H2(
								X.Class("text-2xl font-semibold tabular-nums @[250px]/card:text-3xl"),
								func() string {
									if r.IsAdmin {
										return "Administrator"
									}
									return "Non-administrator"
								}(),
							),
							lucide.Lock(),
						),
					),
					Footer(
						X.Class("text-sm"),
						P(
							X.Class("text-muted-foreground"),
							"Use `make admin` to create an admin account",
						),
					),
				),
				Div(
					X.Class("card w-full"),
					Header(
						P("Github Stars"),
						Div(
							X.Class("flex flex-row gap-0.5 items-center justify-between"),
							H2(
								X.Class("text-2xl font-semibold tabular-nums @[250px]/card:text-3xl"),
								"2,500+",
							),
							lucide.Star(),
						),
					),
					Footer(
						X.Class("text-sm"),
						P(
							X.Class("text-muted-foreground"),
							"Star if you like Pagoda",
						),
					),
				),
			),
			Div(
				H2(X.Class("text-xl mb-2"), "Recent posts"),
				P("Below is an example of both paging and AJAX fetching using HTMX"),
			),
		),
		posts.Render(r.Path(routenames.Home)),
		If(
			r.Htmx.Target != "posts",
			Div(
				X.Class("grid grid-cols-1 md:grid-cols-3 gap-4"),
				Div(
					X.Class("col-span-2 card"),
					Header(
						H2("Serving files"),
						P("In the example posts above, check how the file URL contains a cache-buster query parameter which changes only when the app is restarted. "),
						P("Static files also contain cache-control headers which are configured via middleware."),
					),
				),
				Div(
					X.Class("card"),
					Header(
						H2("Documentation"),
						P("Have you read through the entire documentation? If not, you may be missing functionality or have questions. "),
					),
					Footer(
						X.Class("flex justify-end"),
						A(
							X.Href("https://github.com/mikestefanello/pagoda?tab=readme-ov-file#table-of-contents"),
							X.Class("btn-link"),
							"Learn more",
						),
					),
				),
			),
		),
	)

	return r.Render(layouts.Primary, html)
}
