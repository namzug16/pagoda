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
	return Form(
		X.Id("login"),
		X.Method(http.MethodPost),
		HxBoost(),
		X.Action(r.Path(routenames.LoginSubmit)),
		FlashMessages(r),
		InputField(InputFieldParams{
			Form:      f,
			FormField: "Email",
			Name:      "email",
			InputType: "email",
			Label:     "Email address",
			Value:     f.Email,
		}),
		InputField(InputFieldParams{
			Form:        f,
			FormField:   "Password",
			Name:        "password",
			InputType:   "password",
			Label:       "Password",
			Placeholder: "******",
		}),
		Div(
			X.Class("text-right text-primary mt-2"),
			A(
				X.Href(r.Path(routenames.ForgotPassword)),
				"Forgot password?",
			),
		),
		ControlGroup(
			FormButton(ColorPrimary, "Login"),
			ButtonLink(ColorLink, r.Path(routenames.Home), "Cancel"),
		),
		CSRF(r),
		Div(
			X.Class("text-center text-base-content/50 mt-4"),
			"Don't have an account? ",
			A(
				X.Href(r.Path(routenames.Register)),
				"Register",
			),
		),
	)
}
