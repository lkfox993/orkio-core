package nodes

const (
	NodeTypeUserTask = "user_task"
)

type UserTask struct {
	BaseNode
	Retry *Retry
}

func (UserTask) NodeType() string {
	return "user_task"
}
