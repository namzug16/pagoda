package pages

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/ui"
	"github.com/mikestefanello/pagoda/pkg/ui/basecoat"
	"github.com/mikestefanello/pagoda/pkg/ui/layouts"
	"github.com/mikestefanello/pagoda/pkg/ui/lucide"
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
		//FIX: we missing: on this page...
	)
}

type section struct {
	name    string
	href    string
	content HTML
}

var sections = []section{
	section{name: "Accordion", href: "/components/accordion", content: _accordionSection()},
	section{name: "Alert", href: "/components/alert", content: _alertSection()},
	section{name: "Alert Dialog", href: "/components/alert-dialog", content: _alertDialogSection()},
	section{name: "Avatar", href: "/components/avatar", content: _avatarSection()},
	section{name: "Badge", href: "/components/badge", content: _badgeSection()},
	section{name: "Breadcrumb", href: "/components/breadcrumb", content: _breadcrumbSection()},
	section{name: "Button", href: "/components/button", content: _buttonSection()},
	section{name: "Card", href: "/components/card", content: _cardSection()},
	section{name: "Checkbox", href: "/components/checkbox", content: _checkboxSection()},
	section{name: "Combobox", href: "/components/combobox", content: _comboboxSection()},
	section{name: "Dialog", href: "/components/dialog", content: _dialogSection()},
	section{name: "Dropdown Menu", href: "/components/dropdown-menu", content: _dropdownMenuSection()},
	section{name: "Form", href: "/components/form", content: _formSection()},
	section{name: "Input", href: "/components/input", content: _inputSection()},
	section{name: "Label", href: "/components/label", content: _labelSection()},
	section{name: "Pagination", href: "/components/pagination", content: _paginationSection()},
	section{name: "Popover", href: "/components/popover", content: _popoverSection()},
	section{name: "Radio Group", href: "/components/radio-group", content: _radioGroupSection()},
	section{name: "Select", href: "/components/select", content: _selectSection()},
	section{name: "Skeleton", href: "/components/skeleton", content: _skeletonSection()},
	section{name: "Slider", href: "/components/slider", content: _sliderSection()},
	section{name: "Switch", href: "/components/switch", content: _switchSection()},
	section{name: "Table", href: "/components/table", content: _tableSection()},
	section{name: "Tabs", href: "/components/tabs", content: _tabsSection()},
	section{name: "Textarea", href: "/components/textarea", content: _textareaSection()},
	section{name: "Toast", href: "/components/toast", content: _toastSection()},
	section{name: "Tooltip", href: "/components/tooltip", content: _tooltipSection()},
}

