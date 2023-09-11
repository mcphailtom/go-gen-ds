// Package tree implements a generic tree data structure with an index.
// It includes Breadth First and Depth First traversal methods.
package tree

import (
	"errors"
	"reflect"
)

var (
	NodeExists          = errors.New("node already exists")
	ParentNotFound      = errors.New("node parent not found")
	CyclicReference     = errors.New("node parent is child of node")
	Undefined           = errors.New("tree is nil or undefined")
	MergeDuplicateNodes = errors.New("duplicate nodes in source and target trees")
)

// Tree implements a tree based data structure with an index.
type Tree[T any, id comparable] struct {
	root           Node[T, id]
	nodeIndex      *index[T, id]
	updatesAllowed bool
}

// NewTree creates and returns an empty tree.
func NewTree[T any, id comparable](options ...Option[T, id]) (*Tree[T, id], error) {
	t := Tree[T, id]{
		nodeIndex:      &index[T, id]{},
		updatesAllowed: false,
	}

	if err := t.withOptions(options...); err != nil {
		return nil, err
	}

	return &t, nil
}

// Root returns the root node of the tree. nil indicates that the tree is empty.
func (t *Tree[T, id]) Root() Node[T, id] {
	return t.root
}

// Insert adds a node to the tree.
// Replacement of an existing node is handled by setting the WitheReplaceExisting option
// on the tree at initialization. This is not enabled by default and will cause an error
// if a node with the same nodeIndex key already exists in the tree.
//
// In the case the tree does not have a root this node will become the new root node
// of the tree.
//
// If the parent of the node does not exist in the tree, then the node is not added unless
// it is the parent of the existing root which will cause the tree to be rerooted.
func (t *Tree[T, id]) Insert(node Node[T, id]) error {

	// Check for the existence of the nodeIndex key in the index
	existingNode, err := t.nodeIndex.find(node.GetID())
	if err == nil {
		if !t.updatesAllowed {
			return NodeExists
		}
		existingNode.UpdateNode(node)
		return nil
	}

	if reflect.ValueOf(&t.root).Elem().IsZero() { // always insert the first element
		t.root = node
	} else {
		parent, err := t.nodeIndex.find(node.GetParentID())
		if err != nil {
			if t.root.GetParentID() == node.GetID() { // parent does not exist but incoming node is parent of root
				t.reroot(node)
			} else { // parent does not exist, do not add
				return ParentNotFound
			}
		}

		if t.root.GetParentID() == node.GetID() { // parent exists, but incoming node causes cycle from root
			return CyclicReference
		}

		// parent exists, add
		node.SetParent(parent)
		node.ReplaceChildren() // Reset children, if any
		parent.AddChildren(node)
	}

	// add to nodeIndex index
	t.nodeIndex.insert(node)
	return nil
}

// Exists returns true if the node exists in the tree.
func (t *Tree[T, id]) Exists(pid id) bool {
	_, err := t.nodeIndex.find(pid)
	return err == nil
}

// Merge the source tree (passed in the argument) into the target tree.
// The trees are merged only if a relationship can be found, constituted by
// the parent of the head of the other tree is found in the target tree.
//
// The merge will fail if:
// - The source tree is nil
// - There are duplicate nodeIndex keys between the two trees.
// - The parent of the head of the other tree is not found in the target tree.
func (t *Tree[T, id]) Merge(other *Tree[T, id]) error {

	if other == nil {
		return Undefined
	}

	headParent := other.root.GetParentID()
	// check for parent of head in target tree
	f, err := t.nodeIndex.find(headParent)
	if err != nil {
		return ParentNotFound
	}

	// check for duplicate nodeIndex ids
	for k := range *other.nodeIndex {
		if _, err := t.nodeIndex.find(k); err == nil {
			return MergeDuplicateNodes
		}
	}

	f.AddChildren(other.root)
	other.root.SetParent(f)

	// copy other index to new tree
	for _, n := range *other.nodeIndex {
		t.nodeIndex.insert(n)
	}

	return nil
}

// Find looks up a node by its nodeIndex key. The function returns a node value
// and a boolean indicating whether the node was found in the tree.
func (t *Tree[T, id]) Find(pid id) (Node[T, id], bool) {
	f, err := t.nodeIndex.find(pid)
	if err != nil {
		return f, false
	}
	return f, true
}

// FindParents returns a slice of parent nodes of the node with the provided nodeIndex key.
// Parents are collected from the node's parent up to the root of the tree.
// The function returns a slice of nodes and a boolean indicating whether the node was found in the tree.
func (t *Tree[T, id]) FindParents(pid id) ([]Node[T, id], bool) {
	var parents []Node[T, id]
	f, err := t.nodeIndex.find(pid)
	if err != nil {
		return parents, false
	}

	for n := f.GetParent(); n != nil; n = n.GetParent() {
		parents = append(parents, n)
	}

	return parents, true
}

// reroot reassigns the root node of the
func (t *Tree[T, id]) reroot(newHead Node[T, id]) {
	t.root.SetParent(newHead)
	newHead.AddChildren(t.root)
	t.root = newHead
}
