package kitchensink

import (
	. "github.com/namzug16/gotags"
)

func TableSection() HTML {
	return Div(
		X.Class("p-4"),
		Div(
			X.Class("relative w-full overflow-x-auto"),
			Table(
				X.Class("table"),
				Caption("A list of your recent invoices."),
				Thead(
					Tr(Th("Invoice"), Th("Status"), Th("Method"), Th("Amount")),
				),
				Tbody(
					Tr(Td(X.Class("font-medium"), "INV001"), Td("Paid"), Td("Credit Card"), Td(X.Class("text-right"), "$250.00")),
					Tr(Td(X.Class("font-medium"), "INV002"), Td("Pending"), Td("PayPal"), Td(X.Class("text-right"), "$150.00")),
					Tr(Td(X.Class("font-medium"), "INV003"), Td("Unpaid"), Td("Bank Transfer"), Td(X.Class("text-right"), "$350.00")),
					Tr(Td(X.Class("font-medium"), "INV004"), Td("Paid"), Td("Paypal"), Td(X.Class("text-right"), "$450.00")),
					Tr(Td(X.Class("font-medium"), "INV005"), Td("Paid"), Td("Credit Card"), Td(X.Class("text-right"), "$550.00")),
					Tr(Td(X.Class("font-medium"), "INV006"), Td("Pending"), Td("Bank Transfer"), Td(X.Class("text-right"), "$200.00")),
					Tr(Td(X.Class("font-medium"), "INV007"), Td("Unpaid"), Td("Credit Card"), Td(X.Class("text-right"), "$300.00")),
				),
				Tfoot(
					Tr(
						Td(X.Attr("colspan", "3"), "Total"),
						Td(X.Class("text-right"), "$2,500.00"),
					),
				),
			),
		),
	)
}
