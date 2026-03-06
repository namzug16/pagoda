package kitchensink

import (
	"github.com/mikestefanello/pagoda/pkg/ui/lucide"
	. "github.com/namzug16/gotags"
)

func AlertSection() HTML {
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
