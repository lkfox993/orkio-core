package nodes

type Node interface {
	NodeType() string
	GetChildren() []Node
	Base() *BaseNode
}

type BaseNode struct {
	ID       string
	Children []Node
}

func (n *BaseNode) Base() *BaseNode {
	return n
}

func (n *BaseNode) GetChildren() []Node {
	return n.Children
}

type Retry struct {
	Attempts int
}
