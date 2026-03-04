package pages

import (
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/ui"
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
		X.Class("p-4"),
	)
}
