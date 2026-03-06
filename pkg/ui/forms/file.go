package forms

import (
	"net/http"

	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/ui"
	. "github.com/mikestefanello/pagoda/pkg/ui/components"
	. "github.com/namzug16/gotags"
)

type File struct{}

func (f File) Render(r *ui.Request) HTML {
	return Form(
		X.Class("form grid gap-6"),
		X.Id("files"),
		X.Method(http.MethodPost),
		X.Action(r.Path(routenames.FilesSubmit)),
		X.Attr("enctype", "multipart/form-data"),
		Div(
			X.Class("field"),
			X.Role("group"),
			Label("Test file"),
			Input(
				X.Type("file"),
				X.Name("file"),
			),
			P(X.Class("text-sm opacity-70"), "Pick a file to upload."),
		),
		Div(
			X.Class("flex flex-col items-center gap-2"),
			Button(X.Class("btn w-full"), "Upload"),
		),
		CSRF(r),
	)
}
