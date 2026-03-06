package forms

import (
	"net/http"

	"github.com/mikestefanello/pagoda/pkg/form"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/ui"
	. "github.com/mikestefanello/pagoda/pkg/ui/components"
	. "github.com/namzug16/gotags"
)

type Register struct {
	Name            string `form:"name" validate:"required"`
	Email           string `form:"email" validate:"required,email"`
	Password        string `form:"password" validate:"required"`
	ConfirmPassword string `form:"password-confirm" validate:"required,eqfield=Password"`
	form.Submission
}

func (f *Register) Render(r *ui.Request) HTML {
	return Div(
		X.Class("card w-full max-w-sm"),
		Header(
			If(len(r.Title) > 0, H2(r.Title)),
		),
		Section(
			Form(
				X.Class("form grid gap-6"),
				X.Id("register"),
				X.Method(http.MethodPost),
				X.Action(r.Path(routenames.RegisterSubmit)),
				HxBoost(),
				Div(
					X.Class("field"),
					X.Role("group"),
					Label("Name"),
					Input(
						X.Name("name"),
						X.Type("text"),
						If(FormFieldHasError(f, "Name"), X.Attr("aria-invalid", "true")),
						X.Value(f.Name),
					),
					FormFieldErrors(f, "Name"),
				),
				Div(
					X.Class("field"),
					X.Role("group"),
					Label("Email"),
					Input(
						X.Name("email"),
						X.Type("email"),
						If(FormFieldHasError(f, "Email"), X.Attr("aria-invalid", "true")),
						X.Value(f.Email),
					),
				),
				Div(
					X.Class("field"),
					X.Role("group"),
					Label("Password"),
					Input(
						X.Type("password"),
						X.Name("password"),
						X.Placeholder("******"),
						If(FormFieldHasError(f, "Email"), X.Attr("aria-invalid", "true")),
					),
				),
				Div(
					X.Class("field"),
					X.Role("group"),
					Label("Confirm Password"),
					Input(
						X.Type("password"),
						X.Name("password-confirm"),
						X.Placeholder("******"),
					),
				),
				Div(
					X.Class("flex flex-col items-center gap-2"),
					Button(X.Class("btn w-full"), "Register"),
					A(X.Href(r.Path(routenames.Home)), X.Class("btn-outline btn-error w-full"), "Cancel"),
				),
			),
		),
		CSRF(r),
		Footer(
			P(
				X.Class("mt-4 text-center text-sm"),
				"Already have an account? ",
				A(X.Href(r.Path(routenames.Login)), X.Class("underline-offset-4 hover:underline"), "Sign up"),
			),
		),
	)
}
