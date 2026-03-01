package emails

import (
	"fmt"

	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/ui"
	. "github.com/namzug16/gotags"
)

func ConfirmEmailAddress(ctx echo.Context, username, token string) HTML {
	url := ui.NewRequest(ctx).
		Url(routenames.VerifyEmail, token)

	return Fragment(
		Strong(fmt.Sprintf("Hello %s,", username)),
		Br(),
		Text("Please click on the following link to confirm your email address:"),
		Br(),
		A(X.Href(url), Text(url)),
	)
}
