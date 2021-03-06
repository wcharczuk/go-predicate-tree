package predicate

import (
	"encoding/json"

	"github.com/blendlabs/go-exception"
)

const (
	GraphFieldID             = "id"
	GraphFieldType           = "type"
	GraphFieldChildren       = "children"
	GraphFieldPredicateType  = "predicate_type"
	GraphFieldPredicateState = "predicate_state"
)

func Serialize(node Node) ([]byte, error) {
	graph, err := serializeGraph(node)
	if err != nil {
		return nil, err
	}
	return json.Marshal(graph)
}

func Deserialize(data []byte) (Node, error) {
	graph, err := deserializeGraph(data)
	if err != nil {
		return nil, err
	}
	return umarshalIntermediateNode(graph), nil
}

func umarshalIntermediateNode(graph *intermediateNode) Node {
	nodeType := graph.Type

	node := createNodeByType(nodeType)

	node.SetID(graph.ID)

	switch nodeType {
	case NodeTypeEval:
		{
			if typed, isTyped := node.(*EvalNode); isTyped {
				typed.predicateType = graph.PredicateType
				predicate := CreatePredicate(typed.predicateType)
				if len(graph.Predicate) != 0 {
					err := json.Unmarshal(graph.Predicate, predicate)
					if err == nil {
						typed.predicate = predicate
					}
				}
				typed.predicate = predicate
				return typed
			}
		}
	default:
		{
			if len(graph.Children) != 0 {
				for _, child := range graph.Children {
					node.AddChild(umarshalIntermediateNode(child))
				}
			}
		}
	}
	return node
}

func createNodeByType(nodeType string) Node {
	switch nodeType {
	case NodeTypeAnd:
		return &AndNode{}
	case NodeTypeEval:
		return &EvalNode{}
	case NodeTypeNot:
		return &NotNode{}
	case NodeTypeOr:
		return &OrNode{}
	}
	return nil
}

func serializeGraph(node Node) (map[string]interface{}, error) {
	graphNode := map[string]interface{}{}
	nodeType := node.Type()
	children := node.Children()

	graphNode[GraphFieldID] = node.ID()
	graphNode[GraphFieldType] = nodeType

	// custom steps for specific node types.
	switch nodeType {
	case NodeTypeEval:
		{
			if evalNode, isEvalNode := node.(*EvalNode); isEvalNode {
				graphNode[GraphFieldPredicateType] = evalNode.predicateType
				if evalNode.predicate != nil {
					graphNode[GraphFieldPredicateState] = evalNode.predicate
				}
			} else {
				return nil, exception.New("Node Type Value: `eval` but incorrect type assertion.")
			}
		}
	default:
		{
			if len(children) != 0 {
				var edgeNodes []interface{}
				for _, child := range children {
					edgeNode, err := serializeGraph(child)
					if err != nil {
						return nil, err
					}
					edgeNodes = append(edgeNodes, edgeNode)
				}
				graphNode[GraphFieldChildren] = edgeNodes
			}
		}
	}

	return graphNode, nil
}

func deserializeGraph(data []byte) (*intermediateNode, error) {
	graph := &intermediateNode{}
	err := json.Unmarshal(data, graph)
	return graph, err
}

type intermediateNode struct {
	ID            string              `json:"id"`
	Type          string              `json:"type"`
	Children      []*intermediateNode `json:"children"`
	PredicateType string              `json:"predicate_type"`
	Predicate     json.RawMessage     `json:"predicate_state"`
}
