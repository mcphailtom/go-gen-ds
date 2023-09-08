package tree

import (
	"fmt"

	"github.com/phf/go-queue/queue"
)

// TraversalType determines the order in which an operation is performed on a tree.
type TraversalType int

const (
	TraverseBreadthFirst TraversalType = iota
	TraverseDepthFirst
)

// nodeWithDepth is a convenience struct for tracking the depth of a node in a tree.
type nodeWithDepth[T any, id comparable] struct {
	node  Node[T, id]
	depth int
}

// Traverse visits each node of a tree in a specified order, returning
// those nodes to an iterator-like channel.
// The channel is buffered to the length of the tree's nodeIndex index.
func (t *Tree[T, id]) Traverse(tt TraversalType, startNodeID id, maxDepth int) (<-chan Node[T, id], error) {
	search := make(chan Node[T, id], t.nodeIndex.len())
	startNode, err := t.nodeIndex.find(startNodeID)
	if err != nil {
		return nil, fmt.Errorf("unable to find start node: %w", err)
	}

	q := queue.New()

	sn := nodeWithDepth[T, id]{
		node:  startNode,
		depth: 0,
	}
	q.PushBack(sn)

	switch tt {
	case TraverseBreadthFirst:
		go func() {
			for {
				if bfs(q, search, maxDepth) {
					close(search)
					break
				}
			}
		}()
	case TraverseDepthFirst:
		go func() {
			for {
				if dfs(q, search, maxDepth) {
					close(search)
					break
				}
			}
		}()
	}

	return search, nil
}

// bfs performs a breadth-first search of the tree, returning each node to the search channel.
func bfs[T any, id comparable](q *queue.Queue, search chan<- Node[T, id], maxDepth int) bool {
	if q.Len() == 0 {
		return true
	}

	current := q.PopFront()
	switch cur := current.(type) {
	case nodeWithDepth[T, id]:
		childDepth := cur.depth + 1
		if maxDepth == 0 || childDepth <= maxDepth {
			for _, child := range cur.node.GetChildren() {
				cnd := nodeWithDepth[T, id]{
					node:  child,
					depth: childDepth,
				}
				q.PushBack(cnd)
			}
		}

		search <- cur.node
		return false
	case nil:
		return true
	default:
		return true
	}
}

// dfs performs a depth-first search of the tree, returning each node to the search channel.
func dfs[T any, id comparable](q *queue.Queue, search chan<- Node[T, id], maxDepth int) bool {
	if q.Len() == 0 {
		return true
	}

	current := q.PopBack()
	switch cur := current.(type) {
	case nodeWithDepth[T, id]:
		childDepth := cur.depth + 1
		if maxDepth == 0 || childDepth <= maxDepth {
			for i := len(cur.node.GetChildren()) - 1; i >= 0; i-- {
				child := cur.node.GetChildren()[i]
				cnd := nodeWithDepth[T, id]{
					node:  child,
					depth: childDepth,
				}
				q.PushBack(cnd)
			}
		}
		search <- cur.node
		return false
	case nil:
		return true
	default:
		return true
	}
}
