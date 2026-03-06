package kitchensink

import (
	"github.com/mikestefanello/pagoda/pkg/ui/basecoat"
	"github.com/mikestefanello/pagoda/pkg/ui/lucide"
	. "github.com/namzug16/gotags"
)

func DropdownMenuSection() HTML {
	return Div(
		X.Class("p-4"),
		Div(
			X.Class("flex flex-wrap items-start gap-4"),
			basecoat.DropdownMenu(basecoat.DropdownMenuParams{
				ID:                  "dropdown-menu-default",
				Trigger:             "Open",
				TriggerExtraClasses: "btn-outline",
				PopoverExtraClasses: "min-w-56",
				Content: Fragment(
					Div(
						X.Attr("role", "group"),
						X.Attr("aria-labelledby", "demo-account-options"),
						Div(X.Attr("role", "heading"), X.Id("demo-account-options"), "My Account"),
						Div(X.Attr("role", "menuitem"), "Profile", Span(X.Class("text-muted-foreground ml-auto text-xs tracking-widest"), "Shift+Cmd+P")),
						Div(X.Attr("role", "menuitem"), "Billing", Span(X.Class("text-muted-foreground ml-auto text-xs tracking-widest"), "Cmd+B")),
						Div(X.Attr("role", "menuitem"), "Settings", Span(X.Class("text-muted-foreground ml-auto text-xs tracking-widest"), "Cmd+S")),
						Div(X.Attr("role", "menuitem"), "Keyboard shortcuts", Span(X.Class("text-muted-foreground ml-auto text-xs tracking-widest"), "Cmd+K")),
					),
					Hr(X.Attr("role", "separator")),
					Div(X.Attr("role", "menuitem"), "GitHub"),
					Div(X.Attr("role", "menuitem"), "Support"),
					Div(X.Attr("role", "menuitem"), X.Attr("aria-disabled", "true"), "API"),
					Hr(X.Attr("role", "separator")),
					Div(X.Attr("role", "menuitem"), "Logout", Span(X.Class("text-muted-foreground ml-auto text-xs tracking-widest"), "Shift+Cmd+Q")),
				),
			}),
			Div(
				X.Id("dropdown-menu-checkboxes"),
				X.Class("dropdown-menu"),
				Button(
					X.Type("button"),
					X.Id("dropdown-menu-checkboxes-trigger"),
					X.Attr("aria-haspopup", "menu"),
					X.Attr("aria-controls", "dropdown-menu-checkboxes-menu"),
					X.Attr("aria-expanded", "false"),
					X.Class("btn-outline"),
					"Checkboxes",
				),
				Div(
					X.Id("dropdown-menu-checkboxes-popover"),
					X.Attr("data-popover", ""),
					X.Attr("aria-hidden", "true"),
					X.Class("min-w-56"),
					Div(
						X.Attr("role", "menu"),
						X.Id("dropdown-menu-checkboxes-menu"),
						X.Attr("aria-labelledby", "dropdown-menu-checkboxes-trigger"),
						Div(
							X.Attr("role", "group"),
							X.Attr("aria-labelledby", "demo-appearance-options"),
							Div(X.Attr("role", "heading"), X.Id("demo-appearance-options"), "Appearance"),
							Div(
								X.Attr("role", "menuitemcheckbox"),
								X.Attr("aria-checked", "true"),
								X.Class("group"),
								lucide.Check(X.Class("invisible group-aria-checked:visible"), X.Attr("aria-hidden", "true"), X.Attr("focusable", "false")),
								"Status Bar",
								Span(X.Class("text-muted-foreground ml-auto text-xs tracking-widest"), "Shift+Cmd+P"),
							),
							Div(
								X.Attr("role", "menuitemcheckbox"),
								X.Attr("aria-checked", "false"),
								X.Attr("aria-disabled", "true"),
								X.Class("group"),
								lucide.Check(X.Class("invisible group-aria-checked:visible"), X.Attr("aria-hidden", "true"), X.Attr("focusable", "false")),
								"Activity Bar",
								Span(X.Class("text-muted-foreground ml-auto text-xs tracking-widest"), "Cmd+B"),
							),
							Div(
								X.Attr("role", "menuitemcheckbox"),
								X.Attr("aria-checked", "false"),
								X.Class("group"),
								lucide.Check(X.Class("invisible group-aria-checked:visible"), X.Attr("aria-hidden", "true"), X.Attr("focusable", "false")),
								"Panel",
								Span(X.Class("text-muted-foreground ml-auto text-xs tracking-widest"), "Cmd+S"),
							),
						),
					),
				),
			),
			Div(
				X.Id("dropdown-menu-radio-group"),
				X.Class("dropdown-menu"),
				Button(
					X.Type("button"),
					X.Id("dropdown-menu-radio-group-trigger"),
					X.Attr("aria-haspopup", "menu"),
					X.Attr("aria-controls", "dropdown-menu-radio-group-menu"),
					X.Attr("aria-expanded", "false"),
					X.Class("btn-outline"),
					"Radio Group",
				),
				Div(
					X.Id("dropdown-menu-radio-group-popover"),
					X.Attr("data-popover", ""),
					X.Attr("aria-hidden", "true"),
					X.Class("min-w-56"),
					Div(
						X.Attr("role", "menu"),
						X.Id("dropdown-menu-radio-group-menu"),
						X.Attr("aria-labelledby", "dropdown-menu-radio-group-trigger"),
						Div(
							X.Attr("role", "group"),
							X.Attr("aria-labelledby", "demo-position-options"),
							Span(X.Attr("role", "heading"), X.Id("demo-position-options"), "Panel Position"),
							Hr(X.Attr("role", "separator")),
							Div(
								X.Attr("role", "menuitemradio"),
								X.Attr("aria-checked", "false"),
								X.Class("group"),
								Div(X.Class("size-4 flex items-center justify-center"), Div(X.Class("size-2 rounded-full bg-foreground invisible group-aria-checked:visible"), X.Attr("aria-hidden", "true"), X.Attr("focusable", "false"))),
								"Status Bar",
								Span(X.Class("text-muted-foreground ml-auto text-xs tracking-widest"), "Shift+Cmd+P"),
							),
							Div(
								X.Attr("role", "menuitemradio"),
								X.Attr("aria-checked", "true"),
								X.Class("group"),
								Div(X.Class("size-4 flex items-center justify-center"), Div(X.Class("size-2 rounded-full bg-foreground invisible group-aria-checked:visible"), X.Attr("aria-hidden", "true"), X.Attr("focusable", "false"))),
								"Activity Bar",
								Span(X.Class("text-muted-foreground ml-auto text-xs tracking-widest"), "Cmd+B"),
							),
							Div(
								X.Attr("role", "menuitemradio"),
								X.Attr("aria-checked", "false"),
								X.Class("group"),
								Div(X.Class("size-4 flex items-center justify-center"), Div(X.Class("size-2 rounded-full bg-foreground invisible group-aria-checked:visible"), X.Attr("aria-hidden", "true"), X.Attr("focusable", "false"))),
								"Panel",
								Span(X.Class("text-muted-foreground ml-auto text-xs tracking-widest"), "Cmd+S"),
							),
						),
					),
				),
			),
			basecoat.DropdownMenu(basecoat.DropdownMenuParams{
				ID: "dropdown-menu-profile",
				Trigger: Fragment(
					Img(
						X.Class("size-8 shrink-0 rounded-full"),
						X.Alt("@hunvreus"),
						X.Src("https://github.com/hunvreus.png"),
					),
					Div(
						X.Class("grid flex-1 text-left text-sm leading-tight"),
						Span(X.Class("truncate font-medium"), "hunvreus"),
						Span(X.Class("text-muted-foreground truncate text-xs"), "hunvreus@example.com"),
					),
					lucide.ChevronDown(X.Class("text-muted-foreground ml-auto")),
				),
				TriggerExtraClasses: "btn-outline h-12 justify-start px-2 md:max-w-[200px]",
				Content: Fragment(
					Div(
						X.Class("flex items-center gap-2 px-1 py-1.5 text-left text-sm"),
						Img(X.Alt("@hunvreus"), X.Src("https://github.com/hunvreus.png"), X.Class("size-8 shrink-0 rounded-full")),
						Div(
							X.Class("grid flex-1 text-left text-sm leading-tight"),
							Span(X.Class("truncate font-medium"), "hunvreus"),
							Span(X.Class("text-muted-foreground truncate text-xs"), "hunvreus@example.com"),
						),
					),
					Hr(X.Attr("role", "separator")),
					Div(X.Attr("role", "menuitem"), lucide.CircleCheck(), "Upgrade to Pro"),
					Hr(X.Attr("role", "separator")),
					Div(X.Attr("role", "menuitem"), "Account"),
					Div(X.Attr("role", "menuitem"), "Billing"),
					Div(X.Attr("role", "menuitem"), "Notifications"),
					Hr(X.Attr("role", "separator")),
					Div(X.Attr("role", "menuitem"), "Signout"),
				),
			}),
			basecoat.DropdownMenu(basecoat.DropdownMenuParams{
				ID:                  "dropdown-menu-actions",
				Trigger:             lucide.Ellipsis(),
				TriggerExtraClasses: "btn-icon-ghost",
				PopoverExtraClasses: "min-w-32",
				Items: []basecoat.DropdownMenuItem{
					{Type: "item", Label: "Edit"},
					{Type: "item", Label: "Share"},
					{Type: "separator"},
					{Type: "item", Label: "Delete", Icon: lucide.Trash2(), Attrs: []HTML{X.Class("text-destructive hover:bg-destructive/10 dark:hover:bg-destructive/20 focus:bg-destructive/10 dark:focus:bg-destructive/20 focus:text-destructive [&_svg]:!text-destructive")}},
				},
			}),
			Raw(`
			<script>
			  (() => {
			    const dropdownMenu = document.querySelector('#dropdown-menu-checkboxes');
			    if (!dropdownMenu) return;
			    const checkboxes = dropdownMenu.querySelectorAll('div[role="menuitemcheckbox"]');
			    checkboxes.forEach(checkbox => {
			      checkbox.addEventListener('click', () => {
			        const isChecked = checkbox.getAttribute('aria-checked') === 'true';
			        checkbox.setAttribute('aria-checked', String(!isChecked));
			      });
			    });
			  })();
			</script>
			<script>
			  (() => {
			    const dropdownMenu = document.querySelector('#dropdown-menu-radio-group');
			    if (!dropdownMenu) return;
			    const radioButtons = dropdownMenu.querySelectorAll('div[role="menuitemradio"]');
			    radioButtons.forEach(radioButton => {
			      radioButton.addEventListener('click', () => {
			        radioButtons.forEach(item => {
			          item.setAttribute('aria-checked', 'false');
			        });
			        radioButton.setAttribute('aria-checked', 'true');
			      });
			    });
			  })();
			</script>
			`),
		),
	)
}
