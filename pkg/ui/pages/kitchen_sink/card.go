package kitchensink

import (
	. "github.com/namzug16/gotags"
)

func CardSection() HTML {
	return Div(
		X.Class("p-4"),
		Div(
			X.Class("flex flex-col items-start gap-4"),
			Div(
				X.Class("card w-full max-w-sm"),
				Header(
					H2("Login to your account"),
					P("Enter your details below to login to your account"),
				),
				Section(
					Form(
						X.Class("form grid gap-6"),
						Div(
							X.Class("grid gap-2"),
							Label(X.For("demo-card-form-email"), "Email"),
							Input(X.Type("email"), X.Id("demo-card-form-email")),
						),
						Div(
							X.Class("grid gap-2"),
							Div(
								X.Class("flex items-center gap-2"),
								Label(X.For("demo-card-form-password"), "Password"),
								A(X.Href("#"), X.Class("ml-auto inline-block text-sm underline-offset-4 hover:underline"), "Forgot your password?"),
							),
							Input(X.Type("password"), X.Id("demo-card-form-password")),
						),
					),
				),
				Footer(
					X.Class("flex flex-col items-center gap-2"),
					Button(X.Type("button"), X.Class("btn w-full"), "Login"),
					Button(X.Type("button"), X.Class("btn-outline w-full"), "Login with Google"),
					P(
						X.Class("mt-4 text-center text-sm"),
						"Don't have an account? ",
						A(X.Href("#"), X.Class("underline-offset-4 hover:underline"), "Sign up"),
					),
				),
			),
			Div(
				X.Class("card"),
				Header(
					H2("Meeting Notes"),
					P("Transcript from the meeting with the client."),
				),
				Section(
					X.Class("text-sm"),
					P("Client requested dashboard redesign with focus on mobile responsiveness."),
					Ol(
						X.Class("mt-4 flex list-decimal flex-col gap-2 pl-6"),
						Li("New analytics widgets for daily/weekly metrics"),
						Li("Simplified navigation menu"),
						Li("Dark mode support"),
						Li("Timeline: 6 weeks"),
						Li("Follow-up meeting scheduled for next Tuesday"),
					),
				),
				Footer(
					X.Class("flex items-center"),
					Div(
						X.Class("flex -space-x-2 [&_img]:ring-background [&_img]:ring-2 [&_img]:grayscale [&_img]:size-8 [&_img]:shrink-0 [&_img]:object-cover [&_img]:rounded-full"),
						Img(X.Alt("@hunvreus"), X.Src("https://github.com/hunvreus.png")),
						Img(X.Alt("@shadcn"), X.Src("https://github.com/shadcn.png")),
						Img(X.Alt("@adamwathan"), X.Src("https://github.com/adamwathan.png")),
					),
				),
			),
			Div(
				X.Class("card"),
				Header(
					H2("Is this an image?"),
					P("This is a card with an image."),
				),
				Section(
					X.Class("px-0"),
					Img(
						X.Alt("Photo by Drew Beamer"),
						X.Class("aspect-video object-cover"),
						X.Src("https://images.unsplash.com/photo-1588345921523-c2dcdb7f1dcd?w=1080&q=75"),
					),
				),
				Footer(
					X.Class("flex items-center gap-2"),
					Span(X.Class("badge-outline"), "1 bed"),
					Span(X.Class("badge-outline"), "2 bath"),
					Span(X.Class("badge-outline"), "350m2"),
					Span(X.Class("ml-auto font-medium tabular-nums"), "$135,000"),
				),
			),
			Div(
				X.Class("flex w-full flex-wrap items-start gap-8 md:*:[.card]:basis-1/4"),
				Div(X.Class("card"), Section("Content Only")),
				Div(
					X.Class("card"),
					Header(
						H2("Header Only"),
						P("This is a card with a header and a description."),
					),
				),
				Div(
					X.Class("card"),
					Header(
						H2("Header and Content"),
						P("This is a card with a header and a content."),
					),
					Section("Content only."),
				),
				Div(X.Class("card"), Footer("Footer Only")),
				Div(
					X.Class("card"),
					Header(
						H2("Header + Footer"),
						P("This is a card with a header and a footer."),
					),
					Footer("Footer"),
				),
				Div(X.Class("card"), Section("Content"), Footer("Footer")),
				Div(
					X.Class("card"),
					Header(
						H2("Header + Content + Footer"),
						P("This is a card with a header, content and footer."),
					),
					Section("Content"),
					Footer("Footer"),
				),
			),
		),
	)
}
