package pages

import (
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/ui"
	"github.com/mikestefanello/pagoda/pkg/ui/forms"
	"github.com/mikestefanello/pagoda/pkg/ui/layouts"
	. "github.com/namzug16/gotags"
)

func Login(ctx echo.Context, form *forms.Login) error {
	r := ui.NewRequest(ctx)
	r.Title = "Login"

	return r.Render(layouts.Auth, form.Render(r))
}

func Register(ctx echo.Context, form *forms.Register) error {
	r := ui.NewRequest(ctx)
	r.Title = "Register"

	return r.Render(layouts.Auth, form.Render(r))
}

func ForgotPassword(ctx echo.Context, form *forms.ForgotPassword) error {
	r := ui.NewRequest(ctx)
	r.Title = "Forgot password"

	g := Fragment(
		Div(
			X.Class("content"),
			"Enter your email address and we'll email you a link that allows you to reset your password.",
		),
		form.Render(r),
	)

	return r.Render(layouts.Auth, g)
}

func ResetPassword(ctx echo.Context, form *forms.ResetPassword) error {
	r := ui.NewRequest(ctx)
	r.Title = "Reset your password"

	return r.Render(layouts.Auth, form.Render(r))
}
