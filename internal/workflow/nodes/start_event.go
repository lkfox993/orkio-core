package nodes

const (
	NodeTypeStartEvent = "start_event"
)

type StartEvent struct {
	BaseNode
}

func (StartEvent) NodeType() string {
	return "start_event"
}
