package tree

// Node provides an interface for a generic tree node.
// The Node can be of any type, but must have a nodeIndex key of type id that is comparable.

type Node[T any, id comparable] interface {
	// GetID returns the nodeIndex key of this node.
	GetID() id
	// GetParentID returns the nodeIndex key of this node's parent.
	GetParentID() id

	// GetChildren returns a slice of child nodes of this node.
	GetChildren() []Node[T, id]

	// GetParent returns a node that is the parent of this node.
	GetParent() Node[T, id]

	// AddChildren adds a list of Nodes as children of this node.
	AddChildren(children ...Node[T, id])

	// RemoveChildren removes a list of Nodes as children of this node.
	RemoveChildren(children ...Node[T, id])

	// ReplaceChildren replaces the current children of this node with the provided list of Nodes.
	ReplaceChildren(children ...Node[T, id])

	// SetParent sets the parent of this node.
	SetParent(parent Node[T, id])

	// UpdateNode updates the node with the provided node.
	UpdateNode(node Node[T, id])
}
