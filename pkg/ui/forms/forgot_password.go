package forms

import (
	"net/http"

	"github.com/mikestefanello/pagoda/pkg/form"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/ui"
	. "github.com/mikestefanello/pagoda/pkg/ui/components"
	. "github.com/namzug16/gotags"
)

type ForgotPassword struct {
	Email string `form:"email" validate:"required,email"`
	form.Submission
}

func (f *ForgotPassword) Render(r *ui.Request) HTML {
	return Div(
		X.Class("card w-full max-w-sm"),
		Header(
			If(len(r.Title) > 0, H2(r.Title)),
			P("Enter your email address and we'll email you a link that allows you to reset your password."),
		),
		Section(
			Form(
				X.Class("form grid gap-6"),
				X.Id("forgot-password"),
				X.Method(http.MethodPost),
				HxBoost(),
				X.Action(r.Path(routenames.ForgotPasswordSubmit)),
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
					X.Class("flex flex-col items-center gap-2"),
					Button(X.Class("btn w-full"), "Reset password"),
					A(X.Href(r.Path(routenames.Home)), X.Class("btn-outline btn-destructive w-full"), "Cancel"),
				),
				CSRF(r),
			),
		),
	)
}
