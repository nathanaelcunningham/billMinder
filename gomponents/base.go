package gomponents

import (
	g "github.com/maragudk/gomponents"
	. "github.com/maragudk/gomponents/html"
)

func Page(title, currentPath string, body g.Node) g.Node {
	return Doctype(
		HTML(
			Lang("en"),
			Head(
				TitleEl(g.Text(title)),
				Meta(Charset("utf-8")),
				Meta(
					Name("viewport"),
					Content("width=device-width, initial-scale=1"),
				),
				Link(Rel("shortcut icon"), Href("#")),
				Link(Rel("stylesheet"), Href("https://unpkg.com/css.gg@2.0.0/icons/css/bolt.css")),
				Script(Src("/static/js/tailwind.js")),
				Script(Src("/static/js/htmx.min.js")),
				Script(Src("https://unpkg.com/hyperscript.org@0.9.12")),
			),
			Body(
				Class("p-6 bg-gray-100"),
				Navbar(currentPath),
				Container(body),
			),
		),
	)
}

func Container(children ...g.Node) g.Node {
	return Div(g.Group(children))
}

func About() g.Node {
	return Div(
		g.Text("About"),
	)
}
