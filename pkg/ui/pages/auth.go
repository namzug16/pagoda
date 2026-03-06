package pages

import (
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/ui"
	"github.com/mikestefanello/pagoda/pkg/ui/forms"
	"github.com/mikestefanello/pagoda/pkg/ui/layouts"
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

	return r.Render(layouts.Auth, form.Render(r))
}

func ResetPassword(ctx echo.Context, form *forms.ResetPassword) error {
	r := ui.NewRequest(ctx)
	r.Title = "Reset your password"

	return r.Render(layouts.Auth, form.Render(r))
}
