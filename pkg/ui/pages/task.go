package pages

import (
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/ui"
	. "github.com/mikestefanello/pagoda/pkg/ui/components"
	"github.com/mikestefanello/pagoda/pkg/ui/forms"
	"github.com/mikestefanello/pagoda/pkg/ui/layouts"
	. "github.com/namzug16/gotags"
)

func AddTask(ctx echo.Context, form *forms.Task) error {
	r := ui.NewRequest(ctx)
	r.Title = "Create a task"
	r.Metatags.Description = "Test creating a task to see how it works."

	g := Fragment(
		If(r.Htmx.Target != "task",
			Fragment(
				Raw("Submitting this form will create an <i>ExampleTask</i> in the task queue. After the specified delay, the message will be logged by the queue processor."),
				Raw("See <i>pkg/tasks</i> and the README for more information."),
			),
		),
		form.Render(r),
		If(r.Htmx.Target != "task",
			func() HTML {
				var text string
				if r.IsAdmin {
					text = "View all queued tasks by clicking on the Tasks link in the sidebar."
				} else {
					text = "Log in as an admin in order to access the task and queue monitoring UI."
				}
				return Fragment(
					Div(X.Class("mt-5")),
					Alert(ColorWarning, text),
				)
			}(),
		),
	)

	return r.Render(layouts.Primary, g)
}
