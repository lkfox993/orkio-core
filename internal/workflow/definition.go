package workflow

import "github.com/lkfox993/orkio-core/internal/workflow/nodes"

type Definition struct {
	Key   string       `hcl:"key,optional"`
	Nodes []nodes.Node `hcl:"job,block"`
}

func (d *Definition) StartNodes() []nodes.Node {
	var result []nodes.Node

	for _, node := range d.Nodes {
		if node.NodeType() == nodes.NodeTypeStartEvent {
			result = append(result, node)
		}
	}

	return result
}
