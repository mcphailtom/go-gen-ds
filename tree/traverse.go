package tree

import (
	"github.com/phf/go-queue/queue"
)

// TraversalType determines the order in which an operation is performed on a tree.
type TraversalType int

const (
	TraverseBreadthFirst TraversalType = iota
	TraverseDepthFirst
)

// Traverse visits each node of a tree in a specified order, returning
// those nodes to an iterator-like channel.
// The channel is buffered to the length of the tree's nodeIndex index.
func (t *Tree[T, id]) Traverse(tt TraversalType) <-chan Node[T, id] {
	search := make(chan Node[T, id], t.nodeIndex.len())

	switch tt {
	case TraverseBreadthFirst:
		q := queue.New()
		q.PushBack(t.root)
		go func() {
			for {
				if bfs(q, search) {
					close(search)
					break
				}
			}
		}()
	case TraverseDepthFirst:
		q := queue.New()
		q.PushBack(t.root)
		go func() {
			for {
				if dfs(q, search) {
					close(search)
					break
				}
			}
		}()
	}

	return search
}

// bfs performs a breadth-first search of the tree, returning each node to the search channel.
func bfs[T any, id comparable](q *queue.Queue, search chan<- Node[T, id]) bool {
	current := q.PopFront()
	switch cur := current.(type) {
	case Node[T, id]:
		for _, child := range cur.GetChildren() {
			q.PushBack(child)
		}
		search <- cur
		return false
	case nil:
		return true
	default:
		return true
	}
}

// dfs performs a depth-first search of the tree, returning each node to the search channel.
func dfs[T any, id comparable](q *queue.Queue, search chan<- Node[T, id]) bool {
	current := q.PopBack()

	switch cur := current.(type) {
	case Node[T, id]:
		for _, child := range cur.GetChildren() {
			q.PushBack(child)
		}
		search <- cur
		return false
	case nil:
		return true
	default:
		return true
	}
}
