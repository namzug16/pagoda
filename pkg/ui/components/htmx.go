package components

import (
	"fmt"

	"github.com/mikestefanello/pagoda/pkg/ui"
	. "github.com/namzug16/gotags"
)

func HtmxListeners(r *ui.Request) HTML {
	const htmxErr = `
		document.body.addEventListener('htmx:beforeSwap', function(evt) {
			if (evt.detail.xhr.status >= 400){
				evt.detail.shouldSwap = true;
				evt.detail.target = htmx.find("body");
			}
		});
	`

	const htmxCSRF = `
		document.body.addEventListener('htmx:configRequest', function(evt)  {
			if (evt.detail.verb !== "get") {
				evt.detail.parameters['csrf'] = '%s';
			}
		})
	`

	return Fragment(
		Script(Raw(htmxErr)),
		If(len(r.CSRF) > 0, Script(Raw(fmt.Sprintf(htmxCSRF, r.CSRF)))),
	)
}

func HxBoost() HTML {
	return X.Attr("hx-boost", "true")
}
