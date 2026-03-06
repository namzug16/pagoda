package kitchensink

import (
	. "github.com/namzug16/gotags"
)

func AvatarSection() HTML {
	return Div(
		X.Class("p-4"),
		Div(
			X.Class("flex flex-row flex-wrap items-center gap-4"),
			Img(
				X.Class("size-8 shrink-0 object-cover rounded-full"),
				X.Alt("@hunvreus"),
				X.Src("https://github.com/hunvreus.png"),
			),
			Span(
				X.Class("size-8 shrink-0 bg-muted flex items-center justify-center rounded-full"),
				"CN",
			),
			Img(
				X.Class("size-12 shrink-0 object-cover rounded-full"),
				X.Alt("@hunvreus"),
				X.Src("https://github.com/hunvreus.png"),
			),
			Img(
				X.Class("size-8 shrink-0 object-cover rounded-lg"),
				X.Alt("@hunvreus"),
				X.Src("https://github.com/hunvreus.png"),
			),
			Div(
				X.Class("flex -space-x-2 [&_img]:ring-background [&_img]:ring-2 [&_img]:grayscale [&_img]:size-8 [&_img]:shrink-0 [&_img]:object-cover [&_img]:rounded-full"),
				Img(X.Alt("@hunvreus"), X.Src("https://github.com/hunvreus.png")),
				Img(X.Alt("@shadcn"), X.Src("https://github.com/shadcn.png")),
				Img(X.Alt("@adamwathan"), X.Src("https://github.com/adamwathan.png")),
			),
			Div(
				X.Class("flex -space-x-2 [&_img]:ring-background [&_img]:ring-2 [&_img]:grayscale [&_img]:size-12 [&_img]:shrink-0 [&_img]:object-cover [&_img]:rounded-full"),
				Img(X.Alt("@hunvreus"), X.Src("https://github.com/hunvreus.png")),
				Img(X.Alt("@shadcn"), X.Src("https://github.com/shadcn.png")),
				Img(X.Alt("@adamwathan"), X.Src("https://github.com/adamwathan.png")),
			),
			Div(
				X.Class("flex -space-x-2 hover:space-x-1 [&_img]:ring-background [&_img]:size-12 [&_img]:shrink-0 [&_img]:object-cover [&_img]:rounded-full [&_img]:ring-2 [&_img]:grayscale [&_img]:transition-all [&_img]:ease-in-out [&_img]:duration-300"),
				Img(X.Alt("@hunvreus"), X.Src("https://github.com/hunvreus.png")),
				Img(X.Alt("@shadcn"), X.Src("https://github.com/shadcn.png")),
				Img(X.Alt("@adamwathan"), X.Src("https://github.com/adamwathan.png")),
			),
		),
	)
}
