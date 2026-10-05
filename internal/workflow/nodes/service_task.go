package nodes

const (
	NodeTypeServiceTask = "service_task"
)

type ServiceTask struct {
	BaseNode
	Retry *Retry
}

func (ServiceTask) NodeType() string {
	return "service_task"
}