func _section(s section) HTML {
	return Section(
		X.Id(strings.ToLower(s.name)),
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

func _accordionSection() HTML {
	return Div(
		X.Class("p-4"),
		Section(
			X.Class("accordion"),
			Details(
				X.Class("group border-b last:border-b-0"),
				Summary(
					X.Class("w-full focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] transition-all outline-none rounded-md"),
					H2(
						X.Class("flex flex-1 items-start justify-between gap-4 py-4 text-left text-sm font-medium hover:underline "),
						"Is it accessible?",
						lucide.ChevronDown(X.Class("text-muted-foreground pointer-events-none size-4 shrink-0 translate-y-0.5 transition-transform duration-200 group-open:rotate-180")),
					),
				),
				Section(
					X.Class("pb-4"),
					P(
						X.Class("text-sm"),
						"Yes. It adheres to the WAI-ARIA design pattern.",
					),
				),
			),
			Details(
				X.Class("group border-b last:border-b-0"),
				Summary(
					X.Class("w-full focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] transition-all outline-none rounded-md"),
					H2(
						X.Class("flex flex-1 items-start justify-between gap-4 py-4 text-left text-sm font-medium hover:underline "),
						"Is it styled?",
						lucide.ChevronDown(X.Class("text-muted-foreground pointer-events-none size-4 shrink-0 translate-y-0.5 transition-transform duration-200 group-open:rotate-180")),
					),
				),
				Section(
					X.Class("pb-4"),
					P(
						X.Class("text-sm"),
						"Yes. It comes with default styles that matches the other components' aesthetic.",
					),
				),
			),
			Details(
				X.Class("group border-b last:border-b-0"),
				Summary(
					X.Class("w-full focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] transition-all outline-none rounded-md"),
					H2(
						X.Class("flex flex-1 items-start justify-between gap-4 py-4 text-left text-sm font-medium hover:underline "),
						"Is it animated?",
						lucide.ChevronDown(X.Class("text-muted-foreground pointer-events-none size-4 shrink-0 translate-y-0.5 transition-transform duration-200 group-open:rotate-180")),
					),
				),
				Section(
					X.Class("pb-4"),
					P(
						X.Class("text-sm"),
						"Yes. It's animated by default, but you can disable it if you prefer.",
					),
				),
			),
		),
		Section(
			X.Class("accordion"),
			Details(
				X.Class("group border-b last:border-b-0"),
				Summary(
					X.Class("w-full focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] transition-all outline-none rounded-md"),
					H2(
						X.Class("flex flex-1 items-start justify-between gap-4 py-4 text-left text-sm font-medium hover:underline "),
						"What are the key considerations when implementing a comprehensive enterprise-level authentication system?",
						lucide.ChevronDown(X.Class("text-muted-foreground pointer-events-none size-4 shrink-0 translate-y-0.5 transition-transform duration-200 group-open:rotate-180")),
					),
				),
				Section(
					X.Class("pb-4"),
					P(
						X.Class("text-sm"),
						"Implementing a robust enterprise authentication system requires careful consideration of multiple factors. This includes secure password hashing and storage, multi-factor authentication (MFA) implementation, session management, OAuth2 and SSO integration, regular security audits, rate limiting to prevent brute force attacks, and maintaining detailed audit logs. Additionally, you'll need to consider scalability, performance impact, and compliance with relevant data protection regulations such as GDPR or HIPAA.",
					),
				),
			),
			Details(
				X.Class("group border-b last:border-b-0"),
				Summary(
					X.Class("w-full focus-visible:border-ring focus-visible:ring-ring/50 focus-visible:ring-[3px] transition-all outline-none rounded-md"),
					H2(
						X.Class("flex flex-1 items-start justify-between gap-4 py-4 text-left text-sm font-medium hover:underline "),
						"How does modern distributed system architecture handle eventual consistency and data synchronization across multiple regions?",
						lucide.ChevronDown(X.Class("text-muted-foreground pointer-events-none size-4 shrink-0 translate-y-0.5 transition-transform duration-200 group-open:rotate-180")),
					),
				),
				Section(
					X.Class("pb-4"),
					P(
						X.Class("text-sm"),
						"Modern distributed systems employ various strategies to maintain data consistency across regions. This often involves using techniques like CRDT (Conflict-Free Replicated Data Types), vector clocks, and gossip protocols. Systems might implement event sourcing patterns, utilize message queues for asynchronous updates, and employ sophisticated conflict resolution strategies. Popular solutions like Amazon's DynamoDB and Google's Spanner demonstrate different approaches to solving these challenges, balancing between consistency, availability, and partition tolerance as described in the CAP theorem.",
					),
				),
			),
		),
		Raw(`
      <script>
        (() => {
          const accordions = document.querySelectorAll('.accordion');
          accordions.forEach(accordion => {
            accordion.addEventListener('click', (event) => {
              const summary = event.target.closest('summary');
              if (!summary) return;
              const details = summary.closest('details');
              if (!details) return;
              accordion.querySelectorAll('details').forEach(detailsEl => {
                if (detailsEl !== details) {
                  detailsEl.removeAttribute('open');
                }
              });
            });
          });
        })();
      </script>
		`),
	)
}

func _alertSection() HTML {
	return Div(
		X.Class("grid max-w-xl items-start gap-4 p-4"),
		Div(
			X.Class("alert"),
			lucide.CircleCheck(X.Class("shrink-0")),
			H2("Success! Your changes have been saved"),
			Section(X.Class("text-muted-foreground"), "This is an alert with icon, title and description."),
		),
		Div(
			X.Class("alert"),
			lucide.BookmarkCheck(X.Class("shrink-0")),
			Section(X.Class("text-muted-foreground"), "This is an alert with icon, description and no title."),
		),
		Div(
			X.Class("alert"),
			Section(X.Class("text-muted-foreground"), "This one has a description only. No title. No icon."),
		),
		Div(
			X.Class("alert"),
			lucide.Popcorn(X.Class("shrink-0")),
			H2("Let's try one with icon and title."),
		),
		Div(
			X.Class("alert"),
			lucide.ShieldAlert(X.Class("shrink-0")),
			H2("This is a very long alert title that demonstrates how the component handles extended text content and potentially wraps across multiple lines"),
		),
		Div(
			X.Class("alert"),
			lucide.Gift(X.Class("shrink-0")),
			Section(X.Class("text-muted-foreground"), "This is a very long alert description that demonstrates how the component handles extended text content and potentially wraps across multiple lines"),
		),
		Div(
			X.Class("alert"),
			lucide.Info(X.Class("shrink-0")),
			H2("This is an extremely long alert title that spans multiple lines to demonstrate how the component handles very lengthy headings while maintaining readability and proper text wrapping behavior"),
			Section(X.Class("text-muted-foreground"), "This is an equally long description that contains detailed information about the alert. It shows how the component can accommodate extensive content while preserving proper spacing, alignment, and readability across different screen sizes and viewport widths. This helps ensure the user experience remains consistent regardless of the content length."),
		),
		Div(
			X.Class("alert-destructive"),
			lucide.Info(X.Class("shrink-0")),
			H2("Something went wrong!"),
			Section(X.Class("text-muted-foreground"), "Your session has expired. Please log in again."),
		),
		Div(
			X.Class("alert-destructive"),
			lucide.Info(X.Class("shrink-0")),
			H2("Something went wrong!"),
			Section(
				X.Class("text-muted-foreground"),
				P("Please verify your billing information and try again."),
				Ul(
					X.Class("list-disc pl-5 mt-2"),
					Li("Check your card details"),
					Li("Ensure sufficient funds"),
					Li("Verify billing address"),
				),
			),
		),
		Div(
			X.Class("alert border-amber-50 bg-amber-50 text-amber-900 dark:border-amber-950 dark:bg-amber-950 dark:text-amber-100"),
			lucide.CircleCheck(X.Class("shrink-0")),
			H2("Plot Twist: This Alert is Actually Amber!"),
			Section(X.Class("text-muted-foreground"), "This one has custom colors for light and dark mode."),
		),
	)
}

func _dialogSection() HTML {
	return Div(
		X.Class("p-4"),
		Div(
			X.Class("flex flex-wrap items-center gap-4"),
			Fragment(
				basecoat.Dialog(basecoat.DialogParams{
					ID:                 "demo-dialog-edit-profile",
					DialogExtraClasses: "w-full sm:max-w-[425px] max-h-[612px]",
					Title:              "Edit profile",
					Description:        "Make changes to your profile here. Click save when you're done.",
					Trigger:            "Edit Profile",
					TriggerAttrs: []HTML{
						X.Class("btn-outline"),
					},
					Body: Section(
						Form(
							X.Class("form grid gap-4"),
							Div(
								X.Class("grid gap-3"),
								Label(X.For("demo-dialog-edit-profile-name"), "Name"),
								Input(X.Type("text"), X.Value("Pedro Duarte"), X.Id("demo-dialog-edit-profile-name")),
							),
							Div(
								X.Class("grid gap-3"),
								Label(X.For("demo-dialog-edit-profile-username"), "Username"),
								Input(X.Type("text"), X.Value("@peduarte"), X.Id("demo-dialog-edit-profile-username")),
							),
						),
					),
					Footer: Footer(
						Button(
							X.Type("button"),
							X.Class("btn-outline"),
							X.Attr("onclick", "document.getElementById('demo-dialog-edit-profile').close()"),
							"Cancel",
						),
						Button(
							X.Type("button"),
							X.Class("btn"),
							X.Attr("onclick", "document.getElementById('demo-dialog-edit-profile').close()"),
							"Save changes",
						),
					),
				}),
			),
			basecoat.Dialog(basecoat.DialogParams{
				ID:                 "dialog-example",
				DialogExtraClasses: "w-full sm:max-w-[425px] max-h-[612px]",
				Title:              "Scrollable Content",
				Description:        "This is a dialog with scrollable content.",
				Trigger:            "Scrollable Content",
				TriggerAttrs: []HTML{
					X.Class("btn-outline"),
				},
				Body: Section(
					X.Class("overflow-y-auto scrollbar"),
					Div(
						X.Class("space-y-4 text-sm"),
						P("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."),
						P("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."),
						P("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."),
						P("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."),
						P("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."),
						P("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."),
						P("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."),
						P("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."),
						P("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."),
						P("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."),
						P("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."),
						P("Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum."),
					),
				),
				Footer: Footer(
					Button(
						X.Type("button"),
						X.Class("btn-outline"),
						X.Attr("onclick", "this.closest('dialog').close()"),
						"Close",
					),
				),
			}),
		),
	)
}

func _dropdownMenuSection() HTML {
	return Div(
		X.Class("p-4"),
		Div(
			X.Class("flex flex-wrap items-start gap-4"),
			basecoat.DropdownMenu(basecoat.DropdownMenuParams{
				ID:                  "dropdown-menu-default",
				Trigger:             "Open",
				TriggerExtraClasses: "btn-outline",
				PopoverExtraClasses: "min-w-56",
				Content: Fragment(
					Div(
						X.Attr("role", "group"),
						X.Attr("aria-labelledby", "demo-account-options"),
						Div(X.Attr("role", "heading"), X.Id("demo-account-options"), "My Account"),
						Div(X.Attr("role", "menuitem"), "Profile", Span(X.Class("text-muted-foreground ml-auto text-xs tracking-widest"), "Shift+Cmd+P")),
						Div(X.Attr("role", "menuitem"), "Billing", Span(X.Class("text-muted-foreground ml-auto text-xs tracking-widest"), "Cmd+B")),
						Div(X.Attr("role", "menuitem"), "Settings", Span(X.Class("text-muted-foreground ml-auto text-xs tracking-widest"), "Cmd+S")),
						Div(X.Attr("role", "menuitem"), "Keyboard shortcuts", Span(X.Class("text-muted-foreground ml-auto text-xs tracking-widest"), "Cmd+K")),
					),
					Hr(X.Attr("role", "separator")),
					Div(X.Attr("role", "menuitem"), "GitHub"),
					Div(X.Attr("role", "menuitem"), "Support"),
					Div(X.Attr("role", "menuitem"), X.Attr("aria-disabled", "true"), "API"),
					Hr(X.Attr("role", "separator")),
					Div(X.Attr("role", "menuitem"), "Logout", Span(X.Class("text-muted-foreground ml-auto text-xs tracking-widest"), "Shift+Cmd+Q")),
				),
			}),
			Div(
				X.Id("dropdown-menu-checkboxes"),
				X.Class("dropdown-menu"),
				Button(
					X.Type("button"),
					X.Id("dropdown-menu-checkboxes-trigger"),
					X.Attr("aria-haspopup", "menu"),
					X.Attr("aria-controls", "dropdown-menu-checkboxes-menu"),
					X.Attr("aria-expanded", "false"),
					X.Class("btn-outline"),
					"Checkboxes",
				),
				Div(
					X.Id("dropdown-menu-checkboxes-popover"),
					X.Attr("data-popover", ""),
					X.Attr("aria-hidden", "true"),
					X.Class("min-w-56"),
					Div(
						X.Attr("role", "menu"),
						X.Id("dropdown-menu-checkboxes-menu"),
						X.Attr("aria-labelledby", "dropdown-menu-checkboxes-trigger"),
						Div(
							X.Attr("role", "group"),
							X.Attr("aria-labelledby", "demo-appearance-options"),
							Div(X.Attr("role", "heading"), X.Id("demo-appearance-options"), "Appearance"),
							Div(
								X.Attr("role", "menuitemcheckbox"),
								X.Attr("aria-checked", "true"),
								X.Class("group"),
								lucide.Check(X.Class("invisible group-aria-checked:visible"), X.Attr("aria-hidden", "true"), X.Attr("focusable", "false")),
								"Status Bar",
								Span(X.Class("text-muted-foreground ml-auto text-xs tracking-widest"), "Shift+Cmd+P"),
							),
							Div(
								X.Attr("role", "menuitemcheckbox"),
								X.Attr("aria-checked", "false"),
								X.Attr("aria-disabled", "true"),
								X.Class("group"),
								lucide.Check(X.Class("invisible group-aria-checked:visible"), X.Attr("aria-hidden", "true"), X.Attr("focusable", "false")),
								"Activity Bar",
								Span(X.Class("text-muted-foreground ml-auto text-xs tracking-widest"), "Cmd+B"),
							),
							Div(
								X.Attr("role", "menuitemcheckbox"),
								X.Attr("aria-checked", "false"),
								X.Class("group"),
								lucide.Check(X.Class("invisible group-aria-checked:visible"), X.Attr("aria-hidden", "true"), X.Attr("focusable", "false")),
								"Panel",
								Span(X.Class("text-muted-foreground ml-auto text-xs tracking-widest"), "Cmd+S"),
							),
						),
					),
				),
			),
			Div(
				X.Id("dropdown-menu-radio-group"),
				X.Class("dropdown-menu"),
				Button(
					X.Type("button"),
					X.Id("dropdown-menu-radio-group-trigger"),
					X.Attr("aria-haspopup", "menu"),
					X.Attr("aria-controls", "dropdown-menu-radio-group-menu"),
					X.Attr("aria-expanded", "false"),
					X.Class("btn-outline"),
					"Radio Group",
				),
				Div(
					X.Id("dropdown-menu-radio-group-popover"),
					X.Attr("data-popover", ""),
					X.Attr("aria-hidden", "true"),
					X.Class("min-w-56"),
					Div(
						X.Attr("role", "menu"),
						X.Id("dropdown-menu-radio-group-menu"),
						X.Attr("aria-labelledby", "dropdown-menu-radio-group-trigger"),
						Div(
							X.Attr("role", "group"),
							X.Attr("aria-labelledby", "demo-position-options"),
							Span(X.Attr("role", "heading"), X.Id("demo-position-options"), "Panel Position"),
							Hr(X.Attr("role", "separator")),
							Div(
								X.Attr("role", "menuitemradio"),
								X.Attr("aria-checked", "false"),
								X.Class("group"),
								Div(X.Class("size-4 flex items-center justify-center"), Div(X.Class("size-2 rounded-full bg-foreground invisible group-aria-checked:visible"), X.Attr("aria-hidden", "true"), X.Attr("focusable", "false"))),
								"Status Bar",
								Span(X.Class("text-muted-foreground ml-auto text-xs tracking-widest"), "Shift+Cmd+P"),
							),
							Div(
								X.Attr("role", "menuitemradio"),
								X.Attr("aria-checked", "true"),
								X.Class("group"),
								Div(X.Class("size-4 flex items-center justify-center"), Div(X.Class("size-2 rounded-full bg-foreground invisible group-aria-checked:visible"), X.Attr("aria-hidden", "true"), X.Attr("focusable", "false"))),
								"Activity Bar",
								Span(X.Class("text-muted-foreground ml-auto text-xs tracking-widest"), "Cmd+B"),
							),
							Div(
								X.Attr("role", "menuitemradio"),
								X.Attr("aria-checked", "false"),
								X.Class("group"),
								Div(X.Class("size-4 flex items-center justify-center"), Div(X.Class("size-2 rounded-full bg-foreground invisible group-aria-checked:visible"), X.Attr("aria-hidden", "true"), X.Attr("focusable", "false"))),
								"Panel",
								Span(X.Class("text-muted-foreground ml-auto text-xs tracking-widest"), "Cmd+S"),
							),
						),
					),
				),
			),
			basecoat.DropdownMenu(basecoat.DropdownMenuParams{
				ID: "dropdown-menu-profile",
				Trigger: Fragment(
					Img(
						X.Class("size-8 shrink-0 rounded-full"),
						X.Alt("@hunvreus"),
						X.Src("https://github.com/hunvreus.png"),
					),
					Div(
						X.Class("grid flex-1 text-left text-sm leading-tight"),
						Span(X.Class("truncate font-medium"), "hunvreus"),
						Span(X.Class("text-muted-foreground truncate text-xs"), "hunvreus@example.com"),
					),
					lucide.ChevronDown(X.Class("text-muted-foreground ml-auto")),
				),
				TriggerExtraClasses: "btn-outline h-12 justify-start px-2 md:max-w-[200px]",
				Content: Fragment(
					Div(
						X.Class("flex items-center gap-2 px-1 py-1.5 text-left text-sm"),
						Img(X.Alt("@hunvreus"), X.Src("https://github.com/hunvreus.png"), X.Class("size-8 shrink-0 rounded-full")),
						Div(
							X.Class("grid flex-1 text-left text-sm leading-tight"),
							Span(X.Class("truncate font-medium"), "hunvreus"),
							Span(X.Class("text-muted-foreground truncate text-xs"), "hunvreus@example.com"),
						),
					),
					Hr(X.Attr("role", "separator")),
					Div(X.Attr("role", "menuitem"), lucide.CircleCheck(), "Upgrade to Pro"),
					Hr(X.Attr("role", "separator")),
					Div(X.Attr("role", "menuitem"), "Account"),
					Div(X.Attr("role", "menuitem"), "Billing"),
					Div(X.Attr("role", "menuitem"), "Notifications"),
					Hr(X.Attr("role", "separator")),
					Div(X.Attr("role", "menuitem"), "Signout"),
				),
			}),
			basecoat.DropdownMenu(basecoat.DropdownMenuParams{
				ID:                  "dropdown-menu-actions",
				Trigger:             lucide.Ellipsis(),
				TriggerExtraClasses: "btn-icon-ghost",
				PopoverExtraClasses: "min-w-32",
				Items: []basecoat.DropdownMenuItem{
					{Type: "item", Label: "Edit"},
					{Type: "item", Label: "Share"},
					{Type: "separator"},
					{Type: "item", Label: "Delete", Icon: lucide.Trash2(), Attrs: []HTML{X.Class("text-destructive hover:bg-destructive/10 dark:hover:bg-destructive/20 focus:bg-destructive/10 dark:focus:bg-destructive/20 focus:text-destructive [&_svg]:!text-destructive")}},
				},
			}),
			Raw(`
			<script>
			  (() => {
			    const dropdownMenu = document.querySelector('#dropdown-menu-checkboxes');
			    if (!dropdownMenu) return;
			    const checkboxes = dropdownMenu.querySelectorAll('div[role="menuitemcheckbox"]');
			    checkboxes.forEach(checkbox => {
			      checkbox.addEventListener('click', () => {
			        const isChecked = checkbox.getAttribute('aria-checked') === 'true';
			        checkbox.setAttribute('aria-checked', String(!isChecked));
			      });
			    });
			  })();
			</script>
			<script>
			  (() => {
			    const dropdownMenu = document.querySelector('#dropdown-menu-radio-group');
			    if (!dropdownMenu) return;
			    const radioButtons = dropdownMenu.querySelectorAll('div[role="menuitemradio"]');
			    radioButtons.forEach(radioButton => {
			      radioButton.addEventListener('click', () => {
			        radioButtons.forEach(item => {
			          item.setAttribute('aria-checked', 'false');
			        });
			        radioButton.setAttribute('aria-checked', 'true');
			      });
			    });
			  })();
			</script>
			`),
		),
	)
}

func _formSection() HTML {
	return Div(
		X.Class("p-4"),
		Form(
			X.Class("form grid w-full max-w-sm gap-6"),
			Div(
				X.Class("grid gap-2"),
				Label(X.For("demo-form-text"), "Username"),
				Input(X.Type("text"), X.Id("demo-form-text"), X.Placeholder("hunvreus")),
				P(X.Class("text-muted-foreground text-sm"), "This is your public display name."),
			),
			Div(
				X.Class("grid gap-2"),
				Label(X.For("demo-form-select"), "Email"),
				Select(
					X.Id("demo-form-select"),
					Option(X.Value("bob@example.com"), "m@example.com"),
					Option(X.Value("alice@example.com"), "m@google.com"),
					Option(X.Value("john@example.com"), "m@support.com"),
				),
				P(X.Class("text-muted-foreground text-sm"), "You can manage email addresses in your email settings."),
			),
			Div(
				X.Class("grid gap-2"),
				Label(X.For("demo-form-textarea"), "Bio"),
				Textarea(X.Id("demo-form-textarea"), X.Placeholder("I like to..."), X.Attr("rows", "3")),
				P(X.Class("text-muted-foreground text-sm"), "You can @mention other users and organizations."),
			),
			Div(
				X.Class("flex flex-col gap-3"),
				Label(X.For("demo-form-radio"), "Notify me about..."),
				Fieldset(
					X.Id("demo-form-radio"),
					X.Class("grid gap-3"),
					Label(
						X.Class("font-normal"),
						Input(X.Type("radio"), X.Name("demo-form-radio"), X.Value("1")),
						"All new messages",
					),
					Label(
						X.Class("font-normal"),
						Input(X.Type("radio"), X.Name("demo-form-radio"), X.Value("2")),
						"Direct messages and mentions",
					),
					Label(
						X.Class("font-normal"),
						Input(X.Type("radio"), X.Name("demo-form-radio"), X.Value("3")),
						"Nothing",
					),
				),
			),
			Div(
				X.Class("flex flex-row items-start gap-3 rounded-md border p-4 shadow-xs"),
				Input(X.Type("checkbox"), X.Id("demo-form-checkbox")),
				Div(
					X.Class("flex flex-col gap-1"),
					Label(X.For("demo-form-checkbox"), X.Class("leading-snug"), "Use different settings for my mobile devices"),
					P(X.Class("text-muted-foreground text-sm leading-snug"), "You can manage your mobile notifications in the mobile settings page."),
				),
			),
			Div(
				X.Class("flex flex-col gap-4"),
				Header(
					Label(X.For("demo-form-checkboxes"), X.Class("text-base leading-normal"), "Sidebar"),
					P(X.Class("text-muted-foreground text-sm"), "Select the items you want to display in the sidebar."),
				),
				Fieldset(
					X.Id("demo-form-checkboxes"),
					X.Class("flex flex-col gap-2"),
					Label(
						X.Class("font-normal leading-tight"),
						Input(X.Type("checkbox"), X.Name("demo-form-checkboxes"), X.Value("1"), X.Checked()),
						"Recents",
					),
					Label(
						X.Class("font-normal leading-tight"),
						Input(X.Type("checkbox"), X.Name("demo-form-checkboxes"), X.Value("2"), X.Checked()),
						"Home",
					),
					Label(
						X.Class("font-normal leading-tight"),
						Input(X.Type("checkbox"), X.Name("demo-form-checkboxes"), X.Value("3")),
						"Applications",
					),
					Label(
						X.Class("font-normal leading-tight"),
						Input(X.Type("checkbox"), X.Name("demo-form-checkboxes"), X.Value("4")),
						"Desktop",
					),
					Label(
						X.Class("font-normal leading-tight"),
						Input(X.Type("checkbox"), X.Name("demo-form-checkboxes"), X.Value("5")),
						"Download",
					),
					Label(
						X.Class("font-normal leading-tight"),
						Input(X.Type("checkbox"), X.Name("demo-form-checkboxes"), X.Value("6")),
						"Documents",
					),
				),
			),
			Div(
				X.Class("grid gap-2"),
				Label(X.For("demo-form-date"), "Date of birth"),
				Input(X.Type("date"), X.Id("demo-form-date")),
				P(X.Class("text-muted-foreground text-sm"), "Your date of birth is used to calculate your age."),
			),
			Section(
				X.Class("grid gap-4"),
				H3(X.Class("text-lg font-medium"), "Email Notifications"),
				Div(
					X.Class("gap-2 flex flex-row items-start justify-between rounded-lg border p-4 shadow-xs"),
					Div(
						X.Class("flex flex-col gap-0.5"),
						Label(X.For("demo-form-switch"), X.Class("leading-normal"), "Marketing emails"),
						P(X.Class("text-muted-foreground text-sm"), "Receive emails about new products, features, and more."),
					),
					Input(X.Type("checkbox"), X.Id("demo-form-switch"), X.Attr("role", "switch")),
				),
				Div(
					X.Class("gap-2 flex flex-row items-start justify-between rounded-lg border p-4 shadow-xs"),
					Div(
						X.Class("flex flex-col gap-0.5 opacity-60"),
						Label(X.For("demo-form-switch-disabled"), X.Class("leading-normal"), "Marketing emails"),
						P(X.Class("text-muted-foreground text-sm"), "Receive emails about new products, features, and more."),
					),
					Input(X.Type("checkbox"), X.Id("demo-form-switch-disabled"), X.Attr("role", "switch"), X.Disabled()),
				),
			),
			Button(X.Type("submit"), X.Class("btn"), "Submit"),
		),
	)
}

func _checkboxSection() HTML {
	return Div(
		X.Class("p-4"),
		Div(
			X.Class("flex flex-col gap-6 max-w-lg"),
			Label(
				X.Class("label gap-3"),
				Input(X.Type("checkbox"), X.Class("input")),
				"Accept terms and conditions",
			),
			Div(
				X.Class("flex items-start gap-3"),
				Input(X.Type("checkbox"), X.Id("demo-checkbox-label-and-description"), X.Class("input")),
				Div(
					X.Class("grid gap-2"),
					Label(X.For("demo-checkbox-label-and-description"), X.Class("label"), "Accept terms and conditions"),
					P(X.Class("text-muted-foreground text-sm"), "By clicking this checkbox, you agree to the terms and conditions."),
				),
			),
			Label(
				X.Class("label gap-3"),
				Input(X.Type("checkbox"), X.Class("input"), X.Disabled()),
				"Enable notifications",
			),
			Label(
				X.Class("flex items-start gap-3 border p-3 hover:bg-accent/50 rounded-lg has-[input[type='checkbox']:checked]:border-blue-600 has-[input[type='checkbox']:checked]:bg-blue-50 dark:has-[input[type='checkbox']:checked]:border-blue-900 dark:has-[input[type='checkbox']:checked]:bg-blue-950"),
				Input(X.Type("checkbox"), X.Class("input checked:bg-blue-600 checked:border-blue-600 dark:checked:bg-blue-700 dark:checked:border-blue-700 checked:after:bg-white"), X.Checked()),
				Div(
					X.Class("grid gap-2"),
					H2(X.Class("text-sm leading-none font-medium"), "Enable notifications"),
					P(X.Class("text-muted-foreground text-sm"), "You can enable or disable notifications at any time."),
				),
			),
		),
	)
}

func _inputSection() HTML {
	return Div(
		X.Class("p-4"),
		Div(
			X.Class("flex flex-col gap-y-4"),
			Input(X.Type("text"), X.Class("input"), X.Placeholder("Text")),
			Input(X.Type("text"), X.Class("input"), X.Disabled(), X.Placeholder("Disabled")),
			Input(X.Type("text"), X.Class("input"), X.Attr("aria-invalid", "true"), X.Placeholder("Error")),
			Input(X.Type("email"), X.Class("input"), X.Placeholder("Email")),
			Input(X.Type("password"), X.Class("input"), X.Placeholder("Password")),
			Input(X.Type("number"), X.Class("input"), X.Placeholder("Number")),
			Input(X.Type("file"), X.Class("input")),
			Input(X.Type("tel"), X.Class("input"), X.Placeholder("Tel")),
			Input(X.Type("url"), X.Class("input"), X.Placeholder("URL")),
			Input(X.Type("search"), X.Class("input"), X.Placeholder("Search")),
			Input(X.Type("date"), X.Class("input")),
			Input(X.Type("datetime-local"), X.Class("input")),
			Input(X.Type("month"), X.Class("input")),
			Input(X.Type("week"), X.Class("input")),
			Input(X.Type("time"), X.Class("input")),
		),
	)
}

func _labelSection() HTML {
	return Div(
		X.Class("p-4"),
		Div(
			X.Class("flex flex-col gap-y-4"),
			Div(
				X.Class("grid w-full max-w-sm gap-6"),
				Div(
					X.Class("flex items-center gap-3"),
					Input(X.Type("checkbox"), X.Id("label-demo-terms"), X.Class("input")),
					Label(X.For("label-demo-terms"), X.Class("label"), "Accept terms and conditions"),
				),
				Div(
					X.Class("grid gap-3"),
					Label(X.For("label-demo-text"), X.Class("label"), "Username"),
					Input(X.Type("text"), X.Id("label-demo-text"), X.Class("input"), X.Placeholder("Username")),
				),
				Div(
					X.Class("grid gap-3"),
					Label(X.For("label-demo-disabled"), X.Class("label"), "Disabled"),
					Input(X.Type("text"), X.Id("label-demo-disabled"), X.Class("peer input"), X.Placeholder("Disabled"), X.Disabled()),
				),
				Div(
					X.Class("grid gap-3"),
					Label(X.For("label-demo-textarea"), X.Class("label"), "Message"),
					Textarea(X.Id("label-demo-textarea"), X.Class("textarea"), X.Placeholder("Message")),
				),
			),
		),
	)
}

func _paginationSection() HTML {
	return Div(
		X.Class("p-4"),
		Div(
			X.Class("inline-flex"),
			Nav(
				X.Attr("role", "navigation"),
				X.Attr("aria-label", "pagination"),
				X.Class("mx-auto flex w-full justify-center"),
				Ul(
					X.Class("flex flex-row items-center gap-1"),
					Li(A(X.Href("#"), X.Class("btn-ghost"), lucide.ChevronLeft(), " Previous")),
					Li(A(X.Href("#"), X.Class("btn-ghost size-9"), "1")),
					Li(A(X.Href("#"), X.Class("btn-outline size-9"), "2")),
					Li(A(X.Href("#"), X.Class("btn-ghost size-9"), "3")),
					Li(A(X.Href("#"), X.Class("btn-icon-ghost"), lucide.Ellipsis())),
					Li(A(X.Href("#"), X.Class("btn-ghost"), "Next ", lucide.ChevronRight())),
				),
			),
		),
	)
}

func _popoverSection() HTML {
	return Div(
		X.Class("p-4"),
		basecoat.Popover(basecoat.PopoverParams{
			ID:                  "demo-popover",
			Trigger:             "Open popover",
			TriggerExtraClasses: "btn-outline",
			PopoverExtraClasses: "w-80",
			Content: Div(
				X.Class("grid gap-4"),
				Header(
					X.Class("grid gap-1.5"),
					H4(X.Class("leading-none font-medium"), "Dimensions"),
					P(X.Class("text-muted-foreground text-sm"), "Set the dimensions for the layer."),
				),
				Form(
					X.Class("form grid gap-2"),
					Div(
						X.Class("grid grid-cols-3 items-center gap-4"),
						Label(X.For("demo-popover-width"), "Width"),
						Input(X.Type("text"), X.Id("demo-popover-width"), X.Value("100%"), X.Class("col-span-2 h-8")),
					),
					Div(
						X.Class("grid grid-cols-3 items-center gap-4"),
						Label(X.For("demo-popover-max-width"), "Max. width"),
						Input(X.Type("text"), X.Id("demo-popover-max-width"), X.Value("300px"), X.Class("col-span-2 h-8")),
					),
					Div(
						X.Class("grid grid-cols-3 items-center gap-4"),
						Label(X.For("demo-popover-height"), "Height"),
						Input(X.Type("text"), X.Id("demo-popover-height"), X.Value("25px"), X.Class("col-span-2 h-8")),
					),
					Div(
						X.Class("grid grid-cols-3 items-center gap-4"),
						Label(X.For("demo-popover-max-height"), "Max. height"),
						Input(X.Type("text"), X.Id("demo-popover-max-height"), X.Value("none"), X.Class("col-span-2 h-8")),
					),
				),
			),
		}),
	)
}

func _radioGroupSection() HTML {
	return Div(
		X.Class("p-4"),
		Div(
			X.Class("flex flex-col gap-y-6"),
			Fieldset(
				X.Class("grid gap-3"),
				Label(X.Class("label"), Input(X.Type("radio"), X.Name("demo-radio-group"), X.Value("default"), X.Class("input")), "Default"),
				Label(X.Class("label"), Input(X.Type("radio"), X.Name("demo-radio-group"), X.Value("comfortable"), X.Class("input")), "Comfortable"),
				Label(X.Class("label"), Input(X.Type("radio"), X.Name("demo-radio-group"), X.Value("compact"), X.Class("input")), "Compact"),
			),
			Fieldset(
				X.Class("grid gap-3 max-w-sm"),
				Label(
					X.Class("label gap-3 items-start hover:bg-accent/50 rounded-lg border p-4 has-[input[type='radio']:checked]:border-green-600 has-[input[type='radio']:checked]:bg-green-50 dark:has-[input[type='radio']:checked]:border-green-900 dark:has-[input[type='radio']:checked]:bg-green-950"),
					Input(X.Type("radio"), X.Name("demo-radio-group-styled"), X.Value("starter"), X.Class("input checked:bg-green-600 checked:border-green-600 dark:checked:bg-input/30 checked:before:bg-background dark:checked:before:bg-primary")),
					Div(
						X.Class("grid gap-1 font-normal"),
						H2(X.Class("font-medium"), "Starter Plan"),
						P(X.Class("text-muted-foreground leading-snug"), "Perfect for small businesses getting started with our platform"),
					),
				),
				Label(
					X.Class("label gap-3 items-start hover:bg-accent/50 rounded-lg border p-4 has-[input[type='radio']:checked]:border-green-600 has-[input[type='radio']:checked]:bg-green-50 dark:has-[input[type='radio']:checked]:border-green-900 dark:has-[input[type='radio']:checked]:bg-green-950"),
					Input(X.Type("radio"), X.Name("demo-radio-group-styled"), X.Value("pro"), X.Class("input checked:bg-green-600 checked:border-green-600 dark:checked:bg-input/30 checked:before:bg-background dark:checked:before:bg-primary")),
					Div(
						X.Class("grid gap-1 font-normal"),
						H2(X.Class("font-medium"), "Pro Plan"),
						P(X.Class("text-muted-foreground leading-snug"), "Advanced features for growing businesses with higher demands"),
					),
				),
			),
		),
	)
}

func _selectSection() HTML {
	return Div(
		X.Class("p-4"),
		Div(
			X.Class("flex flex-col gap-4"),
			Div(
				X.Class("flex flex-wrap items-center gap-2 md:flex-row"),
				Select(
					X.Class("select w-[180px]"),
					Optgroup(X.Attr("label", "Fruits"), Option("Apple"), Option("Banana"), Option("Blueberry")),
					Optgroup(X.Attr("label", "Grapes"), Option("Pineapple")),
				),
				Select(X.Class("select w-[180px]"), X.Disabled(), Option("Disabled")),
			),
			Div(
				X.Class("flex flex-wrap items-center gap-4"),
				basecoat.Select(basecoat.SelectParams{
					ID:                  "select-default",
					Selected:            []string{"blueberry"},
					TriggerExtraClasses: "w-[180px]",
					Items: []basecoat.SelectItem{
						{Type: "group", Label: "Fruits", Items: []basecoat.SelectItem{{Type: "item", Value: "apple", Label: "Apple"}, {Type: "item", Value: "banana", Label: "Banana"}, {Type: "item", Value: "blueberry", Label: "Blueberry"}}},
						{Type: "group", Label: "Grapes", Items: []basecoat.SelectItem{{Type: "item", Value: "pineapple", Label: "Pineapple"}}},
					},
				}),
				basecoat.Select(basecoat.SelectParams{
					ID:                  "select-scrollbar",
					Selected:            []string{"item-0"},
					TriggerExtraClasses: "w-[180px]",
					ListboxExtraClasses: "scrollbar overflow-y-auto max-h-64",
					Items: []basecoat.SelectItem{
						{Type: "item", Value: "item-0", Label: "Item 0"}, {Type: "item", Value: "item-1", Label: "Item 1"}, {Type: "item", Value: "item-2", Label: "Item 2"}, {Type: "item", Value: "item-3", Label: "Item 3"}, {Type: "item", Value: "item-4", Label: "Item 4"},
						{Type: "item", Value: "item-5", Label: "Item 5"}, {Type: "item", Value: "item-6", Label: "Item 6"}, {Type: "item", Value: "item-7", Label: "Item 7"}, {Type: "item", Value: "item-8", Label: "Item 8"}, {Type: "item", Value: "item-9", Label: "Item 9"},
						{Type: "item", Value: "item-10", Label: "Item 10"}, {Type: "item", Value: "item-11", Label: "Item 11"}, {Type: "item", Value: "item-12", Label: "Item 12"}, {Type: "item", Value: "item-13", Label: "Item 13"}, {Type: "item", Value: "item-14", Label: "Item 14"},
						{Type: "item", Value: "item-15", Label: "Item 15"}, {Type: "item", Value: "item-16", Label: "Item 16"}, {Type: "item", Value: "item-17", Label: "Item 17"}, {Type: "item", Value: "item-18", Label: "Item 18"}, {Type: "item", Value: "item-19", Label: "Item 19"},
					},
				}),
				basecoat.Select(basecoat.SelectParams{
					ID:                  "select-disabled",
					Selected:            []string{"disabled"},
					TriggerExtraClasses: "w-[180px]",
					TriggerAttrs:        []HTML{X.Attr("disabled", "disabled")},
					Items:               []basecoat.SelectItem{{Type: "item", Value: "disabled", Label: "Disabled"}},
				}),
				basecoat.Select(basecoat.SelectParams{
					ID:                  "select-with-icon",
					Selected:            []string{"bar"},
					Name:                "chart-type",
					TriggerExtraClasses: "w-[180px]",
					Items: []basecoat.SelectItem{
						{Type: "item", Value: "bar", Label: Span(X.Class("flex items-center gap-2"), lucide.ChartColumn(X.Class("text-muted-foreground")), "Bar")},
						{Type: "item", Value: "line", Label: Span(X.Class("flex items-center gap-2"), lucide.ChartLine(X.Class("text-muted-foreground")), "Line")},
						{Type: "item", Value: "pie", Label: Span(X.Class("flex items-center gap-2"), lucide.ChartPie(X.Class("text-muted-foreground")), "Pie")},
					},
				}),
			),
		),
	)
}

func _skeletonSection() HTML {
	return Div(
		X.Class("p-4"),
		Div(
			X.Class("flex flex-col gap-4"),
			Div(
				X.Class("flex items-center gap-4"),
				Div(X.Class("bg-accent animate-pulse size-10 shrink-0 rounded-full")),
				Div(
					X.Class("grid gap-2"),
					Div(X.Class("bg-accent animate-pulse rounded-md h-4 w-[150px]")),
					Div(X.Class("bg-accent animate-pulse rounded-md h-4 w-[100px]")),
				),
			),
			Div(
				X.Class("flex max-sm:flex-col gap-4 w-full"),
				Div(
					X.Class("card w-full @md:w-auto @md:min-w-sm"),
					Header(
						Div(X.Class("bg-accent animate-pulse rounded-md h-4 w-2/3")),
						Div(X.Class("bg-accent animate-pulse rounded-md h-4 w-1/2")),
					),
					Section(Div(X.Class("bg-accent animate-pulse rounded-md aspect-square w-full"))),
				),
				Div(
					X.Class("card w-full @md:w-auto @md:min-w-sm"),
					Header(
						Div(X.Class("bg-accent animate-pulse rounded-md h-4 w-2/3")),
						Div(X.Class("bg-accent animate-pulse rounded-md h-4 w-1/2")),
					),
					Section(Div(X.Class("bg-accent animate-pulse rounded-md aspect-square w-full"))),
				),
			),
		),
	)
}

func _sliderSection() HTML {
	return Div(
		X.Class("p-4"),
		Div(
			X.Class("max-w-sm"),
			Input(
				X.Type("range"),
				X.Class("input w-full"),
				X.Attr("min", "0"),
				X.Attr("max", "27"),
				X.Attr("value", "12"),
				X.Attr("style", "--slider-value: 44.44444444444444%;"),
			),
		),
		Raw(`
		<script>
		  (() => {
		    const sliders = document.querySelectorAll('input[type="range"].input');
		    if (!sliders) return;

		    const updateSlider = (el) => {
		      const min = parseFloat(el.min || 0);
		      const max = parseFloat(el.max || 100);
		      const value = parseFloat(el.value);
		      const percent = (max === min) ? 0 : ((value - min) / (max - min)) * 100;
		      el.style.setProperty('--slider-value', percent + '%');
		    };

		    sliders.forEach(slider => {
		      updateSlider(slider);
		      slider.addEventListener('input', (event) => updateSlider(event.target));
		    });
		  })();
		</script>
		`),
	)
}

func _switchSection() HTML {
	return Div(
		X.Class("p-4"),
		Div(
			X.Class("inline-flex flex-col gap-y-6"),
			Label(
				X.Class("label"),
				Input(X.Type("checkbox"), X.Name("switch"), X.Attr("role", "switch"), X.Class("input")),
				"Airplane Mode",
			),
			Label(
				X.Class("label"),
				Input(X.Type("checkbox"), X.Name("switch"), X.Attr("role", "switch"), X.Class("input checked:bg-blue-500 dark:checked:bg-blue-600"), X.Checked()),
				"Bluetooth",
			),
			Label(
				X.Class("label gap-6 leading-none border rounded-lg p-4 has-[input[type='checkbox']:checked]:border-blue-600"),
				Div(
					X.Class("grid gap-1"),
					H2(X.Class("font-medium text-sm"), "Share across devices"),
					P(X.Class("text-muted-foreground text-sm"), "Focus is shared across devices, and turns off when you leave the app."),
				),
				Input(X.Type("checkbox"), X.Name("switch"), X.Attr("role", "switch"), X.Class("input checked:bg-blue-500 dark:checked:bg-blue-600")),
			),
		),
	)
}

func _tableSection() HTML {
	return Div(
		X.Class("p-4"),
		Div(
			X.Class("relative w-full overflow-x-auto"),
			Table(
				X.Class("table"),
				Caption("A list of your recent invoices."),
				Thead(
					Tr(Th("Invoice"), Th("Status"), Th("Method"), Th("Amount")),
				),
				Tbody(
					Tr(Td(X.Class("font-medium"), "INV001"), Td("Paid"), Td("Credit Card"), Td(X.Class("text-right"), "$250.00")),
					Tr(Td(X.Class("font-medium"), "INV002"), Td("Pending"), Td("PayPal"), Td(X.Class("text-right"), "$150.00")),
					Tr(Td(X.Class("font-medium"), "INV003"), Td("Unpaid"), Td("Bank Transfer"), Td(X.Class("text-right"), "$350.00")),
					Tr(Td(X.Class("font-medium"), "INV004"), Td("Paid"), Td("Paypal"), Td(X.Class("text-right"), "$450.00")),
					Tr(Td(X.Class("font-medium"), "INV005"), Td("Paid"), Td("Credit Card"), Td(X.Class("text-right"), "$550.00")),
					Tr(Td(X.Class("font-medium"), "INV006"), Td("Pending"), Td("Bank Transfer"), Td(X.Class("text-right"), "$200.00")),
					Tr(Td(X.Class("font-medium"), "INV007"), Td("Unpaid"), Td("Credit Card"), Td(X.Class("text-right"), "$300.00")),
				),
				Tfoot(
					Tr(
						Td(X.Attr("colspan", "3"), "Total"),
						Td(X.Class("text-right"), "$2,500.00"),
					),
				),
			),
		),
	)
}

func _tabsSection() HTML {
	return Div(
		X.Class("p-4"),
		Div(
			X.Class("flex flex-col gap-6"),
			basecoat.Tabs(basecoat.TabsParams{
				ID:                  "demo-tabs-with-panels",
				MainExtraClasses:    "max-w-[300px]",
				TablistExtraClasses: "w-full",
				Tabsets: []basecoat.TabsTabset{
					{
						Tab: "Account",
						Panel: Div(
							X.Class("card"),
							Header(
								H2("Account"),
								P("Make changes to your account here. Click save when you're done."),
							),
							Section(
								Form(
									X.Class("form grid gap-6"),
									Div(X.Class("grid gap-3"), Label(X.For("demo-tabs-account-name"), "Name"), Input(X.Type("text"), X.Id("demo-tabs-account-name"), X.Value("Pedro Duarte"))),
									Div(X.Class("grid gap-3"), Label(X.For("demo-tabs-account-username"), "Username"), Input(X.Type("text"), X.Id("demo-tabs-account-username"), X.Value("@peduarte"))),
								),
							),
							Footer(Button(X.Type("button"), X.Class("btn"), "Save changes")),
						),
					},
					{
						Tab: "Password",
						Panel: Div(
							X.Class("card"),
							Header(
								H2("Password"),
								P("Change your password here. After saving, you'll be logged out."),
							),
							Section(
								Form(
									X.Class("form grid gap-6"),
									Div(X.Class("grid gap-3"), Label(X.For("demo-tabs-password-current"), "Current password"), Input(X.Type("password"), X.Id("demo-tabs-password-current"))),
									Div(X.Class("grid gap-3"), Label(X.For("demo-tabs-password-new"), "New password"), Input(X.Type("password"), X.Id("demo-tabs-password-new"))),
								),
							),
							Footer(Button(X.Type("button"), X.Class("btn"), "Save Password")),
						),
					},
				},
			}),
			basecoat.Tabs(basecoat.TabsParams{
				ID:      "demo-tabs-without-panels",
				Tabsets: []basecoat.TabsTabset{{Tab: "Home"}, {Tab: "Settings"}},
			}),
			basecoat.Tabs(basecoat.TabsParams{
				ID: "demo-tabs-disabled",
				Tabsets: []basecoat.TabsTabset{
					{Tab: "Home"},
					{Tab: "Disabled", TabAttrs: []HTML{X.Attr("disabled", "disabled")}},
				},
			}),
			basecoat.Tabs(basecoat.TabsParams{
				ID: "demo-tabs-with-icons",
				Tabsets: []basecoat.TabsTabset{
					{Tab: Fragment(lucide.Newspaper(), Span("Preview"))},
					{Tab: Fragment(lucide.Code(), Span("Code"))},
				},
			}),
		),
	)
}

func _textareaSection() HTML {
	return Div(
		X.Class("p-4"),
		Div(
			X.Class("flex flex-col gap-y-10"),
			Textarea(X.Class("textarea"), X.Placeholder("Type your message here")),
			Textarea(X.Class("textarea"), X.Placeholder("Type your message here"), X.Attr("aria-invalid", "true")),
			Div(
				X.Class("grid gap-3"),
				Label(X.For("textarea-demo-label"), X.Class("label"), "Label"),
				Textarea(X.Id("textarea-demo-label"), X.Class("textarea"), X.Placeholder("Type your message here")),
			),
			Div(
				X.Class("grid gap-3"),
				Label(X.For("textarea-demo-label-and-description"), X.Class("label"), "With label and description"),
				Textarea(X.Id("textarea-demo-label-and-description"), X.Class("textarea"), X.Placeholder("Type your message here")),
				P(X.Class("text-muted-foreground text-sm"), "Type your message and press enter to send."),
			),
			Div(
				X.Class("grid gap-3"),
				Label(X.For("textarea-demo-disabled"), X.Class("label"), "Disabled"),
				Textarea(X.Id("textarea-demo-disabled"), X.Class("textarea"), X.Placeholder("Type your message here"), X.Disabled()),
			),
		),
	)
}

func _toastSection() HTML {
	return Div(
		X.Class("p-4"),
		Div(
			X.Class("flex flex-wrap items-center gap-2"),
			Button(X.Class("btn-outline"), X.Attr("hx-trigger", "click"), X.Attr("hx-get", "/fragments/toast/success"), X.Attr("hx-target", "#toaster"), X.Attr("hx-swap", "beforeend"), "Success"),
			Button(X.Class("btn-outline"), X.Attr("hx-trigger", "click"), X.Attr("hx-get", "/fragments/toast/error"), X.Attr("hx-target", "#toaster"), X.Attr("hx-swap", "beforeend"), "Error"),
			Button(X.Class("btn-outline"), X.Attr("hx-trigger", "click"), X.Attr("hx-get", "/fragments/toast/info"), X.Attr("hx-target", "#toaster"), X.Attr("hx-swap", "beforeend"), "Info"),
			Button(X.Class("btn-outline"), X.Attr("hx-trigger", "click"), X.Attr("hx-get", "/fragments/toast/warning"), X.Attr("hx-target", "#toaster"), X.Attr("hx-swap", "beforeend"), "Warning"),
		),
	)
}

func _tooltipSection() HTML {
	return Div(
		X.Class("p-4"),
		Div(
			X.Class("flex flex-wrap items-center gap-4"),
			Button(X.Type("button"), X.Class("btn-outline"), X.Attr("data-tooltip", "Top tooltip"), "Top"),
			Button(X.Type("button"), X.Class("btn-outline"), X.Attr("data-tooltip", "Right tooltip"), X.Attr("data-side", "right"), "Right"),
			Button(X.Type("button"), X.Class("btn-outline"), X.Attr("data-tooltip", "Bottom tooltip"), X.Attr("data-side", "bottom"), "Bottom"),
			Button(X.Type("button"), X.Class("btn-outline"), X.Attr("data-tooltip", "Left tooltip"), X.Attr("data-side", "left"), "Left"),
		),
	)
}

func _avatarSection() HTML {
	return Div(
		X.Class("p-4"),
		Div(
			X.Class("flex flex-row flex-wrap items-center gap-4"),
			Img(
				X.Class("size-8 shrink-0 object-cover rounded-full"),
				X.Alt("@hunvreus"),
				X.Src("https://github.com/hunvreus.png"),
			),
			Span(
				X.Class("size-8 shrink-0 bg-muted flex items-center justify-center rounded-full"),
				"CN",
			),
			Img(
				X.Class("size-12 shrink-0 object-cover rounded-full"),
				X.Alt("@hunvreus"),
				X.Src("https://github.com/hunvreus.png"),
			),
			Img(
				X.Class("size-8 shrink-0 object-cover rounded-lg"),
				X.Alt("@hunvreus"),
				X.Src("https://github.com/hunvreus.png"),
			),
			Div(
				X.Class("flex -space-x-2 [&_img]:ring-background [&_img]:ring-2 [&_img]:grayscale [&_img]:size-8 [&_img]:shrink-0 [&_img]:object-cover [&_img]:rounded-full"),
				Img(X.Alt("@hunvreus"), X.Src("https://github.com/hunvreus.png")),
				Img(X.Alt("@shadcn"), X.Src("https://github.com/shadcn.png")),
				Img(X.Alt("@adamwathan"), X.Src("https://github.com/adamwathan.png")),
			),
			Div(
				X.Class("flex -space-x-2 [&_img]:ring-background [&_img]:ring-2 [&_img]:grayscale [&_img]:size-12 [&_img]:shrink-0 [&_img]:object-cover [&_img]:rounded-full"),
				Img(X.Alt("@hunvreus"), X.Src("https://github.com/hunvreus.png")),
				Img(X.Alt("@shadcn"), X.Src("https://github.com/shadcn.png")),
				Img(X.Alt("@adamwathan"), X.Src("https://github.com/adamwathan.png")),
			),
			Div(
				X.Class("flex -space-x-2 hover:space-x-1 [&_img]:ring-background [&_img]:size-12 [&_img]:shrink-0 [&_img]:object-cover [&_img]:rounded-full [&_img]:ring-2 [&_img]:grayscale [&_img]:transition-all [&_img]:ease-in-out [&_img]:duration-300"),
				Img(X.Alt("@hunvreus"), X.Src("https://github.com/hunvreus.png")),
				Img(X.Alt("@shadcn"), X.Src("https://github.com/shadcn.png")),
				Img(X.Alt("@adamwathan"), X.Src("https://github.com/adamwathan.png")),
			),
		),
	)
}

func _badgeSection() HTML {
	return Div(
		X.Class("p-4"),
		Div(
			X.Class("flex flex-col gap-2"),
			Div(
				X.Class("flex flex-wrap items-center gap-2 md:flex-row"),
				Span(X.Class("badge"), "Primary"),
				Span(X.Class("badge-secondary"), "Secondary"),
				Span(X.Class("badge-outline"), "Outline"),
				Span(X.Class("badge-destructive"), "Destructive"),
				Span(
					X.Class("badge-outline"),
					lucide.Check(),
					"Badge",
				),
				Span(
					X.Class("badge-destructive"),
					lucide.CircleAlert(),
					"Alert",
				),
				Span(X.Class("badge rounded-full min-w-5 px-1"), "8"),
				Span(X.Class("badge-destructive rounded-full min-w-5 px-1"), "99"),
				Span(X.Class("badge-outline rounded-full min-w-5 px-1 font-mono tabular-nums"), "20+"),
			),
			Div(
				X.Class("flex flex-wrap items-center gap-2 md:flex-row"),
				A(
					X.Href("#"),
					X.Class("badge"),
					"Link",
					lucide.ArrowRight(),
				),
				A(
					X.Href("#"),
					X.Class("badge-secondary"),
					"Link",
					lucide.ArrowRight(),
				),
				A(
					X.Href("#"),
					X.Class("badge-destructive"),
					"Link",
					lucide.ArrowRight(),
				),
				A(
					X.Href("#"),
					X.Class("badge-outline"),
					"Link",
					lucide.ArrowRight(),
				),
			),
		),
	)
}

func _breadcrumbSection() HTML {
	return Div(
		X.Class("p-4"),
		Ol(
			X.Class("text-muted-foreground flex flex-wrap items-center gap-1.5 text-sm break-words sm:gap-2.5"),
			Li(
				X.Class("inline-flex items-center gap-1.5"),
				A(X.Href("#"), X.Class("hover:text-foreground transition-colors"), "Home"),
			),
			Li(lucide.ChevronRight(X.Class("size-3.5"))),
			Li(
				X.Class("inline-flex items-center gap-1.5"),
				Div(
					X.Id("demo-breadcrumb-menu"),
					X.Class("dropdown-menu"),
					X.Attr("data-dropdown-menu-initialized", "true"),
					Button(
						X.Type("button"),
						X.Id("demo-breadcrumb-menu-trigger"),
						X.Attr("aria-haspopup", "menu"),
						X.Attr("aria-controls", "demo-breadcrumb-menu-menu"),
						X.Attr("aria-expanded", "false"),
						X.Class("flex size-9 items-center justify-center h-4 w-4 hover:text-foreground cursor-pointer"),
						lucide.Ellipsis(),
					),
					Div(
						X.Id("demo-breadcrumb-menu-popover"),
						X.Attr("data-popover", ""),
						X.Attr("aria-hidden", "true"),
						X.Class("p-1"),
						Div(
							X.Attr("role", "menu"),
							X.Id("demo-breadcrumb-menu-menu"),
							X.Attr("aria-labelledby", "demo-breadcrumb-menu-trigger"),
							Div(X.Id("demo-breadcrumb-menu-items-1"), X.Attr("role", "menuitem"), "Documentation"),
							Div(X.Id("demo-breadcrumb-menu-items-2"), X.Attr("role", "menuitem"), "Themes"),
							Div(X.Id("demo-breadcrumb-menu-items-3"), X.Attr("role", "menuitem"), "GitHub"),
						),
					),
				),
			),
			Li(lucide.ChevronRight(X.Class("size-3.5"))),
			Li(
				X.Class("inline-flex items-center gap-1.5"),
				A(X.Href("#"), X.Class("hover:text-foreground transition-colors"), "Components"),
			),
			Li(lucide.ChevronRight(X.Class("size-3.5"))),
			Li(
				X.Class("inline-flex items-center gap-1.5"),
				Span(X.Class("text-foreground font-normal"), "Breadcrumb"),
			),
		),
	)
}

func _buttonSection() HTML {
	return Div(
		X.Class("p-4"),
		Div(
			X.Class("flex flex-col gap-6"),
			Div(
				X.Class("flex flex-wrap items-center gap-2 md:flex-row"),
				Button(X.Type("button"), X.Class("btn-primary"), "Primary"),
				Button(X.Type("button"), X.Class("btn-outline"), "Outline"),
				Button(X.Type("button"), X.Class("btn-ghost"), "Ghost"),
				Button(X.Type("button"), X.Class("btn-destructive"), lucide.Send(), "Danger"),
				Button(X.Type("button"), X.Class("btn-secondary"), "Secondary"),
				Button(X.Type("button"), X.Class("btn-link"), "Link"),
				Button(X.Type("button"), X.Class("btn-outline"), lucide.Send(), "Send"),
				Button(X.Type("button"), X.Class("btn-outline"), "Learn more", lucide.ArrowRight()),
				Button(X.Type("button"), X.Class("btn-outline"), X.Disabled(), lucide.Loader(X.Class("animate-spin")), "Loading"),
			),
			Div(
				X.Class("flex flex-wrap items-center gap-2 md:flex-row"),
				Button(X.Type("button"), X.Class("btn-sm-primary"), "Primary"),
				Button(X.Type("button"), X.Class("btn-sm-outline"), "Outline"),
				Button(X.Type("button"), X.Class("btn-sm-ghost"), "Ghost"),
				Button(X.Type("button"), X.Class("btn-sm-destructive"), "Danger"),
				Button(X.Type("button"), X.Class("btn-sm-secondary"), "Secondary"),
				Button(X.Type("button"), X.Class("btn-sm-link"), "Link"),
				Button(X.Type("button"), X.Class("btn-sm-outline"), lucide.Send(), "Send"),
				Button(X.Type("button"), X.Class("btn-sm-outline"), "Learn more", lucide.ArrowRight()),
				Button(X.Type("button"), X.Class("btn-sm-outline"), X.Disabled(), lucide.Loader(X.Class("animate-spin")), "Loading"),
			),
			Div(
				X.Class("flex flex-wrap items-center gap-2 md:flex-row"),
				Button(X.Type("button"), X.Class("btn-lg-primary"), "Primary"),
				Button(X.Type("button"), X.Class("btn-lg-outline"), "Outline"),
				Button(X.Type("button"), X.Class("btn-lg-ghost"), "Ghost"),
				Button(X.Type("button"), X.Class("btn-lg-destructive"), lucide.Send(), "Danger"),
				Button(X.Type("button"), X.Class("btn-lg-secondary"), "Secondary"),
				Button(X.Type("button"), X.Class("btn-lg-link"), "Link"),
				Button(X.Type("button"), X.Class("btn-lg-outline"), lucide.Send(), "Send"),
				Button(X.Type("button"), X.Class("btn-lg-outline"), "Learn more", lucide.ArrowRight()),
				Button(X.Type("button"), X.Class("btn-lg-outline"), X.Disabled(), lucide.Loader(X.Class("animate-spin")), "Loading"),
			),
			Div(
				X.Class("flex flex-wrap items-center gap-2 md:flex-row"),
				Button(X.Type("button"), X.Class("btn-icon-primary"), lucide.Download()),
				Button(X.Type("button"), X.Class("btn-icon-secondary"), lucide.Upload()),
				Button(X.Type("button"), X.Class("btn-icon-outline"), lucide.ArrowRight()),
				Button(X.Type("button"), X.Class("btn-icon-ghost"), lucide.Ellipsis()),
				Button(X.Type("button"), X.Class("btn-icon-destructive"), lucide.Trash2()),
				Button(X.Type("button"), X.Class("btn-icon-outline"), X.Disabled(), lucide.Loader(X.Class("animate-spin"))),
			),
		),
	)
}

func _cardSection() HTML {
	return Div(
		X.Class("p-4"),
		Div(
			X.Class("flex flex-col items-start gap-4"),
			Div(
				X.Class("card w-full max-w-sm"),
				Header(
					H2("Login to your account"),
					P("Enter your details below to login to your account"),
				),
				Section(
					Form(
						X.Class("form grid gap-6"),
						Div(
							X.Class("grid gap-2"),
							Label(X.For("demo-card-form-email"), "Email"),
							Input(X.Type("email"), X.Id("demo-card-form-email")),
						),
						Div(
							X.Class("grid gap-2"),
							Div(
								X.Class("flex items-center gap-2"),
								Label(X.For("demo-card-form-password"), "Password"),
								A(X.Href("#"), X.Class("ml-auto inline-block text-sm underline-offset-4 hover:underline"), "Forgot your password?"),
							),
							Input(X.Type("password"), X.Id("demo-card-form-password")),
						),
					),
				),
				Footer(
					X.Class("flex flex-col items-center gap-2"),
					Button(X.Type("button"), X.Class("btn w-full"), "Login"),
					Button(X.Type("button"), X.Class("btn-outline w-full"), "Login with Google"),
					P(
						X.Class("mt-4 text-center text-sm"),
						"Don't have an account? ",
						A(X.Href("#"), X.Class("underline-offset-4 hover:underline"), "Sign up"),
					),
				),
			),
			Div(
				X.Class("card"),
				Header(
					H2("Meeting Notes"),
					P("Transcript from the meeting with the client."),
				),
				Section(
					X.Class("text-sm"),
					P("Client requested dashboard redesign with focus on mobile responsiveness."),
					Ol(
						X.Class("mt-4 flex list-decimal flex-col gap-2 pl-6"),
						Li("New analytics widgets for daily/weekly metrics"),
						Li("Simplified navigation menu"),
						Li("Dark mode support"),
						Li("Timeline: 6 weeks"),
						Li("Follow-up meeting scheduled for next Tuesday"),
					),
				),
				Footer(
					X.Class("flex items-center"),
					Div(
						X.Class("flex -space-x-2 [&_img]:ring-background [&_img]:ring-2 [&_img]:grayscale [&_img]:size-8 [&_img]:shrink-0 [&_img]:object-cover [&_img]:rounded-full"),
						Img(X.Alt("@hunvreus"), X.Src("https://github.com/hunvreus.png")),
						Img(X.Alt("@shadcn"), X.Src("https://github.com/shadcn.png")),
						Img(X.Alt("@adamwathan"), X.Src("https://github.com/adamwathan.png")),
					),
				),
			),
			Div(
				X.Class("card"),
				Header(
					H2("Is this an image?"),
					P("This is a card with an image."),
				),
				Section(
					X.Class("px-0"),
					Img(
						X.Alt("Photo by Drew Beamer"),
						X.Class("aspect-video object-cover"),
						X.Src("https://images.unsplash.com/photo-1588345921523-c2dcdb7f1dcd?w=1080&q=75"),
					),
				),
				Footer(
					X.Class("flex items-center gap-2"),
					Span(X.Class("badge-outline"), "1 bed"),
					Span(X.Class("badge-outline"), "2 bath"),
					Span(X.Class("badge-outline"), "350m2"),
					Span(X.Class("ml-auto font-medium tabular-nums"), "$135,000"),
				),
			),
			Div(
				X.Class("flex w-full flex-wrap items-start gap-8 md:*:[.card]:basis-1/4"),
				Div(X.Class("card"), Section("Content Only")),
				Div(
					X.Class("card"),
					Header(
						H2("Header Only"),
						P("This is a card with a header and a description."),
					),
				),
				Div(
					X.Class("card"),
					Header(
						H2("Header and Content"),
						P("This is a card with a header and a content."),
					),
					Section("Content only."),
				),
				Div(X.Class("card"), Footer("Footer Only")),
				Div(
					X.Class("card"),
					Header(
						H2("Header + Footer"),
						P("This is a card with a header and a footer."),
					),
					Footer("Footer"),
				),
				Div(X.Class("card"), Section("Content"), Footer("Footer")),
				Div(
					X.Class("card"),
					Header(
						H2("Header + Content + Footer"),
						P("This is a card with a header, content and footer."),
					),
					Section("Content"),
					Footer("Footer"),
				),
			),
		),
	)
}

func _comboboxSection() HTML {
	return Div(
		X.Class("p-4"),
		Div(
			X.Class("flex flex-wrap items-start gap-4"),
			basecoat.Select(basecoat.SelectParams{
				ID:                  "demo-combobox-frameworks",
				IsCombobox:          true,
				Selected:            []string{"Next.js"},
				PopoverExtraClasses: "w-48",
				ListboxAttrs: []HTML{
					X.Attr("data-empty", "No framework found."),
				},
				Items: []basecoat.SelectItem{
					{Type: "item", Value: "Next.js", Label: "Next.js"},
					{Type: "item", Value: "SvelteKit", Label: "SvelteKit"},
					{Type: "item", Value: "Nuxt.js", Label: "Nuxt.js"},
					{Type: "item", Value: "Remix", Label: "Remix"},
					{Type: "item", Value: "Astro", Label: "Astro"},
				},
			}),
			basecoat.Select(basecoat.SelectParams{
				ID:                  "demo-combobox-timezones",
				IsCombobox:          true,
				Selected:            []string{"America/New_York"},
				PopoverExtraClasses: "w-72",
				ListboxExtraClasses: "scrollbar overflow-y-auto max-h-64",
				ListboxAttrs: []HTML{
					X.Data("empty", "No timezone found."),
				},
				Items: []basecoat.SelectItem{
					{
						Type:  "group",
						ID:    "demo-combobox-timezones-group-0",
						Label: "Americas",
						Items: []basecoat.SelectItem{
							{Type: "item", Value: "America/New_York", Label: "(GMT-5) New York"},
							{Type: "item", Value: "America/Los_Angeles", Label: "(GMT-8) Los Angeles"},
							{Type: "item", Value: "America/Chicago", Label: "(GMT-6) Chicago"},
							{Type: "item", Value: "America/Toronto", Label: "(GMT-5) Toronto"},
							{Type: "item", Value: "America/Vancouver", Label: "(GMT-8) Vancouver"},
							{Type: "item", Value: "America/Sao_Paulo", Label: "(GMT-3) Sao Paulo"},
						},
					},
					{
						Type:  "group",
						ID:    "demo-combobox-timezones-group-1",
						Label: "Europe",
						Items: []basecoat.SelectItem{
							{Type: "item", Value: "Europe/London", Label: "(GMT+0) London"},
							{Type: "item", Value: "Europe/Paris", Label: "(GMT+1) Paris"},
							{Type: "item", Value: "Europe/Berlin", Label: "(GMT+1) Berlin"},
							{Type: "item", Value: "Europe/Rome", Label: "(GMT+1) Rome"},
							{Type: "item", Value: "Europe/Madrid", Label: "(GMT+1) Madrid"},
							{Type: "item", Value: "Europe/Amsterdam", Label: "(GMT+1) Amsterdam"},
						},
					},
					{
						Type:  "group",
						ID:    "demo-combobox-timezones-group-2",
						Label: "Asia/Pacific",
						Items: []basecoat.SelectItem{
							{Type: "item", Value: "Asia/Tokyo", Label: "(GMT+9) Tokyo"},
							{Type: "item", Value: "Asia/Shanghai", Label: "(GMT+8) Shanghai"},
							{Type: "item", Value: "Asia/Singapore", Label: "(GMT+8) Singapore"},
							{Type: "item", Value: "Asia/Hong_Kong", Label: "(GMT+8) Hong Kong"},
							{Type: "item", Value: "Asia/Bangkok", Label: "(GMT+7) Bangkok"},
							{Type: "item", Value: "Asia/Jakarta", Label: "(GMT+7) Jakarta"},
							{Type: "item", Value: "Asia/Seoul", Label: "(GMT+9) Seoul"},
							{Type: "item", Value: "Australia/Sydney", Label: "(GMT+10) Sydney"},
							{Type: "item", Value: "Australia/Melbourne", Label: "(GMT+10) Melbourne"},
						},
					},
				},
			}),
		),
	)
}

func _alertDialogSection() HTML {
	return Div(
		X.Class("p-4"),
		basecoat.Dialog(basecoat.DialogParams{
			ID:          "alert-dialog-demo",
			Title:       "Are you absolutely sure?",
			Description: "This action cannot be undone. This will permanently delete your account and remove your data from our servers.",
			Trigger:     "Open dialog",
			TriggerAttrs: []HTML{
				X.Class("btn-outline"),
			},
			Footer: Footer(
				Button(
					X.Type("button"),
					X.Class("btn-outline"),
					X.Attr("onclick", "document.getElementById('alert-dialog-demo').close()"),
					"Cancel",
				),
				Button(
					X.Type("button"),
					X.Class("btn-primary"),
					X.Attr("onclick", "document.getElementById('alert-dialog-demo').close()"),
					"Continue",
				),
			),
		}),
	)
}
