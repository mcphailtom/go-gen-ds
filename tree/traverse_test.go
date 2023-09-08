package tree

func (tts *TreeTestSuite) TestTraversal() {
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
	}

	for _, node := range nodes {
		err := st.Insert(node)
		tts.Require().NoError(err, "Failed to insert node into tree")
	}

	tts.Run("Breadth first traversal", func() {
		// start from root node (id: 0) and limit depth at 2
		c, err := st.Traverse(TraverseBreadthFirst, 0, 2)
		tts.Require().NoError(err, "Failed to traverse tree")

		// Grab the resulting BFS output and convert to a queue for easier index-based comparison
		// returns should look like 0,1,2,3,4,7,8
		var results []int
		for node := range c {
			results = append(results, node.GetID())
		}

		bfsExpectedOrder := []int{0, 1, 2, 3, 4, 7, 8}
		tts.Equal(bfsExpectedOrder, results, "BFS traversal did not match expected order")
	})

	tts.Run("Depth first traversal", func() {
		// start from root node (id: 0) and limit depth at 2
		c, err := st.Traverse(TraverseDepthFirst, 0, 2)
		tts.Require().NoError(err, "Failed to traverse tree")

		// Grab the resulting DFS output and convert to a queue for easier index-based comparison
		// returns should look like depending on your implementation of DFS: either the left-most or right-most children first.
		// For this example, assuming right-most children first, we expect: 0,2,8,7,1,4,3
		var results []int
		for node := range c {
			results = append(results, node.GetID())
		}

		dfsExpectedOrder := []int{0, 2, 8, 7, 1, 4, 3}
		tts.Equal(dfsExpectedOrder, results, "DFS traversal did not match expected order")
	})

	tts.Run("Breadth first traversal with maxDepth = 0", func() {
		// start from root node (id: 0) without limit depth
		c, err := st.Traverse(TraverseBreadthFirst, 0, 0)
		tts.Require().NoError(err, "Failed to traverse tree")

		// Grab the resulting BFS output
		var results []int
		for node := range c {
			results = append(results, node.GetID())
		}

		// Expected result should include all nodes in the tree
		bfsExpectedOrder := []int{0, 1, 2, 3, 4, 7, 8, 9, 10, 5, 11, 6}
		tts.Equal(bfsExpectedOrder, results, "BFS traversal did not match expected order")
	})

	tts.Run("Depth first traversal with maxDepth = 0", func() {
		// start from root node (id: 0) and limit depth at 2
		c, err := st.Traverse(TraverseDepthFirst, 0, 0)
		tts.Require().NoError(err, "Failed to traverse tree")

		// Grab the resulting DFS output and convert to a queue for easier index-based comparison
		// returns should look like depending on your implementation of DFS: either the left-most or right-most children first.
		// For this example, assuming right-most children first, we expect: 0,2,8,7,1,4,3
		var results []int
		for node := range c {
			results = append(results, node.GetID())
		}

		dfsExpectedOrder := []int{0, 2, 8, 7, 1, 4, 5, 6, 3, 10, 11, 9}
		tts.Equal(dfsExpectedOrder, results, "DFS traversal did not match expected order")
	})

}
