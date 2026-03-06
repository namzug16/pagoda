package forms

import (
	"fmt"
	"net/http"
	"net/url"

	"entgo.io/ent/schema/field"
	"github.com/mikestefanello/pagoda/ent/admin"
	"github.com/mikestefanello/pagoda/pkg/routenames"
	"github.com/mikestefanello/pagoda/pkg/ui"
	. "github.com/mikestefanello/pagoda/pkg/ui/components"
	. "github.com/namzug16/gotags"
)

func AdminEntity(r *ui.Request, entityType admin.EntityType, values url.Values) HTML {
	// TODO:(pagoda) inline validation?
	isNew := values == nil
	nodes := make([]HTML, 0)

	getValue := func(name string) string {
		// Values in the submitted form take precedence.
		if value := r.Context.FormValue(name); value != "" {
			return value
		}

		// Fallback to the entity's values, if being edited.
		if values != nil && len(values[name]) > 0 {
			return values[name][0]
		}

		return ""
	}

	// Attempt to add form elements for all editable entity fields.
	for _, f := range entityType.GetSchema() {
		// TODO:(pagoda) cardinality?
		if !isNew && f.Immutable {
			continue
		}

		switch f.Type {
		case field.TypeString:
			inputType := "text"
			placeholder := ""
			help := ""

			if f.Sensitive {
				inputType = "password"
				if !isNew {
					placeholder = "*****"
					help = "SENSITIVE: This field will only be updated if a value is provided."
				}
			}

			nodes = append(nodes, Div(
				X.Class("field"),
				X.Role("group"),
				Label(admin.FieldLabel(f.Name)),
				Input(
					X.Name(f.Name),
					X.Type(inputType),
					X.Value(getValue(f.Name)),
					If(placeholder != "", X.Placeholder(placeholder)),
				),
				If(help != "", P(X.Class("text-sm opacity-70"), help)),
			))

		case field.TypeTime:
			nodes = append(nodes, Div(
				X.Class("field"),
				X.Role("group"),
				Label(admin.FieldLabel(f.Name)),
				Input(
					X.Name(f.Name),
					X.Type("datetime-local"),
					X.Value(getValue(f.Name)),
				),
			))

		case field.TypeInt, field.TypeInt8, field.TypeInt16, field.TypeInt32, field.TypeInt64,
			field.TypeUint, field.TypeUint8, field.TypeUint16, field.TypeUint32, field.TypeUint64,
			field.TypeFloat32, field.TypeFloat64:
			nodes = append(nodes, Div(
				X.Class("field"),
				X.Role("group"),
				Label(admin.FieldLabel(f.Name)),
				Input(
					X.Name(f.Name),
					X.Type("number"),
					X.Value(getValue(f.Name)),
				),
			))

		case field.TypeBool:
			nodes = append(nodes, Div(
				X.Class("field"),
				Label(
					X.Class("inline-flex items-center gap-2"),
					Input(
						X.Type("checkbox"),
						X.Name(f.Name),
						X.Value("true"),
						If(getValue(f.Name) == "true", X.Attr("checked")),
					),
					admin.FieldLabel(f.Name),
				),
			))

		case field.TypeEnum:
			options := make([]HTML, 0, len(f.Enums)+1)
			if f.Optional {
				options = append(options, Option(
					"-",
					X.Value(""),
					If(getValue(f.Name) == "", X.Attr("selected")),
				))
			}

			for _, enum := range f.Enums {
				options = append(options, Option(
					enum,
					X.Value(enum),
					If(getValue(f.Name) == enum, X.Attr("selected")),
				))
			}

			nodes = append(nodes, Div(
				X.Class("field"),
				X.Role("group"),
				Label(admin.FieldLabel(f.Name)),
				Select(
					X.Name(f.Name),
					options,
				),
			))

		default:
			nodes = append(nodes, P(fmt.Sprintf("%s not supported", f.Name)))
		}
	}

	return Form(
		X.Class("form grid gap-6"),
		X.Method(http.MethodPost),
		nodes,
		Div(
			X.Class("flex flex-col items-center gap-2"),
			Button(X.Class("btn w-full"), "Submit"),
			A(
				X.Href(r.Path(routenames.AdminEntityList(entityType.GetName()))),
				X.Class("btn-outline btn-destructive w-full"),
				"Cancel",
			),
		),
		CSRF(r),
	)
}
