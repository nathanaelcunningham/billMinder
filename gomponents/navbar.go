package gomponents

import (
	g "github.com/maragudk/gomponents"
	c "github.com/maragudk/gomponents/components"
	. "github.com/maragudk/gomponents/html"
	s "github.com/maragudk/gomponents/svg"
)

func Navbar(currentPath string) g.Node {
	return Header(
		Class("bg-white"),
		Nav(
			Class("flex justify-between items-center py-4 px-6 bg-white shadow-md"),
			NavbarIcon("/"),
			Hamburger(),
			NavbarItems(currentPath),
		),
		NavbarHamburger(currentPath),
	)
}

func NavbarIcon(href string) g.Node {
	return A(
		Href(href),
		Class("text-xl font-bold text-gray-800 hover:text-gray-700"),
		g.Text("Bill Minder"),
	)
}

func NavbarHamburger(currentPath string) g.Node {
	linkClass := "block px-4 py-2 text-gray-800 hover:bg-gray-100"
	return Div(
		ID("mobileMenu"),
		Class("hidden lg:hidden"),
		NavbarLink("/", "Bills", currentPath, &linkClass),
		NavbarLink("/about", "About", currentPath, &linkClass),
	)
}

func Hamburger() g.Node {
	return Div(
		Class("lg:hidden"),
		HScript("on click toggle .hidden on #mobileMenu"),
		Div(
			Class("text-gray-800 focus:outline-none"),
			SVG(
				Class("w-6 h-6"),
				s.Fill("none"),
				s.Stroke("currentColor"),
				g.Raw(`stroke-linecap="round"`),
				g.Raw(`stroke-linejoin="round"`),
				s.StrokeWidth("2"),
				s.ViewBox("0 0 24 24"),
				s.Path(s.D("M4 6h16M4 12h16m-7 6h7")),
			),
		),
	)
}

func NavbarItems(currentPath string) g.Node {
	linkClass := "px-4 text-gray-800 hover:text-gray-700"
	return Div(
		Class("hidden lg:flex"),
		NavbarLink("/", "Bills", currentPath, &linkClass),
		NavbarLink("/about", "About", currentPath, &linkClass),
	)
}

func NavbarLink(href, name, currentPath string, class *string) g.Node {
	return A(
		Href(href),
		c.Classes{"text-sky-600": currentPath == href, *class: class != nil},
		g.Text(name),
	)
}
