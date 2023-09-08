package tree

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

// testnode is a test implementation of a Node.
type testnode struct {
	id       int
	parentID int
	name     string
	children []Node[*testnode, int]
	parent   Node[*testnode, int]
}

func (c *testnode) GetID() int {
	return c.id
}

func (c *testnode) GetParentID() int {
	return c.parentID
}

func (c *testnode) GetChildren() []Node[*testnode, int] {
	return c.children
}

func (c *testnode) GetParent() Node[*testnode, int] {
	return c.parent
}

func (c *testnode) AddChildren(children ...Node[*testnode, int]) {
	for _, v := range children {
		c.children = append(c.children, v)
	}
}

func (c *testnode) RemoveChildren(children ...Node[*testnode, int]) {
	for _, v := range children {
		for i, child := range c.children {
			if child.GetID() == v.GetID() {
				c.children = append(c.children[:i], c.children[i+1:]...)
			}
		}
	}
}

func (c *testnode) ReplaceChildren(children ...Node[*testnode, int]) {
	c.children = []Node[*testnode, int]{}
	c.AddChildren(children...)
}

func (c *testnode) SetParent(parent Node[*testnode, int]) {
	if parent == nil || parent.GetID() == c.GetID() {
		return
	}
	c.parent = parent
}

type TreeTestSuite struct {
	suite.Suite
}

func (tts *TreeTestSuite) SetupSuite() {
}

func TestTreeTestSuite(t *testing.T) {
	suite.Run(t, new(TreeTestSuite))
}
