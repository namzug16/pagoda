package forms

import (
	"net/http"

	"github.com/mikestefanello/pagoda/pkg/form"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/ui"
	. "github.com/mikestefanello/pagoda/pkg/ui/components"
	. "github.com/namzug16/gotags"
)

type Cache struct {
	CurrentValue string
	Value        string `form:"value"`
	form.Submission
}

func (f *Cache) Render(r *ui.Request) HTML {
	return Form(
		X.Id("cache"),
		X.Method(http.MethodPost),
		X.Attr("hx-post", r.Path(routenames.CacheSubmit)),
		Card(CardParams{
			Title: "Test the cache",
			Body: []HTML{
				Text("This route handler shows how the default in-memory cache works. Try updating the value using the form below and see how it persists after you reload the page."),
				Text("HTMX makes it easy to re-render the cached value after the form is submitted."),
			},
			Color: ColorInfo,
			Size:  SizeMedium,
		}),
		Label(
			X.For("value"),
			X.Class("value"),
			"Value in cache: ",
		),
		If(f.CurrentValue != "", Badge(ColorSuccess, f.CurrentValue)),
		If(f.CurrentValue == "", Badge(ColorWarning, "empty")),
		InputField(InputFieldParams{
			Form:      f,
			FormField: "Value",
			Name:      "value",
			InputType: "text",
			Label:     "Value",
			Value:     f.Value,
		}),
		ControlGroup(
			FormButton(ColorPrimary, "Update cache"),
		),
		CSRF(r),
	)
}
