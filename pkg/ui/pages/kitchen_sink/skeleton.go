package kitchensink

import (
	. "github.com/namzug16/gotags"
)

func SkeletonSection() HTML {
	return Div(
		X.Class("p-4"),
		Div(
			X.Class("flex flex-col gap-4"),
			Div(
				X.Class("flex items-center gap-4"),
				Div(X.Class("bg-accent animate-pulse size-10 shrink-0 rounded-full")),
				Div(
					X.Class("grid gap-2"),
					Div(X.Class("bg-accent animate-pulse rounded-md h-4 w-[150px]")),
					Div(X.Class("bg-accent animate-pulse rounded-md h-4 w-[100px]")),
				),
			),
			Div(
				X.Class("flex max-sm:flex-col gap-4 w-full"),
				Div(
					X.Class("card w-full @md:w-auto @md:min-w-sm"),
					Header(
						Div(X.Class("bg-accent animate-pulse rounded-md h-4 w-2/3")),
						Div(X.Class("bg-accent animate-pulse rounded-md h-4 w-1/2")),
					),
					Section(Div(X.Class("bg-accent animate-pulse rounded-md aspect-square w-full"))),
				),
				Div(
					X.Class("card w-full @md:w-auto @md:min-w-sm"),
					Header(
						Div(X.Class("bg-accent animate-pulse rounded-md h-4 w-2/3")),
						Div(X.Class("bg-accent animate-pulse rounded-md h-4 w-1/2")),
					),
					Section(Div(X.Class("bg-accent animate-pulse rounded-md aspect-square w-full"))),
				),
			),
		),
	)
}
