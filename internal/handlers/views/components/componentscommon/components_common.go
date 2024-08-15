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

type hxAttr struct {
	Key HxAttrKey
	Val HxAttrVal
}

func NewHxAttr(k HxAttrKey, v HxAttrVal) hxAttr {
	return hxAttr{Key: k}
}

type HxAttrKey string
type HxAttrVal string

func NewURL(url string) HxAttrVal {
	return HxAttrVal(url)
}

const (
	HxGet     HxAttrKey = "hx-get"
	HxPost    HxAttrKey = "hx-post"
	HxPut     HxAttrKey = "hx-put"
	HxDelete  HxAttrKey = "hx-delete"
	HxPatch   HxAttrKey = "hx-patch"
	HxTrigger HxAttrKey = "hx-trigger"
	HxSwap    HxAttrKey = "hx-swap"
	HxPushURL HxAttrKey = "hx-push-url"
)

const (
	SwapNone     HxAttrVal = ""
	SwapOuter    HxAttrVal = "outerHTML"
	SwapInner    HxAttrVal = "innerHTML"
	PushURLTrue  HxAttrVal = "true"
	PushURLFalse HxAttrVal = "false"
)
