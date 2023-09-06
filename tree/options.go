package tree

import "fmt"

type Option[T any, id comparable] func(t *Tree[T, id]) error

// WithReplaceExisting allows nodes to be replaced when inserting into the tree.
// If this option is not set, then an error will be returned if a node with the
// same nodeIndex key already exists in the tree.
func WithReplaceExisting[T any, id comparable](rep bool) Option[T, id] {
	return func(t *Tree[T, id]) error {
		t.replaceExisting = rep
		return nil
	}
}

func (t *Tree[T, id]) withOptions(opts ...Option[T, id]) error {
	for _, opt := range opts {
		err := opt(t)
		if err != nil {
			return fmt.Errorf("cannot apply Option: %w", err)
		}
	}
	return nil
}
