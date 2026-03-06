package basecoat

import (
	"fmt"

	gt "github.com/namzug16/gotags"
)

func normalizeAnyComponent(c any) gt.HTML {
	switch cc := c.(type) {
	case nil:
		return gt.Fragment()
	case gt.HTML:
		return cc
	case []gt.HTML:
		return gt.Fragment(cc)
	case string:
		return gt.Text(cc)
	default:
		panic(fmt.Errorf(
			"Invalid component type %T; expected string, HTML or []HTML",
			c,
		))
	}
}
