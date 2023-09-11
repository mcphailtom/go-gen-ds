package tree

import "fmt"

type Option[T any, id comparable] func(t *Tree[T, id]) error

// WithUpdatesAllowed allows updates to existing nodes.
func WithUpdatesAllowed[T any, id comparable](rep bool) Option[T, id] {
	return func(t *Tree[T, id]) error {
		t.updatesAllowed = rep
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
