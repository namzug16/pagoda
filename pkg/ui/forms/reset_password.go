package forms

import (
	"net/http"

	"github.com/mikestefanello/pagoda/pkg/form"
	"github.com/mikestefanello/pagoda/pkg/ui"
	. "github.com/mikestefanello/pagoda/pkg/ui/components"
	. "github.com/namzug16/gotags"
)

type ResetPassword struct {
	Password        string `form:"password" validate:"required"`
	ConfirmPassword string `form:"password-confirm" validate:"required,eqfield=Password"`
	form.Submission
}

func (f *ResetPassword) Render(r *ui.Request) HTML {
	return Form(
		X.Id("reset-password"),
		X.Method(http.MethodPost),
		HxBoost(),
		X.Action(r.CurrentPath),
		InputField(InputFieldParams{
			Form:        f,
			FormField:   "Password",
			Name:        "password",
			InputType:   "password",
			Label:       "Password",
			Placeholder: "******",
		}),
		InputField(InputFieldParams{
			Form:        f,
			FormField:   "PasswordConfirm",
			Name:        "password-confirm",
			InputType:   "password",
			Label:       "Confirm password",
			Placeholder: "******",
		}),
		ControlGroup(
			FormButton(ColorPrimary, "Update password"),
		),
		CSRF(r),
	)
}
