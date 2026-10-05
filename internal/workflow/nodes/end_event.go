package nodes

const (
	NodeTypeEndEvent = "end_event"
)

type EndEvent struct {
	BaseNode
}

func (EndEvent) NodeType() string {
	return "end_event"
}
