package exec

import (
	"fmt"
	"reflect"
)

var RootNode = newNode()

func newNode() *Node {
	return &Node{
		metadata: make(map[reflect.Type]any),
	}
}

type Node struct {
	parent   *Node
	children []*Node
	content  any // the essence of the node (e.g. Type, Field, Package)
	metadata map[reflect.Type]any
}

func (n *Node) Parent() interface{} {
	return nil
}

func (n *Node) AttachMetadata(value any) {
	n.metadata[reflect.TypeOf(value)] = value
}

func (n *Node) FetchMetdata(target any) bool {
	if target == nil {
		panic(fmt.Errorf("target cannot be nil"))
	}
	value := reflect.ValueOf(target)
	valueType := value.Type()
	if valueType.Kind() != reflect.Ptr {
		panic(fmt.Errorf("target %T must be a pointer", target))
	}
	if value.IsNil() {
		panic(fmt.Errorf("target pointer cannot be nil"))
	}
	targetRefType := valueType.Elem()
	resultValue, ok := n.metadata[targetRefType]
	if !ok {
		return false
	}
	value.Elem().Set(reflect.ValueOf(resultValue))
	return true
}

func (n *Node) AppendChild(node *Node) {
	// TODO
}
