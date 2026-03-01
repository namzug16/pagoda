package pages

import (
	"github.com/labstack/echo/v4"
	"github.com/mikestefanello/pagoda/pkg/ui"
	. "github.com/mikestefanello/pagoda/pkg/ui/components"
	"github.com/mikestefanello/pagoda/pkg/ui/forms"
	"github.com/mikestefanello/pagoda/pkg/ui/layouts"
	"github.com/mikestefanello/pagoda/pkg/ui/models"
	. "github.com/namzug16/gotags"
)

func UploadFile(ctx echo.Context, files []*models.File) error {
	r := ui.NewRequest(ctx)
	r.Title = "Upload a file"

	fileList := make([]HTML, len(files))
	for i, file := range files {
		fileList[i] = file.Render()
	}

	n := Fragment(
		Text("This is a very basic example of how to handle file uploads. Files uploaded will be saved to the directory specified in your configuration."),
		Divider(""),
		forms.File{}.Render(r),
		Divider(""),
		H3(
			X.Class("title"),
			"Uploaded files",
		),
		Card(CardParams{
			Body:  []HTML{Text("Below are all files in the configured upload directory.")},
			Color: ColorWarning,
			Size:  SizeMedium,
		}),
		Table(
			X.Class("table"),
			Thead(
				Tr(
					Th("Filename"),
					Th("Size"),
					Th("Modified on"),
				),
			),
			Tbody(fileList),
		),
	)

	return r.Render(layouts.Primary, n)
}
