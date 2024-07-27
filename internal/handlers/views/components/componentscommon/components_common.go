package componentscommon

import "github.com/a-h/templ"

type Templifier interface {
	Templify() templ.Component
}

func Templify(t Templifier) templ.Component {
	return t.Templify()
}