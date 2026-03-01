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
			Div(
				X.Class("drawer lg:drawer-open"),
				Input(
					X.Id("sidebar"),
					X.Type("checkbox"),
					X.Class("drawer-toggle"),
				),
				Div(
					X.Class("drawer-content flex flex-col p-7 prose-base"),
					If(len(r.Title) > 0, H1(r.Title)),
					FlashMessages(r),
					content,
					Label(
						X.For("sidebar"),
						X.Class("btn btn-primary drawer-button lg:hidden"),
						"Open drawer",
					),
				),
				sidebarMenu(r),
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
	header := func(text string) HTML {
		return Li(
			X.Class("menu-title mt-3 uppercase"),
			Span(text),
		)
	}

	adminSubMenu := func() HTML {
		entityTypeLinks := make([]HTML, len(admin.GetEntityTypes()))
		for i, n := range admin.GetEntityTypes() {
			entityTypeLinks[i] = MenuLink(r, icons.PencilSquare(), n.GetName(), routenames.AdminEntityList(n.GetName()))
		}

		return Fragment(
			header("Entities"),
			Fragment(entityTypeLinks...),
			header("Monitoring"),
			Li(
				A(
					icons.CircleStack(),
					X.Href(r.Path(routenames.AdminTasks)),
					"Tasks",
					X.Target("_blank"),
				),
			),
		)
	}

	return Div(
		X.Class("drawer-side"),
		Label(
			X.For("sidebar"),
			X.Attr("aria-label", "close sidebar"),
			X.Class("drawer-overlay"),
		),
		Div(
			X.Class("menu bg-base-200 text-base-content min-h-full w-80 p-4"),
			Div(
				X.Class("w-2/3 mx-auto mt-3 mb-10"),
				Img(
					X.Src(ui.StaticFile("logo.png")),
				),
			),
			search(),
			Ul(
				HxBoost(),
				header("General"),
				MenuLink(r, icons.Home(), "Dashboard", routenames.Home),
				MenuLink(r, icons.Info(), "About", routenames.About),
				MenuLink(r, icons.Mail(), "Contact", routenames.Contact),
				MenuLink(r, icons.Archive(), "Cache", routenames.Cache),
				MenuLink(r, icons.CircleStack(), "Task", routenames.Task),
				MenuLink(r, icons.Document(), "Files", routenames.Files),
				header("Account"),
				If(r.IsAuth, MenuLink(r, icons.Exit(), "Logout", routenames.Logout)),
				If(!r.IsAuth, MenuLink(r, icons.Enter(), "Login", routenames.Login)),
				If(!r.IsAuth, MenuLink(r, icons.UserPlus(), "Register", routenames.Register)),
				If(!r.IsAuth, MenuLink(r, icons.QuestionCircle(), "Forgot password", routenames.ForgotPasswordSubmit)),
				If(r.IsAdmin, adminSubMenu()),
			),
		),
	)
}
