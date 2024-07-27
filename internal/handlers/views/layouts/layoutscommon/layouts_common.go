package layoutscommon

type LayoutTarget string

const (
	Details LayoutTarget = "details"
	Modal   LayoutTarget = "modal"
)

func (lt LayoutTarget) String() string {
	return string(lt)
}

func (lt LayoutTarget) Selector() string {
	return "#" + lt.String()
}
