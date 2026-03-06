package kitchensink

import (
	. "github.com/namzug16/gotags"
)

func ToastSection() HTML {
	return Div(
		X.Class("p-4"),
		Div(
			X.Class("flex flex-wrap items-center gap-2"),
			Button(X.Class("btn-outline"), X.Attr("hx-trigger", "click"), X.Attr("hx-get", "/fragments/toast/success"), X.Attr("hx-target", "#toaster"), X.Attr("hx-swap", "beforeend"), "Success"),
			Button(X.Class("btn-outline"), X.Attr("hx-trigger", "click"), X.Attr("hx-get", "/fragments/toast/error"), X.Attr("hx-target", "#toaster"), X.Attr("hx-swap", "beforeend"), "Error"),
			Button(X.Class("btn-outline"), X.Attr("hx-trigger", "click"), X.Attr("hx-get", "/fragments/toast/info"), X.Attr("hx-target", "#toaster"), X.Attr("hx-swap", "beforeend"), "Info"),
			Button(X.Class("btn-outline"), X.Attr("hx-trigger", "click"), X.Attr("hx-get", "/fragments/toast/warning"), X.Attr("hx-target", "#toaster"), X.Attr("hx-swap", "beforeend"), "Warning"),
		),
	)
}
