package tree

import (
	"errors"
	"log"
)

var (
	IndexUndefined    = errors.New("index is undefined")
	IndexNodeNotFound = errors.New("node id not found in index")
)

// index is a map of id to Node[T, id] with convenience methods for finding and inserting nodes.
type index[T any, id comparable] map[id]Node[T, id]

// find returns the Node[T, id] with the provided id, or an error if the node is not found.
func (idx *index[T, id]) find(pid id) (Node[T, id], error) {
	var val Node[T, id]
	if idx == nil {
		log.Println("Attempting to find in an undefined index")
		return val, IndexUndefined
	}
	m := *idx
	val, exists := m[pid]
	if !exists {
		return val, IndexNodeNotFound
	}
	return val, nil
}

// insert inserts the provided Node[T, id] into the index with the provided id.
func (idx *index[T, id]) insert(node Node[T, id]) error {
	if idx == nil {
		log.Println("Attempting to insert in an undefined index")
		return IndexUndefined
	}
	m := *idx
	m[node.GetID()] = node
	return nil
}

func (idx *index[T, id]) len() int {
	if idx == nil {
		return 0
	}
	return len(*idx)
}
