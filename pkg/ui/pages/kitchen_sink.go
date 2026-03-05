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
	section{name: "Dialog", href: "/components/dialog", content: _dialogSection()},
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
