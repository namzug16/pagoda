package models

import (
	"fmt"

	. "github.com/namzug16/gotags"
)

type File struct {
	Name     string
	Size     int64
	Modified string
}

func (f *File) Render() HTML {
	return Tr(
		Td(f.Name),
		Td(fmt.Sprint(f.Size)),
		Td(f.Modified),
	)
}
