package layouts

import (
	"github.com/mikestefanello/pagoda/ent/admin"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/ui"
	"github.com/mikestefanello/pagoda/pkg/ui/cache"
	. "github.com/mikestefanello/pagoda/pkg/ui/components"
	"github.com/mikestefanello/pagoda/pkg/ui/icons"
	"github.com/mikestefanello/pagoda/pkg/ui/lucide"
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
					lucide.Database(),
					Span(n.GetName()),
				),
			)
		}

		return Fragment(
			Li(
				A(
					X.Href(r.Path(routenames.AdminTasks)),
					lucide.ListTodo(),
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
				icons.Home(),
				Span("Dashboard"),
			),
		),
		Li(
			A(
				X.Href(r.Path(routenames.KitchenSink)),
				lucide.Blocks(),
				Span("Kitchen Sink"),
			),
		),
		Li(
			A(
				X.Href(r.Path(routenames.About)),
				icons.Info(),
				Span("About"),
			),
		),
		Li(
			A(
				X.Href(r.Path(routenames.Contact)),
				icons.Mail(),
				Span("Contact"),
			),
		),
		Li(
			A(
				X.Href(r.Path(routenames.Cache)),
				icons.CircleStack(),
				Span("Cache"),
			),
		),
		Li(
			A(
				X.Href(r.Path(routenames.Task)),
				lucide.ListTodo(),
				Span("Task"),
			),
		),
		Li(
			A(
				X.Href(r.Path(routenames.Files)),
				lucide.Folder(),
				Span("Files"),
			),
		),
	}

	accountItems := []HTML{}
	if r.IsAuth {
		accountItems = append(accountItems, Li(
			A(
				X.Href(r.Path(routenames.Logout)),
				icons.Exit(),
				Span("Logout"),
			),
		))
	} else {
		accountItems = append(accountItems, Li(
			A(
				X.Href(r.Path(routenames.Login)),
				icons.Enter(),
				Span("Login"),
			),
		))
		accountItems = append(accountItems, Li(
			A(
				X.Href(r.Path(routenames.Register)),
				icons.UserPlus(),
				Span("Register"),
			),
		))
		accountItems = append(accountItems, Li(
			A(
				X.Href(r.Path(routenames.ForgotPasswordSubmit)),
				icons.LockClosed(),
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
										lucide.Shield(),
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
