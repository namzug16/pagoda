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
	return Form(
		X.Id("forgot-password"),
		X.Method(http.MethodPost),
		HxBoost(),
		X.Action(r.Path(routenames.ForgotPasswordSubmit)),
		InputField(InputFieldParams{
			Form:      f,
			FormField: "Email",
			Name:      "email",
			InputType: "email",
			Label:     "Email address",
			Value:     f.Email,
		}),
		ControlGroup(
			FormButton(ColorPrimary, "Reset password"),
			ButtonLink(ColorLink, r.Path(routenames.Home), "Cancel"),
		),
		CSRF(r),
	)
}
