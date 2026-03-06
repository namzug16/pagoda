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
		X.Class("form grid gap-6"),
		X.Id("cache"),
		X.Method(http.MethodPost),
		X.Attr("hx-post", r.Path(routenames.CacheSubmit)),
		Div(
			X.Class("grid gap-2 rounded-lg border border-base-300 p-4"),
			H3(X.Class("text-lg font-semibold"), "Test the cache"),
			P("This route handler shows how the default in-memory cache works. Try updating the value using the form below and see how it persists after you reload the page."),
			P("HTMX makes it easy to re-render the cached value after the form is submitted."),
		),
		Div(
			X.Class("field"),
			X.Role("group"),
			Label("Value in cache"),
			If(f.CurrentValue != "", Span(X.Class("badge badge-success"), f.CurrentValue)),
			If(f.CurrentValue == "", Span(X.Class("badge badge-warning"), "empty")),
		),
		Div(
			X.Class("field"),
			X.Role("group"),
			Label("Value"),
			Input(
				X.Name("value"),
				X.Type("text"),
				X.Value(f.Value),
				If(FormFieldHasError(f, "Value"), X.Attr("aria-invalid", "true")),
			),
			FormFieldErrors(f, "Value"),
		),
		Div(
			X.Class("flex flex-col items-center gap-2"),
			Button(X.Class("btn w-full"), "Update cache"),
		),
		CSRF(r),
	)
}
