package forms

import (
	"net/http"

	"github.com/mikestefanello/pagoda/pkg/form"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/ui"
	. "github.com/mikestefanello/pagoda/pkg/ui/components"
	. "github.com/namzug16/gotags"
)

type Login struct {
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required"`
	form.Submission
}

func (f *Login) Render(r *ui.Request) HTML {
	return Div(
		X.Class("card w-full max-w-sm"),
		Header(
			If(len(r.Title) > 0, H2(r.Title)),
		),
		Section(
			Form(
				X.Class("form grid gap-6"),
				X.Id("login"),
				X.Method(http.MethodPost),
				HxBoost(),
				X.Action(r.Path(routenames.LoginSubmit)),
				FlashMessages(r),
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
					Label("Password"),
					Input(
						X.Name("password"),
						X.Type("password"),
						X.Placeholder("******"),
						If(FormFieldHasError(f, "Password"), X.Attr("aria-invalid", "true")),
					),
					FormFieldErrors(f, "Password"),
				),
				Div(
					X.Class("text-right text-primary mt-2"),
					A(
						X.Href(r.Path(routenames.ForgotPassword)),
						"Forgot password?",
					),
				),
				Div(
					X.Class("flex flex-col items-center gap-2"),
					Button(X.Class("btn w-full"), "Login"),
					A(X.Href(r.Path(routenames.Home)), X.Class("btn-outline btn-destructive w-full"), "Cancel"),
				),
				CSRF(r),
			),
		),

		Footer(
			Div(
				X.Class("mt-4 text-center text-sm"),
				"Don't have an account? ",
				A(
					X.Href(r.Path(routenames.Register)),
					X.Class("underline-offset-4 hover:underline"),
					"Register",
				),
			),
		),
	)
}
