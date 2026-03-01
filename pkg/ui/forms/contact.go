package forms

import (
	"fmt"
	"net/http"

	"github.com/mikestefanello/pagoda/pkg/form"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/ui"
	. "github.com/mikestefanello/pagoda/pkg/ui/components"
	. "github.com/namzug16/gotags"
)

type Contact struct {
	Email      string `form:"email" validate:"required,email"`
	Department string `form:"department" validate:"required,oneof=sales marketing hr"`
	Message    string `form:"message" validate:"required"`
	form.Submission
}

func (f *Contact) Render(r *ui.Request) HTML {
	fmt.Println("ALKSJDLKJASLKDJKLASJD")
	fmt.Println(f.Email)
	return Form(
		X.Id("contact"),
		X.Method(http.MethodPost),
		X.Attr("hx-post", r.Path(routenames.ContactSubmit)),
		InputField(InputFieldParams{
			Form:      f,
			FormField: "Email",
			Name:      "email",
			InputType: "email",
			Label:     "Email address",
			Value:     f.Email,
		}),
		Radios(OptionsParams{
			Form:      f,
			FormField: "Department",
			Name:      "department",
			Label:     "Department",
			Value:     f.Department,
			Options: []Choice{
				{Value: "sales", Label: "Sales"},
				{Value: "marketing", Label: "Marketing"},
				{Value: "hr", Label: "HR"},
			},
		}),
		TextareaField(TextareaFieldParams{
			Form:      f,
			FormField: "Message",
			Name:      "message",
			Label:     "Message",
			Value:     f.Message,
		}),
		ControlGroup(
			FormButton(ColorPrimary, "Submit"),
		),
		CSRF(r),
	)
}
