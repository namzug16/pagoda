package layouts

import (
	"github.com/mikestefanello/pagoda/ent/admin"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/ui"
	"github.com/mikestefanello/pagoda/pkg/ui/cache"
	. "github.com/mikestefanello/pagoda/pkg/ui/components"
	"github.com/mikestefanello/pagoda/pkg/ui/icons"
	. "github.com/namzug16/gotags"
)

func Primary(r *ui.Request, content HTML) HTML {
	return Doc(
		Head(
			Metatags(r),
			CSS(),
			JS(),
		),
		Body(
			X.Attr("data-theme", "dark"),
			sidebarMenu(r),
			Main(
				X.Id("content"),
				X.Class("flex flex-col p-7 prose-base"),
				If(len(r.Title) > 0, H1(r.Title)),
				FlashMessages(r),
				content,
			),
			searchModal(r),
			HtmxListeners(r),
		),
	)
}

func search() HTML {
	return cache.SetIfNotExists("layout.search", func() HTML {
		return Div(
			X.Class("ml-2"),
			X.Attr("x-data", ""),
			Label(
				X.Class("input"),
				icons.MagnifyingGlass(),
				Input(
					X.Type("search"),
					X.Class("grow"),
					X.Placeholder("Search"),
					X.Attr("@click", "search_modal.showModal();"),
				),
			),
		)
	})
}

func searchModal(r *ui.Request) HTML {
	return cache.SetIfNotExists("layout.searchModal", func() HTML {
		return Dialog(
			X.Id("search_modal"),
			X.Class("modal"),
			Div(
				X.Class("modal-box"),
				Form(
					X.Method("dialog"),
					Button(
						X.Class("btn btn-sm btn-circle btn-ghost absolute right-2 top-2"),
						"✕",
					),
				),
				H3(
					X.Class("text-lg font-bold mb-2"),
					"Search",
				),
				Input(
					X.Attr("hx-get", r.Path(routenames.Search)),
					X.Attr("hx-trigger", "keyup changed delay:500ms"),
					X.Attr("hx-target", "#results"),
					X.Name("query"),
					X.Class("input w-full"),
					X.Type("search"),
					X.Placeholder("Search..."),
				),
				Ul(
					X.Id("results"),
					X.Class("list"),
				),
			),
			Form(
				X.Method("dialog"),
				X.Class("modal-backdrop"),
				Button(
					"close",
				),
			),
		)
	})
}

func sidebarMenu(r *ui.Request) HTML {
	adminSubMenu := func() HTML {
		entityTypeLinks := make([]HTML, len(admin.GetEntityTypes()))
		for i, n := range admin.GetEntityTypes() {
			entityTypeLinks[i] = Li(
				A(
					X.Href(r.Path(routenames.AdminEntityList(n.GetName()))),
					Span(n.GetName()),
				),
			)
		}

		return Fragment(
			Li(
				A(
					X.Href(r.Path(routenames.AdminTasks)),
					Span("Tasks"),
					X.Target("_blank"),
				),
			),
			Fragment(entityTypeLinks...),
		)
	}

	group := func(id, title string, items []HTML) HTML {
		return Div(
			X.Role("group"),
			X.Attr("aria-labelledby", id),
			H3(
				X.Id(id),
				title,
			),
			Ul(
				HxBoost(),
				Fragment(items...),
			),
		)
	}

	generalItems := []HTML{
		Li(
			A(
				X.Href(r.Path(routenames.Home)),
				Span("Dashboard"),
			),
		),
		Li(
			A(
				X.Href(r.Path(routenames.About)),
				Span("About"),
			),
		),
		Li(
			A(
				X.Href(r.Path(routenames.Contact)),
				Span("Contact"),
			),
		),
		Li(
			A(
				X.Href(r.Path(routenames.Cache)),
				Span("Cache"),
			),
		),
		Li(
			A(
				X.Href(r.Path(routenames.Task)),
				Span("Task"),
			),
		),
		Li(
			A(
				X.Href(r.Path(routenames.Files)),
				Span("Files"),
			),
		),
	}

	accountItems := []HTML{}
	if r.IsAuth {
		accountItems = append(accountItems, Li(
			A(
				X.Href(r.Path(routenames.Logout)),
				Span("Logout"),
			),
		))
	} else {
		accountItems = append(accountItems, Li(
			A(
				X.Href(r.Path(routenames.Login)),
				Span("Login"),
			),
		))
		accountItems = append(accountItems, Li(
			A(
				X.Href(r.Path(routenames.Register)),
				Span("Register"),
			),
		))
		accountItems = append(accountItems, Li(
			A(
				X.Href(r.Path(routenames.ForgotPasswordSubmit)),
				Span("Forgot password"),
			),
		))
	}

	return Aside(
		X.Class("sidebar"),
		X.Attr("data-side", "left"),
		X.Attr("aria-hidden", "false"),
		Nav(
			X.Attr("aria-label", "Sidebar navigation"),
			Header(
				A(
					X.Href("/"),
					X.Class("btn-ghost p-2 h-12 w-full justify-start"),
					Div(
						X.Class("bg-sidebar-primary text-sidebar-primary-foreground flex aspect-square size-8 items-center justify-center rounded-lg"),
						Img(
							X.Src(ui.StaticFile("logo.png")),
							X.Alt("Logo"),
						),
					),
					Div(
						X.Class("grid flex-1 text-left text-sm leading-tight"),
						Span(X.Class("truncate font-medium"), Text("Pagoda")),
						Span(X.Class("truncate text-xs"), Text("v0.3.11")),
					),
				),
			),
			Section(
				X.Class("scrollbar"),
				group("group-label-general", "Getting started", generalItems),
				If(r.IsAdmin,
					group("group-label-admin", "Admin",
						[]HTML{
							Li(
								Details(
									X.Id("submenu-admin"),
									Summary(
										X.Attr("aria-controls", "submenu-admin-content"),
										"Admin",
									),
									Ul(
										X.Id("submenu-admin-content"),
										adminSubMenu(),
									),
								),
							),
						},
					),
				),
				group("group-label-account", "Account", accountItems),
			),
		),
	)
}

// <header>
// <a href="/" class="btn-ghost p-2 h-12 w-full justify-start">
//   <div class="bg-sidebar-primary text-sidebar-primary-foreground flex aspect-square size-8 items-center justify-center rounded-lg">
//     <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 256 256" class="h-4 w-4"><rect width="256" height="256" fill="none"></rect><line x1="208" y1="128" x2="128" y2="208" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32"></line><line x1="192" y1="40" x2="40" y2="192" fill="none" stroke="currentColor" stroke-linecap="round" stroke-linejoin="round" stroke-width="32"></line></svg>
//   </div>
//   <div class="grid flex-1 text-left text-sm leading-tight">
//     <span class="truncate font-medium">Basecoat</span>
//     <span class="truncate text-xs">v0.3.11</span>
//   </div>
// </a>
// </header>
