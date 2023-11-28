package gomponents

import (
	"io"

	g "github.com/maragudk/gomponents"
)

type rawAttr struct {
	name  string
	value string
}

func HScript(v string) g.Node {
	return attr(v)
}

func attr(value string) g.Node {
	return g.Attr("_", value)
}

func (r *rawAttr) Render(w io.Writer) error {
	_, err := w.Write([]byte(" _" + `="` + r.value + `"`))
	return err
}

func (r *rawAttr) Type() g.NodeType {
	return g.AttributeType
}
