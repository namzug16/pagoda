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
		X.Class("form grid gap-6"),
		X.Id("reset-password"),
		X.Method(http.MethodPost),
		HxBoost(),
		X.Action(r.CurrentPath),
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
			X.Class("field"),
			X.Role("group"),
			Label("Confirm password"),
			Input(
				X.Name("password-confirm"),
				X.Type("password"),
				X.Placeholder("******"),
				If(FormFieldHasError(f, "ConfirmPassword"), X.Attr("aria-invalid", "true")),
			),
			FormFieldErrors(f, "ConfirmPassword"),
		),
		Div(
			X.Class("flex flex-col items-center gap-2"),
			Button(X.Class("btn w-full"), "Update password"),
		),
		CSRF(r),
	)
}
