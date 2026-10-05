package nodes

const (
	NodeTypeMessageEvent = "message_event"
)

type MessageEvent struct {
	BaseNode
}

func (EndEvent) MessageType() string {
	return "message_event"
}
