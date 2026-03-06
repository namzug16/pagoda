package pages

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/ui"
	"github.com/mikestefanello/pagoda/pkg/ui/layouts"
	"github.com/mikestefanello/pagoda/pkg/ui/lucide"
	kitchensink "github.com/mikestefanello/pagoda/pkg/ui/pages/kitchen_sink"
	. "github.com/namzug16/gotags"
)

func KitchenSink(ctx echo.Context) error {
	r := ui.NewRequest(ctx)

	return r.Render(layouts.Primary, html())
}

func html() HTML {
	return Div(
		X.Class("w-full flex gap-x-10"),
		Div(
			X.Class("w-full"),
			Header(
				X.Class("space-y-2 mb-8"),
				H1(X.Class("text-3xl font-semibold tracking-tight"), "Basecoat Kitchen Sink"),
				P(X.Class("text-muted-foreground"), "A collection of all the components available in Basecoat"),
			),
			Div(
				X.Class("grid gap-4 flex-1"),
				Range(sections, func(i int, s section) HTML {
					return _section(s)
				}),
			),
		),
		Div(
			X.Class("text-sm w-full max-w-[200px] shrink-0"),
			Nav(
				X.Class("sticky top-22 max-h-[calc(100vh-6rem)] overflow-y-auto scrollbar pr-2 space-y-2 [&_ul]:m-0 [&_ul]:list-none [&_ul_ul]:pl-4 [&_li]:mt-0 [&_li]:pt-2 [&_a]:inline-block [&_a]:no-underline [&_a]:transition-colors [&_a]:hover:text-foreground [&_a]:text-muted-foreground"),
				H4(X.Class("font-medium"), "On This Page"),
				Ul(
					Range(sections, func(i int, s section) HTML {
						return Li(A(X.Href("#"+sectionID(s.name)), s.name))
					}),
				),
			),
		),
	)
}

type section struct {
	name    string
	href    string
	content HTML
}

var sections = []section{
	section{name: "Accordion", href: "/components/accordion", content: kitchensink.Accordion()},
	section{name: "Alert", href: "/components/alert", content: kitchensink.AlertSection()},
	section{name: "Alert Dialog", href: "/components/alert-dialog", content: kitchensink.AlertDialogSection()},
	section{name: "Avatar", href: "/components/avatar", content: kitchensink.AvatarSection()},
	section{name: "Badge", href: "/components/badge", content: kitchensink.BadgeSection()},
	section{name: "Breadcrumb", href: "/components/breadcrumb", content: kitchensink.BreadcrumbSection()},
	section{name: "Button", href: "/components/button", content: kitchensink.ButtonSection()},
	section{name: "Card", href: "/components/card", content: kitchensink.CardSection()},
	section{name: "Checkbox", href: "/components/checkbox", content: kitchensink.CheckboxSection()},
	section{name: "Combobox", href: "/components/combobox", content: kitchensink.ComboboxSection()},
	section{name: "Dialog", href: "/components/dialog", content: kitchensink.DialogSection()},
	section{name: "Dropdown Menu", href: "/components/dropdown-menu", content: kitchensink.DropdownMenuSection()},
	section{name: "Form", href: "/components/form", content: kitchensink.FormSection()},
	section{name: "Input", href: "/components/input", content: kitchensink.InputSection()},
	section{name: "Label", href: "/components/label", content: kitchensink.LabelSection()},
	section{name: "Pagination", href: "/components/pagination", content: kitchensink.PaginationSection()},
	section{name: "Popover", href: "/components/popover", content: kitchensink.PopoverSection()},
	section{name: "Radio Group", href: "/components/radio-group", content: kitchensink.RadioGroupSection()},
	section{name: "Select", href: "/components/select", content: kitchensink.SelectSection()},
	section{name: "Skeleton", href: "/components/skeleton", content: kitchensink.SkeletonSection()},
	section{name: "Slider", href: "/components/slider", content: kitchensink.SliderSection()},
	section{name: "Switch", href: "/components/switch", content: kitchensink.SwitchSection()},
	section{name: "Table", href: "/components/table", content: kitchensink.TableSection()},
	section{name: "Tabs", href: "/components/tabs", content: kitchensink.TabsSection()},
	section{name: "Textarea", href: "/components/textarea", content: kitchensink.TextareaSection()},
	section{name: "Toast", href: "/components/toast", content: kitchensink.ToastSection()},
	section{name: "Tooltip", href: "/components/tooltip", content: kitchensink.TooltipSection()},
}

func _section(s section) HTML {
	return Section(
		X.Id(sectionID(s.name)),
		X.Class("w-full rounded-lg border scroll-mt-14"),
		Header(
			X.Class("border-b px-4 py-3 flex items-center justify-between"),
			H2(X.Class("text-sm font-medium"), s.name),
			A(
				X.Href("https://basecoatui.com"+s.href),
				X.Class("text-muted-foreground hover:text-foreground"),
				X.Attr("data-tooltip", "See documentation"),
				X.Attr("data-side", "left"),
				lucide.BookOpen(X.Class("size-4")),
			),
		),
		s.content,
	)
}

func sectionID(name string) string {
	return strings.ReplaceAll(strings.ToLower(name), " ", "-")
}
