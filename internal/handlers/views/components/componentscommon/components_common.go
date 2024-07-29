package componentscommon

import "github.com/a-h/templ"

type Templifier interface {
	Templify() templ.Component
}

func Templify(t Templifier) templ.Component {
	return t.Templify()
}

type NavItem struct {
	Text          string
	URL           string
	PushURLString string
}

type Header struct {
	NavItems []NavItem
}

func (i NavItem) PushURL() string {
	if i.PushURLString != "" {
		return i.PushURLString
	} else {
		return "true"
	}
}
