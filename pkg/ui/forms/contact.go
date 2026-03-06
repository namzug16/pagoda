package forms

import (
	"net/http"

	"github.com/mikestefanello/pagoda/pkg/form"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/ui"
	. "github.com/mikestefanello/pagoda/pkg/ui/components"
	. "github.com/namzug16/gotags"
)

type Contact struct {
	Email      string `form:"email" validate:"required,email"`
	Department string `form:"department" validate:"required,oneof=sales marketing hr"`
	Message    string `form:"message" validate:"required"`
	form.Submission
}

func (f *Contact) Render(r *ui.Request) HTML {
	return Form(
		X.Class("form grid gap-6"),
		X.Id("contact"),
		X.Method(http.MethodPost),
		X.Attr("hx-post", r.Path(routenames.ContactSubmit)),
		Div(
			X.Class("field"),
			X.Role("group"),
			Label("Email address"),
			Input(
				X.Name("email"),
				X.Type("email"),
				X.Value(f.Email),
				If(FormFieldHasError(f, "Email"), X.Attr("aria-invalid", "true")),
			),
			FormFieldErrors(f, "Email"),
		),
		Div(
			X.Class("field"),
			X.Role("group"),
			Label("Department"),
			Div(
				X.Class("grid gap-2"),
				Label(
					X.Class("inline-flex items-center gap-2"),
					Input(
						X.Type("radio"),
						X.Name("department"),
						X.Value("sales"),
						If(f.Department == "sales", X.Attr("checked")),
					),
					"Sales",
				),
				Label(
					X.Class("inline-flex items-center gap-2"),
					Input(
						X.Type("radio"),
						X.Name("department"),
						X.Value("marketing"),
						If(f.Department == "marketing", X.Attr("checked")),
					),
					"Marketing",
				),
				Label(
					X.Class("inline-flex items-center gap-2"),
					Input(
						X.Type("radio"),
						X.Name("department"),
						X.Value("hr"),
						If(f.Department == "hr", X.Attr("checked")),
					),
					"HR",
				),
			),
			FormFieldErrors(f, "Department"),
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
			FormFieldErrors(f, "Message"),
		),
		Div(
			X.Class("flex flex-col items-center gap-2"),
			Button(X.Class("btn w-full"), "Submit"),
		),
		CSRF(r),
	)
}
