package components

import (
	"github.com/mikestefanello/pagoda/pkg/form"
	"github.com/mikestefanello/pagoda/pkg/ui"
	. "github.com/namzug16/gotags"
)

type (
	InputFieldParams struct {
		Form        form.Form
		FormField   string
		Name        string
		InputType   string
		Label       string
		Value       string
		Placeholder string
		Help        string
	}

	FileFieldParams struct {
		Name  string
		Label string
		Help  string
	}

	OptionsParams struct {
		Form      form.Form
		FormField string
		Name      string
		Label     string
		Value     string
		Options   []Choice
		Help      string
	}

	Choice struct {
		Value string
		Label string
	}

	TextareaFieldParams struct {
		Form      form.Form
		FormField string
		Name      string
		Label     string
		Value     string
		Help      string
	}

	CheckboxParams struct {
		Form      form.Form
		FormField string
		Name      string
		Label     string
		Checked   bool
	}
)

func ControlGroup(controls ...HTML) HTML {
	return Div(
		X.Class("mt-2 flex gap-2"),
		Fragment(controls),
	)
}

func TextareaField(el TextareaFieldParams) HTML {
	return FieldSetComponent(
		el.Label,
		Textarea(
			X.Class("textarea h-24 w-2/3 "+FormFieldStatusClass(el.Form, el.FormField)),
			X.Id(el.Name),
			X.Name(el.Name),
			el.Value,
		),
		HelpText(el.Help),
		FormFieldErrors(el.Form, el.FormField),
	)
}

func Radios(el OptionsParams) HTML {
	buttons := make([]HTML, len(el.Options))
	for i, opt := range el.Options {
		id := "radio-" + el.Name + "-" + opt.Value
		buttons[i] = Div(
			X.Class("mb-2"),
			Input(
				X.Id(id),
				X.Type("radio"),
				X.Name(el.Name),
				X.Value(opt.Value),
				X.Class("radio mr-1 "+FormFieldStatusClass(el.Form, el.FormField)),
				If(el.Value == opt.Value, X.Checked()),
			),
			Label(
				opt.Label,
				X.For(id),
			),
		)
	}

	return FieldSetComponent(
		el.Label,
		Fragment(buttons),
		FormFieldErrors(el.Form, el.FormField),
	)
}

func SelectList(el OptionsParams) HTML {
	buttons := make([]HTML, len(el.Options))
	for i, opt := range el.Options {
		buttons[i] = Option(
			opt.Label,
			X.Value(opt.Value),
			If(opt.Value == el.Value, X.Attr("selected")),
		)
	}

	return FieldSetComponent(
		el.Label,
		Select(
			X.Class("select "+FormFieldStatusClass(el.Form, el.FormField)),
			X.Name(el.Name),
			buttons,
		),
		HelpText(el.Help),
		FormFieldErrors(el.Form, el.FormField),
	)
}

func Checkbox(el CheckboxParams) HTML {
	return Div(
		Label(
			X.Class("label"),
			Input(
				X.Class("checkbox"),
				X.Type("checkbox"),
				X.Name(el.Name),
				If(el.Checked, X.Checked()),
				X.Value("true"),
			),
			" "+el.Label,
		),
		FormFieldErrors(el.Form, el.FormField),
	)
}

func InputField(el InputFieldParams) HTML {
	return FieldSetComponent(
		el.Label,
		Input(
			X.Id(el.Name),
			X.Name(el.Name),
			X.Type(el.InputType),
			X.Class("input "+FormFieldStatusClass(el.Form, el.FormField)),
			X.Value(el.Value),
			If(el.Placeholder != "", X.Placeholder(el.Placeholder)),
		),
		HelpText(el.Help),
		FormFieldErrors(el.Form, el.FormField),
	)
}

func HelpText(text string) HTML {
	return Fragment()
	// return If(len(text) > 0, Div(
	// 	X.Class("label"),
	// 	text,
	// ))
}

func FieldSetComponent(label string, els ...HTML) HTML {
	return Fieldset(
		X.Class("fieldset"),
		If(len(label) > 0, Legend(
			X.Class("fieldset-legend"),
			label,
		)),
		Fragment(els),
	)
}

func FileField(el FileFieldParams) HTML {
	return FieldSetComponent(
		el.Label,
		Input(
			X.Type("file"),
			X.Class("file-input"),
			X.Name(el.Name),
		),
		HelpText(el.Help),
	)
}

func FormFieldStatusClass(fm form.Form, formField string) string {
	switch {
	case fm == nil:
		return ""
	case !fm.IsSubmitted():
		return ""
	case fm.FieldHasErrors(formField):
		return "input-error"
	default:
		return "input-success"
	}
}

func FormFieldHasError(fm form.Form, field string) bool {
	if fm == nil {
		return false
	}

	errs := fm.GetFieldErrors(field)
	if len(errs) == 0 {
		return false
	}

	return true
}

func FormFieldErrors(fm form.Form, field string) HTML {
	if fm == nil {
		return Fragment()
	}

	errs := fm.GetFieldErrors(field)
	if len(errs) == 0 {
		return Fragment()
	}

	g := make([]HTML, len(errs))
	for i, err := range errs {
		g[i] = P(
			X.Role("alert"),
			err,
		)
	}

	return Fragment(g)
}

func CSRF(r *ui.Request) HTML {
	return Input(
		X.Type("hidden"),
		X.Name("csrf"),
		X.Value(r.CSRF),
	)
}

func FormButton(color Color, label string) HTML {
	return Button(
		X.Class("btn "+buttonColor(color)),
		label,
	)
}

func ButtonLink(color Color, href, label string) HTML {
	return A(
		X.Href(href),
		X.Class("btn "+buttonColor(color)),
		label,
	)
}

func buttonColor(color Color) string {
	// Only colors being used are included so unused styles are not compiled.
	switch color {
	case ColorPrimary:
		return "btn-primary"
	case ColorInfo:
		return "btn-info"
	case ColorAccent:
		return "btn-accent"
	case ColorError:
		return "btn-error"
	case ColorLink:
		return "btn-link"
	default:
		return ""
	}
}
