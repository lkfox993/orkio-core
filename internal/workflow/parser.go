package workflow

import (
	"fmt"

	"github.com/hashicorp/hcl/v2"
	"github.com/hashicorp/hcl/v2/hclparse"

	"github.com/lkfox993/orkio-core/internal/workflow/nodes"
)

func ParseDefinition(data []byte) (*Definition, error) {
	parser := hclparse.NewParser()

	file, diags := parser.ParseHCL(data, "workflow.hcl")

	if diags.HasErrors() {
		return nil, fmt.Errorf("parse HCL: %s", diags.Error())
	}

	return parseWorkflow(file.Body)
}

func parseWorkflow(body hcl.Body) (*Definition, error) {

	schema := &hcl.BodySchema{
		Blocks: []hcl.BlockHeaderSchema{
			{
				Type:       "workflow",
				LabelNames: []string{"key"},
			},
		},
	}

	content, diags := body.Content(schema)

	if diags.HasErrors() {
		return nil, fmt.Errorf("parse workflow: %s", diags.Error())
	}

	if len(content.Blocks) != 1 {
		return nil, fmt.Errorf(
			"expected exactly one workflow block",
		)
	}

	block := content.Blocks[0]

	if len(block.Labels) != 1 {
		return nil, fmt.Errorf(
			"workflow requires exactly one key",
		)
	}

	nodes, err := parseNodes(block.Body)

	if err != nil {
		return nil, err
	}

	return &Definition{
		Key:   block.Labels[0],
		Nodes: nodes,
	}, nil
}

func parseNodes(body hcl.Body) ([]nodes.Node, error) {
	content, diags := body.Content(&hcl.BodySchema{
		Blocks: []hcl.BlockHeaderSchema{
			{
				Type:       "start_event",
				LabelNames: nil,
			},
			{
				Type:       "service_task",
				LabelNames: []string{"name"},
			},
			{
				Type:       "user_task",
				LabelNames: []string{"name"},
			},
			{
				Type:       "end_event",
				LabelNames: nil,
			},
		},
	})

	if diags.HasErrors() {
		return nil, fmt.Errorf("parse nodes: %s", diags.Error())
	}

	result := make([]nodes.Node, 0, len(content.Blocks))

	for _, block := range content.Blocks {
		node, err := parseNode(block)
		if err != nil {
			return nil, err
		}

		result = append(result, node)
	}

	return result, nil
}

func parseNode(block *hcl.Block) (nodes.Node, error) {

	switch block.Type {
	case "start_event":
		return parseStartEvent(block)

	case "service_task":
		return parseServiceTask(block)

	case "user_task":
		return parseUserTask(block)

	case "end_event":
		return parseEndEvent(block)

	default:
		return nil, fmt.Errorf("unsupported node type: %s", block.Type)
	}
}

func parseStartEvent(block *hcl.Block) (nodes.Node, error) {

	event := &nodes.StartEvent{}

	children, err := parseNodes(block.Body)

	if err != nil {
		return nil, err
	}

	event.Children = children

	return event, nil
}

func parseServiceTask(block *hcl.Block) (nodes.Node, error) {
	if len(block.Labels) != 1 {
		return nil, fmt.Errorf("service_task requires a name")
	}

	task := &nodes.ServiceTask{
		BaseNode: nodes.BaseNode{
			ID: block.Labels[0],
		},
	}

	children, err := parseNodes(block.Body)

	if err != nil {
		return nil, err
	}

	task.Children = children

	return task, nil
}

func parseUserTask(block *hcl.Block) (nodes.Node, error) {
	if len(block.Labels) != 1 {
		return nil, fmt.Errorf("service_task requires a name")
	}

	task := &nodes.UserTask{
		BaseNode: nodes.BaseNode{
			ID: block.Labels[0],
		},
	}

	children, err := parseNodes(block.Body)

	if err != nil {
		return nil, err
	}

	task.Children = children

	return task, nil
}

func parseEndEvent(block *hcl.Block) (nodes.Node, error) {
	return &nodes.EndEvent{}, nil
}
