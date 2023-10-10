package tree

func (tts *TreeTestSuite) TestNoParent() {
	st, err := NewTree[*testnode, int]()
	tts.Require().NoError(err, "Failed to initialize tree for testing")

	nodes := []Node[*testnode, int]{
		&testnode{id: 0, parentID: 0, name: "node0"},
		&testnode{id: 1, parentID: 0, name: "node1"},
		&testnode{id: 2, parentID: 0, name: "node2"},
		&testnode{id: 3, parentID: 1, name: "node3"},
		&testnode{id: 4, parentID: 1, name: "node4"},
		&testnode{id: 5, parentID: 4, name: "node5"},
		&testnode{id: 6, parentID: 5, name: "node6"},
		&testnode{id: 7, parentID: 2, name: "node7"},
		&testnode{id: 8, parentID: 2, name: "node8"},
		&testnode{id: 9, parentID: 3, name: "node9"},
		&testnode{id: 10, parentID: 3, name: "node10"},
		&testnode{id: 11, parentID: 10, name: "node11"},
		&testnode{id: 13, parentID: 12, name: "node13"},
	}

	for _, node := range nodes {
		err := st.Insert(node)
		if node.GetID() == 13 {
			tts.Require().Error(err, "Node without parent should trigger error")
			continue
		}
		tts.Require().NoError(err, "Failed to insert node into tree")
	}
}

func (tts *TreeTestSuite) TestReroot() {
	st, err := NewTree[*testnode, int]()
	tts.Require().NoError(err, "Failed to initialize tree for testing")

	nodes := []Node[*testnode, int]{
		&testnode{id: 0, parentID: 12, name: "node0"},
		&testnode{id: 1, parentID: 0, name: "node1"},
		&testnode{id: 2, parentID: 0, name: "node2"},
		&testnode{id: 3, parentID: 1, name: "node3"},
		&testnode{id: 4, parentID: 1, name: "node4"},
		&testnode{id: 5, parentID: 4, name: "node5"},
		&testnode{id: 6, parentID: 5, name: "node6"},
		&testnode{id: 7, parentID: 2, name: "node7"},
		&testnode{id: 8, parentID: 2, name: "node8"},
		&testnode{id: 9, parentID: 3, name: "node9"},
		&testnode{id: 10, parentID: 3, name: "node10"},
		&testnode{id: 11, parentID: 10, name: "node11"},
		&testnode{id: 12, parentID: 27, name: "node13"},
	}

	for _, node := range nodes {
		err := st.Insert(node)
		tts.Require().NoError(err, "Failed to insert node into tree")
	}

	// start from root node (id: 0) without limit depth
	c, err := st.Traverse(TraverseBreadthFirst, st.Root().GetID(), 0)
	tts.Require().NoError(err, "Failed to traverse tree")

	// Grab the resulting BFS output
	var results []int
	for node := range c {
		results = append(results, node.GetID())
	}

	// Expected result should include all nodes in the tree
	bfsExpectedOrder := []int{12, 0, 1, 2, 3, 4, 7, 8, 9, 10, 5, 11, 6}
	tts.Equal(bfsExpectedOrder, results, "BFS traversal did not match expected order")
}
