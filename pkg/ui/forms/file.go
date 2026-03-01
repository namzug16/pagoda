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
		X.Id("files"),
		X.Method(http.MethodPost),
		X.Action(r.Path(routenames.FilesSubmit)),
		X.Attr("enctype", "multipart/form-data"),
		FileField(FileFieldParams{
			Name:  "file",
			Label: "Test file",
			Help:  "Pick a file to upload.",
		}),
		ControlGroup(
			FormButton(ColorPrimary, "Upload"),
		),
		CSRF(r),
	)
}
