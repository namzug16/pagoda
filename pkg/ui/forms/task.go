package forms

import (
	"fmt"
	"net/http"

	"github.com/mikestefanello/pagoda/pkg/form"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/ui"
	. "github.com/mikestefanello/pagoda/pkg/ui/components"
	. "github.com/namzug16/gotags"
)

type Task struct {
	Delay   int    `form:"delay" validate:"gte=0"`
	Message string `form:"message" validate:"required"`
	form.Submission
}

func (f *Task) Render(r *ui.Request) HTML {
	return Form(
		X.Class("form grid gap-6"),
		X.Id("task"),
		X.Method(http.MethodPost),
		X.Attr("hx-post", r.Path(routenames.TaskSubmit)),
		FlashMessages(r),
		Div(
			X.Class("field"),
			X.Role("group"),
			Label("Delay (in seconds)"),
			Input(
				X.Name("delay"),
				X.Type("number"),
				X.Value(fmt.Sprint(f.Delay)),
				If(FormFieldHasError(f, "Delay"), X.Attr("aria-invalid", "true")),
			),
			P(X.Class("text-sm opacity-70"), "How long to wait until the task is executed"),
			FormFieldErrors(f, "Delay"),
		),
		Div(
			X.Class("field"),
			X.Role("group"),
			Label("Message"),
			Textarea(
				X.Name("message"),
				X.Class("min-h-24"),
				If(FormFieldHasError(f, "Message"), X.Attr("aria-invalid", "true")),
				f.Message,
			),
			P(X.Class("text-sm opacity-70"), "The message the task will output to the log"),
			FormFieldErrors(f, "Message"),
		),
		Div(
			X.Class("flex flex-col items-center gap-2"),
			Button(X.Class("btn w-full"), "Add task to queue"),
		),
		CSRF(r),
	)
}
